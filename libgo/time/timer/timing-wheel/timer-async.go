/* For license and copyright information please see the LEGAL file in the code repository */

package timing_wheel

import (
	"memar/computer/runtime/scheduler"
	error_p "memar/process/error/protocol"
	race "memar/process/race/thread_sanitizer"
	"memar/time/duration"
	"memar/time/monotonic"
	"memar/time/timer"
	timer_errs "memar/time/timer/errors"
	timer_p "memar/time/timer/protocol"
)

// NewAsync waits for the duration to elapse and then calls callback.
// If callback need blocking operation it must do its logic in new thread(goroutine).
// It returns a Timer that can be used to cancel the call using its Stop method.
func NewAsync(d duration.NanoSecond, callback timer_p.TimerListener) (t *Async, err error_p.Error) {
	var timer Async
	timer.Init(callback)
	err = timer.Start(d)
	t = &timer
	return
}

type Async struct {
	// Timer wakes up at when, and then at when+period, ... (period > 0 only)
	// when must be positive on an active timer.
	when   monotonic.Time
	period duration.NanoSecond

	// The status field holds one of the values in status file.
	status timer.Status

	// callback function that call when reach
	// it is possible that callback will be called a little after the delay.
	// * NOTE: each time calling callback() in the timer goroutine, so callback must be
	// * a well-behaved function and not block.
	callback timer_p.TimerListener

	timers *TimingWheel
}

// Init initialize the timer with given callback function or make the channel and send signal on it
// Be aware that given function must not be closure and must not block the caller.
//
//memar:impl memar/timer/protocol.Timer
func (t *Async) Init(callback timer_p.TimerListener) (err error_p.Error) {
	if t.callback != nil {
		err = &timer_errs.TimerAlreadyInit
		return
	}

	t.callback = callback
	return
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (t *Async) Reinit(callback timer_p.TimerListener) (err error_p.Error) {
	// var status = t.status.Load()
	// if !(status == Status_Unset || status == Status_Deleted) {
	// 	panic("timer: Reinit called with non stopped timer")
	// }
	err = t.Stop()
	if err != nil {
		return
	}
	t.callback = callback
	t.timers = nil
	return
}
func (t *Async) Deinit() (err error_p.Error) {
	err = t.Stop()
	return
}

//memar:impl memar/timer/protocol.Timer
func (t *Async) Status() (activeStatus timer.Status) { return t.status.Load() }

// Start adds the timer to the running cpu core timing.
// This should only be called with a newly created timer.
// That avoids the risk of changing the when field of a timer in some P's heap,
// which could cause the heap to become unsorted.
//
//memar:impl memar/timer/protocol.Timer
func (t *Async) Start(d duration.NanoSecond) (err error_p.Error) {
	if t.callback == nil {
		err = &timer_errs.TimerNotInit
		return
	}
	if t.status != status(Status_Unset) {
		err = &timer_errs.TimerAlreadyStarted
		return
	}
	if t.timers != nil {
		panic("timer: timers already set in timer")
	}
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timers.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		panic("timer: timer must have positive duration.")
	}

	if race.DetectorEnabled {
		race.Release(t)
	}
	t.when = when(d)
	t.status = status(Status_Waiting)
	t.timers = getCoreTimingHeap()
	t.timers.AddTimer(t)
	return
}

// Stop deletes the timer t. We can't actually remove it from the timers heap.
// We can only mark it as deleted. It will be removed in due course by the timing whose heap it is on.
// Reports whether the timer was removed before it was run.
//
//memar:impl memar/timer/protocol.Timer
func (t *Async) Stop() (err error_p.Error) {
	if t.callback == nil {
		err = &timer_errs.TimerNotInit
		return
	}

	var activeStatus Status
	for {
		activeStatus = t.status.Load()
		switch activeStatus {
		case timer.Status_Waiting, timer.Status_ModifiedLater, timer.Status_ModifiedEarlier:
			if t.status.CompareAndSwap(activeStatus, timer.Status_Modifying) {
				// Must fetch t.timers before changing status,
				// as ts.cleanTimers in another goroutine can clear t.timers of a Status_Deleted timer.
				var timers = t.timers
				if !t.status.CompareAndSwap(timer.Status_Modifying, timer.Status_Deleted) {
					badTimer()
					// err =
					return
				}
				timers.deletedTimers.Add(1)
				// Timer was not yet run.
				return
			}
		case timer.Status_Deleted, timer.Status_Removing, timer.Status_Removed:
			// Timer was already run.
			return
		case timer.Status_Running, timer.Status_Moving:
			// The timer is being run or moved, by a different P Wait for it to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case timer.Status_Unset:
			err = &timer_errs.TimerNotInit
			return
		case timer.Status_Modifying:
			// Simultaneous calls to Delete() and Reset(). Wait for the other call to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		default:
			badTimer()
		}
	}
}

// Reset modifies an existing timer to new deadline.
// It's OK to call Reset() on a newly allocated Timer.
// Reports whether the timer was modified before it was run.
//
//memar:impl memar/timer/protocol.Timer
func (t *Async) Reset(d duration.NanoSecond) (err error_p.Error) {
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timers.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		panic("timer: timer must have positive duration")
	}
	if t.callback == nil {
		err = &timer_errs.TimerNotInit
		return
	}

	if race.DetectorEnabled {
		race.Release(t)
	}

	var wasRemoved = false
	var activeStatus timer.Status
loop:
	for {
		activeStatus = t.status.Load()
		switch activeStatus {
		case timer.Status_Waiting, timer.Status_ModifiedEarlier, timer.Status_ModifiedLater:
			if t.status.CompareAndSwap(activeStatus, timer.Status_Modifying) {
				break loop
			}
		case timer.Status_Unset, timer.Status_Removed:
			// Timer was already run and t is no longer in a heap.
			// Act like AddTimer.
			if t.status.CompareAndSwap(activeStatus, timer.Status_Modifying) {
				wasRemoved = true
				break loop
			}
		case timer.Status_Deleted:
			if t.status.CompareAndSwap(activeStatus, timer.Status_Modifying) {
				t.timers.deletedTimers.Add(-1)
				break loop
			}
		case timer.Status_Running, timer.Status_Removing, timer.Status_Moving:
			// The timer is being run or moved, by a different P.
			// Wait for it to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case timer.Status_Modifying:
			// Multiple simultaneous calls to Reset().
			// Wait for the other call to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		default:
			badTimer()
		}
	}

	var timerOldWhen = t.when
	var timerNewWhen = when(d)
	t.when = timerNewWhen
	if t.period != 0 {
		t.period = d
	}
	if wasRemoved {
		t.timers = getCoreTimingHeap()
		t.timers.AddTimer(t)
		if !t.status.CompareAndSwap(timer.Status_Modifying, timer.Status_Waiting) {
			badTimer()
		}
	} else {
		var newStatus = timer.Status_ModifiedLater
		if timerNewWhen < timerOldWhen {
			newStatus = timer.Status_ModifiedEarlier
			t.timers.updateTimerModifiedEarliest(timerNewWhen)
		}

		// Set the new status of the timer.
		if !t.status.CompareAndSwap(timer.Status_Modifying, newStatus) {
			badTimer()
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
//memar:impl memar/timer/protocol.Ticker
func (t *Async) Tick(first, interval duration.NanoSecond) (err error_p.Error) {
	if first < 1 || interval < 1 {
		panic("timer: non-positive interval to tick. period must be non-negative,")
	}
	t.period = interval
	t.Start(first)
	return
}

//memar:impl memar/timer/protocol.Stringer
func (t *Async) ToString() string                        { return "TODO:::" }
func (t *Async) FromString(s string) (err error_p.Error) { return }
