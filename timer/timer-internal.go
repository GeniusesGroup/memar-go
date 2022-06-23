/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"runtime/internal/atomic"

	"../protocol"
	"../time/monotonic"
)

// when is a helper function for setting the 'when' field of a Timer.
// It returns what the time will be, in nanoseconds, Duration d in the future.
// If d is negative, it is ignored. If the returned value would be less than
// zero because of an overflow, MaxInt64 is returned.
func when(d protocol.Duration) (t int64) {
	t = monotonic.RuntimeNano()
	if d <= 0 {
		return
	}
	t += int64(d)
	if t < 0 {
		// N.B. monotonic.RuntimeNano() and d are always positive, so addition
		// (including overflow) will never result in t == 0.
		t = 1<<63 - 1 // math.MaxInt64
	}
	return
}

type goFunc func()

func (g goFunc) concurrentRun() {
	go g()
}

// NotifyChannel does a non-blocking send the signal on t.signal
func (t *Timer) notifyChannel() {
	select {
	case t.signal <- struct{}{}:
	default:
	}
}

func (t *Timer) checkAndPanicInStart() {
	t.checkAndPanic()
	if t.status != status_Unset {
		panic("timer: start called with initialized timer")
	}
}

func (t *Timer) checkAndPanicInModify() {
	t.checkAndPanic()
}

func (t *Timer) checkAndPanicInStop() {
	if t.function == nil {
		panic("timer: Stop called on uninitialized Timer")
	}
}

func (t *Timer) checkAndPanic() {
	// when must be positive. A negative value will cause runtimer to
	// overflow during its delta calculation and never expire other runtime
	// timers. Zero will cause checkTimers to fail to notice the timer.
	if t.when <= 0 {
		panic("timer: timer must have positive duration")
	}
	if t.period < 0 {
		panic("timer: period must be non-negative")
	}
	if t.function == nil {
		panic("timer: Timer must initialized before start or reset")
	}
}

// add adds a timer to the current P.
// This should only be called with a newly created timer.
// That avoids the risk of changing the when field of a timer in some P's heap,
// which could cause the heap to become unsorted.
func (t *Timer) add() {
	t.status = status_Waiting

	// Disable preemption while using pp to avoid changing another P's heap.
	mp := acquirem()

	pp := getg().m.p.ptr()
	lock(&pp.timersLock)
	cleantimers(pp)
	doaddtimer(pp, t)
	unlock(&pp.timersLock)

	wakeNetPoller(t.when)

	releasem(mp)
}

// delete deletes the timer t. It may be on some other P, so we can't
// actually remove it from the timers heap. We can only mark it as deleted.
// It will be removed in due course by the P whose heap it is on.
// Reports whether the timer was removed before it was run.
func (t *Timer) delete() bool {
	for {
		var status = t.status.Load()
		switch status {
		case status_Waiting, status_ModifiedLater:
			// Prevent preemption while the timer is in status_Modifying.
			// This could lead to a self-deadlock. See #38070.
			mp := acquirem()
			if t.status.CompareAndSwap(status, status_Modifying) {
				// Must fetch t.pp before changing status,
				// as cleantimers in another goroutine
				// can clear t.pp of a status_Deleted timer.
				tpp := t.pp.ptr()
				if !t.status.CompareAndSwap(status_Modifying, status_Deleted) {
					badTimer()
				}
				releasem(mp)
				atomic.Xadd(&tpp.deletedTimers, 1)
				// Timer was not yet run.
				return true
			} else {
				releasem(mp)
			}
		case status_ModifiedEarlier:
			// Prevent preemption while the timer is in status_Modifying.
			// This could lead to a self-deadlock. See #38070.
			mp := acquirem()
			if t.status.CompareAndSwap(status, status_Modifying) {
				// Must fetch t.pp before setting status
				// to status_Deleted.
				tpp := t.pp.ptr()
				if !t.status.CompareAndSwap(status_Modifying, status_Deleted) {
					badTimer()
				}
				releasem(mp)
				atomic.Xadd(&tpp.deletedTimers, 1)
				// Timer was not yet run.
				return true
			} else {
				releasem(mp)
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
			// Simultaneous calls to deltimer and modtimer.
			// Wait for the other call to complete.
			osyield()
		default:
			badTimer()
		}
	}
}

// modify modifies an existing timer.
// This is called by the netpoll code or time.Ticker.Reset or time.Timer.Reset.
// Reports whether the timer was modified before it was run.
func (t *Timer) modify() bool {
	wasRemoved := false
	var pending bool
	var mp *m
loop:
	for {
		var status = t.status.Load()
		switch status {
		case status_Waiting, status_ModifiedEarlier, status_ModifiedLater:
			// Prevent preemption while the timer is in status_Modifying.
			// This could lead to a self-deadlock. See #38070.
			mp = acquirem()
			if t.status.CompareAndSwap(status, status_Modifying) {
				pending = true // timer not yet run
				break loop
			}
			releasem(mp)
		case status_Unset, status_Removed:
			// Prevent preemption while the timer is in status_Modifying.
			// This could lead to a self-deadlock. See #38070.
			mp = acquirem()

			// Timer was already run and t is no longer in a heap.
			// Act like addtimer.
			if t.status.CompareAndSwap(status, status_Modifying) {
				wasRemoved = true
				pending = false // timer already run or stopped
				break loop
			}
			releasem(mp)
		case status_Deleted:
			// Prevent preemption while the timer is in status_Modifying.
			// This could lead to a self-deadlock. See #38070.
			mp = acquirem()
			if t.status.CompareAndSwap(status, status_Modifying) {
				atomic.Xadd(&t.pp.ptr().deletedTimers, -1)
				pending = false // timer already stopped
				break loop
			}
			releasem(mp)
		case status_Running, status_Removing, status_Moving:
			// The timer is being run or moved, by a different P.
			// Wait for it to complete.
			osyield()
		case status_Modifying:
			// Multiple simultaneous calls to modtimer.
			// Wait for the other call to complete.
			osyield()
		default:
			badTimer()
		}
	}

	if wasRemoved {
		t.when = when
		pp := getg().m.p.ptr()
		lock(&pp.timersLock)
		doaddtimer(pp, t)
		unlock(&pp.timersLock)
		if !t.status.CompareAndSwap(status_Modifying, status_Waiting) {
			badTimer()
		}
		releasem(mp)
		wakeNetPoller(when)
	} else {
		// The timer is in some other P's heap, so we can't change
		// the when field. If we did, the other P's heap would
		// be out of order. So we put the new when value in the
		// nextwhen field, and let the other P set the when field
		// when it is prepared to resort the heap.
		t.nextwhen = when

		newStatus := status_ModifiedLater
		if when < t.when {
			newStatus = status_ModifiedEarlier
		}

		tpp := t.pp.ptr()

		if newStatus == status_ModifiedEarlier {
			updateTimerModifiedEarliest(tpp, when)
		}

		// Set the new status of the timer.
		if !t.status.CompareAndSwap(status_Modifying, newStatus) {
			badTimer()
		}
		releasem(mp)

		// If the new status is earlier, wake up the poller.
		if newStatus == status_ModifiedEarlier {
			wakeNetPoller(when)
		}
	}

	return pending
}
