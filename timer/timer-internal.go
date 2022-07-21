/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"sync/atomic"
	"unsafe"

	"../cpu"
	"../protocol"
	"../race"
)

type timer struct {
	// Timer wakes up at when, and then at when+period, ... (period > 0 only)
	// when must be positive on an active timer.
	when         int64
	period       int64
	periodNumber int64 // -1 means no limit

	// The status field holds one of the values in status file.
	status status

	// callback function that call when reach
	// it is possible that callback will be called a little after the delay.
	// **NOTE**: each time calling callback() in the timer goroutine, so callback must be
	// a well-behaved function and not block.
	callback protocol.TimerListener

	timers *Timers
}

// use to prevent memory leak
func (t *timer) reset() {
	t.callback = nil
	t.timers = nil
}

// add adds a timer to the running cpu core timers.
// This should only be called with a newly created timer.
// That avoids the risk of changing the when field of a timer in some P's heap,
// which could cause the heap to become unsorted.
func (t *timer) add(d protocol.Duration) {
	if t.callback == nil {
		panic("timer: Timer must initialized before start")
	}
	if t.status != status_Unset {
		panic("timer: start called with initialized timer")
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
		race.Release(unsafe.Pointer(t))
	}
	t.when = when(d)
	t.status = status_Waiting
	t.timers = &poolByCores[cpu.ActiveCoreID()]
	t.timers.AddTimer(t)
}

// delete deletes the timer t. It may be on some other P, so we can't
// actually remove it from the timers heap. We can only mark it as deleted.
// It will be removed in due course by the P whose heap it is on.
// Reports whether the timer was removed before it was run.
func (t *timer) delete() bool {
	if t.callback == nil {
		panic("timer: Stop called on uninitialized Timer")
	}

	for {
		var status = t.status.Load()
		switch status {
		case status_Waiting, status_ModifiedLater:
			if t.status.CompareAndSwap(status, status_Modifying) {
				// Must fetch t.timers before changing status,
				// as ts.cleanTimers in another goroutine can clear t.timers of a status_Deleted timer.
				var timers = t.timers
				if !t.status.CompareAndSwap(status_Modifying, status_Deleted) {
					badTimer()
				}
				atomic.AddInt32(&timers.deletedTimers, 1)
				// Timer was not yet run.
				return true
			}
		case status_ModifiedEarlier:
			if t.status.CompareAndSwap(status, status_Modifying) {
				var timers = t.timers
				if !t.status.CompareAndSwap(status_Modifying, status_Deleted) {
					badTimer()
				}
				atomic.AddInt32(&timers.deletedTimers, 1)
				// Timer was not yet run.
				return true
			}
		case status_Deleted, status_Removing, status_Removed:
			// Timer was already run.
			return false
		case status_Running, status_Moving:
			// The timer is being run or moved, by a different P.
			// Wait for it to complete.
			osyield()
		case status_Unset:
			// Removing timer that was never added or
			// has already been run. Also see issue 21874.
			return false
		case status_Modifying:
			// Simultaneous calls to delete and modify.
			// Wait for the other call to complete.
			osyield()
		default:
			badTimer()
		}
	}
}

// modify modifies an existing timer.
// It's OK to call modify() on a newly allocated Timer.
// Reports whether the timer was modified before it was run.
func (t *timer) modify(d protocol.Duration) (pending bool) {
	// when must be positive. A negative value will cause ts.runTimer to
	// overflow during its delta calculation and never expire other runtime timers.
	// Zero will cause checkTimers to fail to notice the timer.
	if d < 1 {
		panic("timer: timer must have positive duration")
	}
	if t.callback == nil {
		panic("timer: Timer must initialized before reset")
	}

	if race.DetectorEnabled {
		race.Release(unsafe.Pointer(t))
	}

	var wasRemoved = false
loop:
	for {
		var status = t.status.Load()
		switch status {
		case status_Waiting, status_ModifiedEarlier, status_ModifiedLater:
			if t.status.CompareAndSwap(status, status_Modifying) {
				pending = true // timer not yet run
				break loop
			}
		case status_Unset, status_Removed:
			// Timer was already run and t is no longer in a heap.
			// Act like AddTimer.
			if t.status.CompareAndSwap(status, status_Modifying) {
				wasRemoved = true
				pending = false // timer already run or stopped
				break loop
			}
		case status_Deleted:
			if t.status.CompareAndSwap(status, status_Modifying) {
				atomic.AddInt32(&t.timers.deletedTimers, -1)
				pending = false // timer already stopped
				break loop
			}
		case status_Running, status_Removing, status_Moving:
			// The timer is being run or moved, by a different P.
			// Wait for it to complete.
			osyield()
		case status_Modifying:
			// Multiple simultaneous calls to modify.
			// Wait for the other call to complete.
			osyield()
		default:
			badTimer()
		}
	}

	var timerOldWhen = t.when
	var timerNewWhen = when(d)
	t.when = timerNewWhen
	if t.period != 0 {
		t.period = int64(d)
	}
	if wasRemoved {
		t.timers = poolByCores[cpu.ActiveCoreID()]
		t.timers.AddTimer(t)
		if !t.status.CompareAndSwap(status_Modifying, status_Waiting) {
			badTimer()
		}
	} else {
		var newStatus = status_ModifiedLater
		if timerNewWhen < timerOldWhen {
			newStatus = status_ModifiedEarlier
			t.timers.updateTimerModifiedEarliest(timerNewWhen)
		}

		// Set the new status of the timer.
		if !t.status.CompareAndSwap(status_Modifying, newStatus) {
			badTimer()
		}
	}

	return
}
