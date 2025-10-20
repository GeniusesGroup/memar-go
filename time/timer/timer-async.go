/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	error_p "memar/process/error/protocol"
	"memar/process/race"
	"memar/computer/runtime/scheduler"
	"memar/time/duration"
	"memar/time/monotonic"
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
//memar:impl memar/time/timer/protocol.Timer
func (self *Async) Init(callback timer_p.TimerListener) (err error_p.Error) {
	if self.callback != nil {
		err = &timer_errs.ErrTimerAlreadyInit
		return
	}

	self.callback = callback
	return
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Async) Reinit(callback timer_p.TimerListener) (err error_p.Error) {
	// var status = self.status.Load()
	// if !(status == Status_Unset || status == Status_Deleted) {
	// 	panic("timer: Reinit called with non stopped timer")
	// }
	err = self.Stop()
	if err != nil {
		return
	}
	self.callback = callback
	self.timing = nil
	return
}
func (self *Async) Deinit() (err error_p.Error) {
	err = self.Stop()
	// TODO::: Can we remove t from related timing heap?
	return
}

//memar:impl memar/time/timer/protocol.Timer
func (self *Async) Status() (activeStatus Status) { return self.status.Load() }
func (self *Async) When() monotonic.Time          { return self.when }

// Start adds the timer to the running cpu core timing.
// This should only be called with a newly created timer.
// That avoids the risk of changing the when field of a timer in some P's heap,
// which could cause the heap to become unsorted.
//
//memar:impl memar/time/timer/protocol.Timer
func (self *Async) Start(d duration.NanoSecond) (err error_p.Error) {
	if self.callback == nil {
		err = &timer_errs.ErrTimerNotInit
		return
	}
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timing.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		err = &timer_errs.ErrNegativeDuration
		return
	}
	var activeStatus = self.status.Load()
	if activeStatus != Status_Unset || self.timing != nil {
		err = &timer_errs.ErrTimerAlreadyStarted
		return
	}

	if !self.status.CompareAndSwap(Status_Unset, Status_Waiting) {
		err = &timer_errs.ErrTimerRacyAccess
		return
	}

	if race.DetectorEnabled {
		race.Release(self)
	}

	self.when = when(d)
	self.timing = getActiveTiming()
	self.timing.AddTimer(self)
	return
}

// Stop deletes the timer. We can't actually remove it from the timing heap.
// We can only mark it as deleted. It will be removed in due course by the timing whose heap it is on.
// Reports whether the timer was removed before it was run.
//
//memar:impl memar/time/timer/protocol.Timer
func (self *Async) Stop() (err error_p.Error) {
	if self.callback == nil {
		err = &timer_errs.ErrTimerNotInit
		return
	}

	var activeStatus Status
	for {
		activeStatus = self.status.Load()
		switch activeStatus {
		case Status_Unset:
			err = &timer_errs.ErrTimerNotInit
			return
		case Status_Waiting, Status_ModifiedLater, Status_ModifiedEarlier:
			// Must fetch self.timing before changing status,
			// due to ts.cleanTimers in another goroutine can clear self.timing of timing in Status_Deleted status.
			var timing = self.timing

			// Timer was not yet run.
			if self.status.CompareAndSwap(activeStatus, Status_Deleted) {
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
			err = &timer_errs.ErrTimerBadStatus
			return
		}
	}
}

// Reset modifies an existing timer to new deadline.
// It's OK to call Reset() on a newly allocated Timer.
// Reports whether the timer was modified before it was run.
//
//memar:impl memar/time/timer/protocol.Timer
func (self *Async) Reset(d duration.NanoSecond) (err error_p.Error) {
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timing.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		err = &timer_errs.ErrNegativeDuration
		return
	}
	if self.callback == nil {
		err = &timer_errs.ErrTimerNotInit
		return
	}

	if race.DetectorEnabled {
		race.Release(t)
	}

	var wasRemovedFromTiming = false
	var activeStatus Status
loop:
	for {
		activeStatus = self.status.Load()
		switch activeStatus {
		case Status_Waiting, Status_ModifiedEarlier, Status_ModifiedLater:
			if self.status.CompareAndSwap(activeStatus, Status_Modifying) {
				break loop
			}
		case Status_Unset, Status_Removed:
			// Timer was already run and t is no longer in a timing.
			// Act like AddTimer.
			if self.status.CompareAndSwap(activeStatus, Status_Modifying) {
				wasRemovedFromTiming = true
				break loop
			}
		case Status_Deleted:
			if self.status.CompareAndSwap(activeStatus, Status_Modifying) {
				self.timing.deletedTimersCount.Add(-1)
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
			err = &timer_errs.ErrTimerBadStatus
			return
		}
	}

	var timerOldWhen = self.when
	var timerNewWhen = when(d)
	self.when = timerNewWhen
	if self.period != 0 {
		self.period = d
	}
	if wasRemovedFromTiming {
		self.timing = getActiveTiming()
		self.timing.AddTimer(self)
		if !self.status.CompareAndSwap(Status_Modifying, Status_Waiting) {
			err = &timer_errs.ErrTimerRacyAccess
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
			self.timing.updateTimerModifiedEarliest(timerNewWhen)
		}

		// Set the new status of the timer.
		if !self.status.CompareAndSwap(Status_Modifying, newStatus) {
			err = &timer_errs.ErrTimerRacyAccess
			// TODO::: Easily just return??
			return
		}
	}

	return
}

// Tick will call the self.callback.TimerHandler() after each tick on initialized Timer.
// The period of the ticks is specified by the duration arguments.
// The ticker will adjust the time interval or drop ticks to make up for slow receivers.
// The durations must be greater than zero; if not, Tick() will panic.
// Stop the ticker to release associated resources.
//
//memar:impl memar/time/timer/protocol.Ticker
func (self *Async) Tick(first, interval duration.NanoSecond) (err error_p.Error) {
	if first < 1 || interval < 1 {
		err = &timer_errs.ErrNegativeDuration
		return
	}
	self.period = interval
	err = self.Start(first)
	return
}

//memar:impl memar/codec/string/protocol.Stringer
func (self *Async) ToString() (str string, err error_p.Error) {
	var until = self.when.UntilNow()
	var untilSecond, untilNanoSecond = until.ToSecAndNano()
	var untilSecondString, untilNanoSecondString string
	untilSecondString, err = untilSecond.ToString()
	untilNanoSecondString, err = untilNanoSecond.ToString()
	str = "Timer awake after " + untilSecondString + " seconds and" + untilNanoSecondString + "nano-second"
	return
}
func (self *Async) FromString(str string) (err error_p.Error) { return }
