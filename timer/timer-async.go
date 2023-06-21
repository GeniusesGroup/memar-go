/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"strconv"
	"unsafe"

	"libgo/protocol"
	"libgo/race"
	"libgo/scheduler"
	"libgo/time/monotonic"
)

// NewAsync waits for the duration to elapse and then calls callback.
// If callback need blocking operation it must do its logic in new thread(goroutine).
// It returns a Timer that can be used to cancel the call using its Stop method.
func NewAsync(d protocol.Duration, callback protocol.TimerListener) (t *Async, err protocol.Error) {
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
	period protocol.Duration

	// The status field holds one of the values in status file.
	status status

	// callback function that call when reach
	// it is possible that callback will be called a little after the delay.
	// * NOTE: each time calling callback() in the timer goroutine, so callback must be
	// * a well-behaved function and not block.
	callback protocol.TimerListener

	timing *TimingHeap
}

// Init initialize the timer with given callback function or make the channel and send signal on it
// Be aware that given function must not be closure and must not block the caller.
//
//libgo:impl libgo/protocol.Timer
func (t *Async) Init(callback protocol.TimerListener) (err protocol.Error) {
	if t.callback != nil {
		err = &ErrTimerAlreadyInit
		return
	}

	t.callback = callback
	return
}

//libgo:impl libgo/protocol.SoftwareLifeCycle
func (t *Async) Reinit(callback protocol.TimerListener) (err protocol.Error) {
	// var status = t.status.Load()
	// if !(status == protocol.TimerStatus_Unset || status == protocol.TimerStatus_Deleted) {
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

//libgo:impl libgo/protocol.Timer
func (t *Async) Status() (activeStatus protocol.TimerStatus) { return t.status.Load() }

// Start adds the timer to the running cpu core timing.
// This should only be called with a newly created timer.
// That avoids the risk of changing the when field of a timer in some P's heap,
// which could cause the heap to become unsorted.
//
//libgo:impl libgo/protocol.Timer
func (t *Async) Start(d protocol.Duration) (err protocol.Error) {
	if t.callback == nil {
		err = &ErrTimerNotInit
		return
	}
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timing.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		err = &ErrNegativeDuration
		return
	}
	var activeStatus = t.status.Load()
	if activeStatus != protocol.TimerStatus_Unset || t.timing != nil {
		err = &ErrTimerAlreadyStarted
		return
	}

	if !t.status.CompareAndSwap(protocol.TimerStatus_Unset, protocol.TimerStatus_Waiting) {
		err = &ErrTimerRacyAccess
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
//libgo:impl libgo/protocol.Timer
func (t *Async) Stop() (err protocol.Error) {
	if t.callback == nil {
		err = &ErrTimerNotInit
		return
	}

	var activeStatus protocol.TimerStatus
	for {
		activeStatus = t.status.Load()
		switch activeStatus {
		case protocol.TimerStatus_Unset:
			err = &ErrTimerNotInit
			return
		case protocol.TimerStatus_Waiting, protocol.TimerStatus_ModifiedLater, protocol.TimerStatus_ModifiedEarlier:
			// Must fetch t.timing before changing status,
			// due to ts.cleanTimers in another goroutine can clear t.timing of timing in protocol.TimerStatus_Deleted status.
			var timing = t.timing

			// Timer was not yet run.
			if t.status.CompareAndSwap(activeStatus, protocol.TimerStatus_Deleted) {
				timing.deletedTimers.Add(1)
				return
			}
		case protocol.TimerStatus_Deleted, protocol.TimerStatus_Removing, protocol.TimerStatus_Removed:
			// Timer was already run.
			return
		case protocol.TimerStatus_Running, protocol.TimerStatus_Moving:
			// The timer is being run or moved, by a different P Wait for it to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case protocol.TimerStatus_Modifying:
			// Simultaneous calls to Reset(). Wait for the other call to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		default:
			err = &ErrTimerBadStatus
			return
		}
	}
}

// Reset modifies an existing timer to new deadline.
// It's OK to call Reset() on a newly allocated Timer.
// Reports whether the timer was modified before it was run.
//
//libgo:impl libgo/protocol.Timer
func (t *Async) Reset(d protocol.Duration) (err protocol.Error) {
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timing.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		err = &ErrNegativeDuration
		return
	}
	if t.callback == nil {
		err = &ErrTimerNotInit
		return
	}

	if race.DetectorEnabled {
		race.Release(unsafe.Pointer(t))
	}

	var wasRemovedFromTiming = false
	var activeStatus protocol.TimerStatus
loop:
	for {
		activeStatus = t.status.Load()
		switch activeStatus {
		case protocol.TimerStatus_Waiting, protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
			if t.status.CompareAndSwap(activeStatus, protocol.TimerStatus_Modifying) {
				break loop
			}
		case protocol.TimerStatus_Unset, protocol.TimerStatus_Removed:
			// Timer was already run and t is no longer in a timing.
			// Act like AddTimer.
			if t.status.CompareAndSwap(activeStatus, protocol.TimerStatus_Modifying) {
				wasRemovedFromTiming = true
				break loop
			}
		case protocol.TimerStatus_Deleted:
			if t.status.CompareAndSwap(activeStatus, protocol.TimerStatus_Modifying) {
				t.timing.deletedTimers.Add(-1)
				break loop
			}
		case protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Moving:
			// The timer is being run or moved, by a different P.
			// Wait for it to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case protocol.TimerStatus_Modifying:
			// Multiple simultaneous calls to Reset().
			// Wait for the other call to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		default:
			err = &ErrTimerBadStatus
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
		if !t.status.CompareAndSwap(protocol.TimerStatus_Modifying, protocol.TimerStatus_Waiting) {
			err = &ErrTimerRacyAccess
			// TODO::: Easily just return??
			return
		}
	} else {
		// TODO::: as describe here: https://github.com/golang/go/issues/53953#issuecomment-1189769955
		// we need to access to timerBucket.when to decide correctly about new timer status,
		// updateTimerModifiedEarliest() may call wrongly and waste resource. Any idea to fix?
		var newStatus = protocol.TimerStatus_ModifiedLater
		if timerNewWhen < timerOldWhen {
			newStatus = protocol.TimerStatus_ModifiedEarlier
			t.timing.updateTimerModifiedEarliest(timerNewWhen)
		}

		// Set the new status of the timer.
		if !t.status.CompareAndSwap(protocol.TimerStatus_Modifying, newStatus) {
			err = &ErrTimerRacyAccess
			// TODO::: Easily just return??
			return
		}
	}

	return
}

// Tick will send a signal on the t.Signal() channel after each tick on initialized Timer.
// The period of the ticks is specified by the duration arguments.
// The ticker will adjust the time interval or drop ticks to make up for slow receivers.
// The durations must be greater than zero; if not, Tick() will panic.
// Stop the ticker to release associated resources.
//
//libgo:impl libgo/protocol.Ticker
func (t *Async) Tick(first, interval protocol.Duration) (err protocol.Error) {
	if first < 1 || interval < 1 {
		err = &ErrNegativeDuration
		return
	}
	t.period = interval
	err = t.Start(first)
	return
}

//libgo:impl libgo/protocol.Stringer
func (t *Async) ToString() string {
	var until = t.when.UntilNow()
	var untilSecond = until / monotonic.Second
	var untilSecondString = strconv.FormatInt(int64(untilSecond), 10)
	return "Timer sleep for " + untilSecondString + " seconds"
}
func (t *Async) FromString(s string) (err protocol.Error) { return }
