/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"unsafe"

	"memar/math/integer"
	"memar/protocol"
	"memar/runtime/race"
	"memar/runtime/scheduler"
	"memar/time/duration"
	"memar/time/monotonic"
	errs "memar/timer/errors"
	timer_p "memar/timer/protocol"
)

// NewAsync waits for the duration to elapse and then calls callback.
// If callback need blocking operation it must do its logic in new thread(goroutine).
// It returns a Timer that can be used to cancel the call using its Stop method.
func NewAsync(d duration.NanoSecond, callback timer_p.TimerListener) (t *Async, err protocol.Error) {
	var timer Async
	timer.Init(callback)
	err = timer.Start(d)
	t = &timer
	return
}

// Async is a async timer object.
// - It is not safe to call its method concurrently.
// - It can be cause memory leak if you embed it directly, Due to Timing can't remove reference to it quickly.
type Async struct {
	// Timer wakes up at when, and then at when+period, ... (period > 0 only)
	// when must be positive on an active timer.
	when   monotonic.Time
	period duration.NanoSecond

	// The status field holds one of the values in status file.
	status Status

	// callback function that call when reach
	// it is possible that callback will be called a little after the delay.
	// * NOTE: each time calling callback() in the timer goroutine, so callback must be
	// * a well-behaved function and not block.
	callback timer_p.TimerListener

	timing *Timing
}

// Init initialize the timer with given callback.
//
//memar:impl memar/protocol.Timer
func (t *Async) Init(callback timer_p.TimerListener) (err protocol.Error) {
	if t.callback != nil {
		err = &errs.ErrTimerAlreadyInit
		return
	}

	t.callback = callback
	return
}

//memar:impl memar/protocol.SoftwareLifeCycle
func (t *Async) Reinit(callback timer_p.TimerListener) (err protocol.Error) {
	// var status = t.status.Load()
	// if !(status == Status_Unset || status == Status_Deleted) {
	// 	panic("timer: Reinit called with non stopped timer")
	// }
	err = t.Stop()
	if err != nil {
		return
	}
	t.callback = callback
	t.timing = nil
	return
}
func (t *Async) Deinit() (err protocol.Error) {
	err = t.Stop()
	// TODO::: Can we remove t from related timing heap?
	return
}

//memar:impl memar/protocol.Timer
func (t *Async) Status() (activeStatus Status) { return t.status.Load() }
func (t *Async) When() monotonic.Time          { return t.when }

// Start adds the timer to the running cpu core timing.
// This should only be called with a newly created timer.
// That avoids the risk of changing the when field of a timer in some P's heap,
// which could cause the heap to become unsorted.
//
//memar:impl memar/protocol.Timer
func (t *Async) Start(d duration.NanoSecond) (err protocol.Error) {
	if t.callback == nil {
		err = &errs.ErrTimerNotInit
		return
	}
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timing.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		err = &errs.ErrNegativeDuration
		return
	}
	var activeStatus = t.status.Load()
	if activeStatus != Status_Unset || t.timing != nil {
		err = &errs.ErrTimerAlreadyStarted
		return
	}

	if !t.status.CompareAndSwap(Status_Unset, Status_Waiting) {
		err = &errs.ErrTimerRacyAccess
		return
	}

	if race.DetectorEnabled {
		race.Release(unsafe.Pointer(t))
	}

	t.when = when(d)
	t.timing = getActiveTiming()
	t.timing.AddTimer(t)
	return
}

// Stop deletes the timer t. We can't actually remove it from the timing heap.
// We can only mark it as deleted. It will be removed in due course by the timing whose heap it is on.
// Reports whether the timer was removed before it was run.
//
//memar:impl memar/protocol.Timer
func (t *Async) Stop() (err protocol.Error) {
	if t.callback == nil {
		err = &errs.ErrTimerNotInit
		return
	}

	var activeStatus Status
	for {
		activeStatus = t.status.Load()
		switch activeStatus {
		case Status_Unset:
			err = &errs.ErrTimerNotInit
			return
		case Status_Waiting, Status_ModifiedLater, Status_ModifiedEarlier:
			// Must fetch t.timing before changing status,
			// due to ts.cleanTimers in another goroutine can clear t.timing of timing in Status_Deleted status.
			var timing = t.timing

			// Timer was not yet run.
			if t.status.CompareAndSwap(activeStatus, Status_Deleted) {
				timing.deletedTimersCount.Add(1)
				return
			}
		case Status_Deleted, Status_Removing, Status_Removed:
			// Timer was already run.
			return
		case Status_Running, Status_Moving:
			// The timer is being run or moved, by a different P Wait for it to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case Status_Modifying:
			// Simultaneous calls to Reset(). Wait for the other call to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		default:
			err = &errs.ErrTimerBadStatus
			return
		}
	}
}

// Reset modifies an existing timer to new deadline.
// It's OK to call Reset() on a newly allocated Timer.
// Reports whether the timer was modified before it was run.
//
//memar:impl memar/protocol.Timer
func (t *Async) Reset(d duration.NanoSecond) (err protocol.Error) {
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timing.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		err = &errs.ErrNegativeDuration
		return
	}
	if t.callback == nil {
		err = &errs.ErrTimerNotInit
		return
	}

	if race.DetectorEnabled {
		race.Release(unsafe.Pointer(t))
	}

	var wasRemovedFromTiming = false
	var activeStatus Status
loop:
	for {
		activeStatus = t.status.Load()
		switch activeStatus {
		case Status_Waiting, Status_ModifiedEarlier, Status_ModifiedLater:
			if t.status.CompareAndSwap(activeStatus, Status_Modifying) {
				break loop
			}
		case Status_Unset, Status_Removed:
			// Timer was already run and t is no longer in a timing.
			// Act like AddTimer.
			if t.status.CompareAndSwap(activeStatus, Status_Modifying) {
				wasRemovedFromTiming = true
				break loop
			}
		case Status_Deleted:
			if t.status.CompareAndSwap(activeStatus, Status_Modifying) {
				t.timing.deletedTimersCount.Add(-1)
				break loop
			}
		case Status_Running, Status_Removing, Status_Moving:
			// The timer is being run or moved, by a different P.
			// Wait for it to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case Status_Modifying:
			// Multiple simultaneous calls to Reset().
			// Wait for the other call to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		default:
			err = &errs.ErrTimerBadStatus
			return
		}
	}

	var timerOldWhen = t.when
	var timerNewWhen = when(d)
	t.when = timerNewWhen
	if t.period != 0 {
		t.period = d
	}
	if wasRemovedFromTiming {
		t.timing = getActiveTiming()
		t.timing.AddTimer(t)
		if !t.status.CompareAndSwap(Status_Modifying, Status_Waiting) {
			err = &errs.ErrTimerRacyAccess
			// TODO::: Easily just return??
			return
		}
	} else {
		// TODO::: as describe here: https://github.com/golang/go/issues/53953#issuecomment-1189769955
		// we need to access to timerBucket.when to decide correctly about new timer status,
		// updateTimerModifiedEarliest() may call wrongly and waste resource. Any idea to fix?
		var newStatus = Status_ModifiedLater
		if timerNewWhen < timerOldWhen {
			newStatus = Status_ModifiedEarlier
			t.timing.updateTimerModifiedEarliest(timerNewWhen)
		}

		// Set the new status of the timer.
		if !t.status.CompareAndSwap(Status_Modifying, newStatus) {
			err = &errs.ErrTimerRacyAccess
			// TODO::: Easily just return??
			return
		}
	}

	return
}

// Tick will call the t.callback.TimerHandler() after each tick on initialized Timer.
// The period of the ticks is specified by the duration arguments.
// The ticker will adjust the time interval or drop ticks to make up for slow receivers.
// The durations must be greater than zero; if not, Tick() will panic.
// Stop the ticker to release associated resources.
//
//memar:impl memar/protocol.Ticker
func (t *Async) Tick(first, interval duration.NanoSecond) (err protocol.Error) {
	if first < 1 || interval < 1 {
		err = &errs.ErrNegativeDuration
		return
	}
	t.period = interval
	err = t.Start(first)
	return
}

//memar:impl memar/protocol.Stringer
func (t *Async) ToString() (str string, err protocol.Error) {
	var until = t.when.UntilNow()
	var untilSecond = until / duration.OneSecond
	var untilSecondINT = integer.S64(untilSecond)
	var untilSecondString string
	untilSecondString, err = untilSecondINT.ToString()
	str = "Timer sleep for " + untilSecondString + " seconds"
	return
}
func (t *Async) FromString(str string) (err protocol.Error) { return }
