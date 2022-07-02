/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"../cpu"
	"../race"
	"../time/monotonic"
)

// TimingHeap ...
// Active timers live in the timers field as heap structure.
// Inactive timers live there too temporarily, until they are removed.
// Due to atomic function need memory alignment, Don't change fields order.
//
// https://github.com/search?l=go&q=timer&type=Repositories
// https://github.com/RussellLuo/timingwheel/blob/master/delayqueue/delayqueue.go
type TimingHeap struct {
	coreID uint64 // CPU core number this heap run on it

	// The when field of the first entry on the timer heap.
	// This is updated using atomic functions.
	// This is 0 if the timer heap is empty.
	timer0When int64

	// The earliest known when field of a timer with
	// timerModifiedEarlier status. Because the timer may have been
	// modified again, there need not be any timer with this value.
	// This is updated using atomic functions.
	// This is 0 if there are no timerModifiedEarlier timers.
	timerModifiedEarliest int64

	// Number of timers in P's heap.
	// Modified using atomic instructions.
	numTimers int32

	// Number of timerDeleted timers in P's heap.
	// Modified using atomic instructions.
	deletedTimers int32

	// Race context used while executing timer functions.
	timerRaceCtx uintptr

	// Lock for timers. We normally access the timers while running
	// on this P, but the scheduler can also do it from a different P.
	timersLock sync.Mutex
	// Must hold timersLock to access.
	// https://en.wikipedia.org/wiki/Heap_(data_structure)#Comparison_of_theoretic_bounds_for_variants
	// Balancing a heap is done by th.siftUp or th.siftDown methods
	timers []timerBucket
}

type timerBucket struct {
	timer *Timer
	// Two reason to have timer when here:
	// - hot cache to prevent dereference timer to get when field
	// - It can be difference with timer when filed in timerModifiedXX status.
	when int64
}

func (th *TimingHeap) Init() {
	// TODO::: let application flow choose timers init cap or force it?
	// th.timers = make([]timerBucket, 1024)
	th.coreID = cpu.ActiveCoreID()
}

// addTimer adds t to the timers queue.
// The caller must have locked the th.timersLock
func (th *TimingHeap) addTimer(t *Timer) {
	th.timersLock.Lock()

	th.cleanTimers()

	var timerWhen = t.when
	t.timers = th
	var i = len(th.timers)
	th.timers = append(th.timers, timerBucket{t, timerWhen})

	th.siftUpTimer(i)
	if t == th.timers[0].timer {
		atomic.StoreInt64(&th.timer0When, timerWhen)
	}
	atomic.AddInt32(&th.numTimers, 1)

	th.timersLock.Unlock()
}

// deleteTimer removes timer i from the timers heap.
// It returns the smallest changed index in th.timers
// The caller must have locked the th.timersLock
func (th *TimingHeap) deleteTimer(i int) int {
	th.timers[i].timer.timers = nil

	var last = len(th.timers) - 1
	if i != last {
		th.timers[i] = th.timers[last]
	}
	th.timers[last].timer = nil
	th.timers = th.timers[:last]

	var smallestChanged = i
	if i != last {
		// Moving to i may have moved the last timer to a new parent,
		// so sift up to preserve the heap guarantee.
		smallestChanged = th.siftUpTimer(i)
		th.siftDownTimer(i)
	}
	if i == 0 {
		th.updateTimer0When()
	}
	atomic.AddInt32(&th.numTimers, -1)
	return smallestChanged
}

// deleteTimer0 removes timer 0 from the timers heap.
// It reports whether it saw no problems due to races.
// The caller must have locked the th.timersLock
func (th *TimingHeap) deleteTimer0() {
	th.timers[0].timer.timers = nil

	var last = len(th.timers) - 1
	if last > 0 {
		th.timers[0] = th.timers[last]
	}
	th.timers[last].timer = nil
	th.timers = th.timers[:last]
	if last > 0 {
		th.siftDownTimer(0)
	}
	th.updateTimer0When()
	atomic.AddInt32(&th.numTimers, -1)
}

// cleanTimers cleans up the head of the timer queue. This speeds up
// programs that create and delete timers; leaving them in the heap
// slows down addTimer. Reports whether no timer problems were found.
// The caller must have locked the th.timersLock
func (th *TimingHeap) cleanTimers() {
	if len(th.timers) == 0 {
		return
	}

	for {
		// This loop can theoretically run for a while, and because
		// it is holding timersLock it cannot be preempted.
		// If someone is trying to preempt us, just return.
		// We can clean the timers later.
		// if gp.preemptStop {
		// 	return
		// }

		var timerBucket = th.timers[0]
		var timer = timerBucket.timer
		var status = timer.status.Load()
		switch status {
		case status_Deleted:
			if !timer.status.CompareAndSwap(status, status_Removing) {
				continue
			}
			th.deleteTimer0()
			if !timer.status.CompareAndSwap(status_Removing, status_Removed) {
				badTimer()
			}
			atomic.AddInt32(&th.deletedTimers, -1)
		case status_ModifiedEarlier, status_ModifiedLater:
			if !timer.status.CompareAndSwap(status, status_Moving) {
				continue
			}
			// Now we can change the when field of timerBucket.
			th.timers[0].when = timer.when
			// Move timer to the right position.
			th.deleteTimer0()
			th.addTimer(timer)
			if !timer.status.CompareAndSwap(status_Moving, status_Waiting) {
				badTimer()
			}
		default:
			// Head of timers does not need adjustment.
			return
		}
	}
}

// moveTimers moves a slice of timers to the timers heap.
// The slice has been taken from a different Timers.
// This is currently called when the world is stopped, but the caller
// is expected to have locked the th.timersLock
func (th *TimingHeap) moveTimers(timers []timerBucket) {
	for _, timerBucket := range timers {
		var timer = timerBucket.timer
	loop:
		for {
			var status = timer.status.Load()
			switch status {
			case status_Waiting, status_ModifiedEarlier, status_ModifiedLater:
				if !timer.status.CompareAndSwap(status, status_Moving) {
					continue
				}
				timer.timers = nil
				th.addTimer(timer)
				if !timer.status.CompareAndSwap(status_Moving, status_Waiting) {
					badTimer()
				}
				break loop
			case status_Deleted:
				if !timer.status.CompareAndSwap(status, status_Removed) {
					continue
				}
				timer.timers = nil
				// We no longer need this timer in the heap.
				break loop
			case status_Modifying:
				// Loop until the modification is complete.
				osyield()
			case status_Unset, status_Removed:
				// We should not see these status values in a timers heap.
				badTimer()
			case status_Running, status_Removing, status_Moving:
				// Some other P thinks it owns this timer,
				// which should not happen.
				badTimer()
			default:
				badTimer()
			}
		}
	}
}

// adjustTimers looks through the timers for any timers that have been modified to run earlier,
// and puts them in the correct place in the heap. While looking for those timers,
// it also moves timers that have been modified to run later, and removes deleted timers.
// The caller must have locked the th.timersLock
func (th *TimingHeap) adjustTimers(now int64) {
	// If we haven't yet reached the time of the first status_ModifiedEarlier
	// timer, don't do anything. This speeds up programs that adjust
	// a lot of timers back and forth if the timers rarely expire.
	// We'll postpone looking through all the adjusted timers until
	// one would actually expire.
	var first = atomic.LoadInt64(&th.timerModifiedEarliest)
	if first == 0 || int64(first) > now {
		if verifyTimers {
			th.verifyTimerHeap()
		}
		return
	}

	// We are going to clear all status_ModifiedEarlier timers.
	atomic.StoreInt64(&th.timerModifiedEarliest, 0)

	var moved []*Timer
	var timers = th.timers
	var timersLen = len(timers)
	for i := 0; i < timersLen; i++ {
		var timerBucket = timers[i]
		var timer = timerBucket.timer
		var status = timer.status.Load()
		switch status {
		case status_Deleted:
			if timer.status.CompareAndSwap(status, status_Removing) {
				var changed = th.deleteTimer(i)
				if !timer.status.CompareAndSwap(status_Removing, status_Removed) {
					badTimer()
				}
				atomic.AddInt32(&th.deletedTimers, -1)
				// Go back to the earliest changed heap entry.
				// "- 1" because the loop will add 1.
				i = changed - 1
			}
		case status_ModifiedEarlier, status_ModifiedLater:
			if timer.status.CompareAndSwap(status, status_Moving) {
				// Take t off the heap, and hold onto it.
				// We don't add it back yet because the
				// heap manipulation could cause our
				// loop to skip some other timer.
				var changed = th.deleteTimer(i)
				moved = append(moved, timer)
				// Go back to the earliest changed heap entry.
				// "- 1" because the loop will add 1.
				i = changed - 1
			}
		case status_Unset, status_Running, status_Removing, status_Removed, status_Moving:
			badTimer()
		case status_Waiting:
			// OK, nothing to do.
		case status_Modifying:
			// Check again after modification is complete.
			osyield()
			i--
		default:
			badTimer()
		}
	}

	if len(moved) > 0 {
		th.addAdjustedTimers(moved)
	}

	if verifyTimers {
		th.verifyTimerHeap()
	}
}

// addAdjustedTimers adds any timers we adjusted in th.adjustTimers
// back to the timer heap.
func (th *TimingHeap) addAdjustedTimers(moved []*Timer) {
	for _, t := range moved {
		th.addTimer(t)
		if !t.status.CompareAndSwap(status_Moving, status_Waiting) {
			badTimer()
		}
	}
}

// runTimer examines the first timer in timers. If it is ready based on now,
// it runs the timer and removes or updates it.
// Returns 0 if it ran a timer, -1 if there are no more timers, or the time
// when the first timer should run.
// The caller must have locked the th.timersLock
// If a timer is run, this will temporarily unlock the timers.
func (th *TimingHeap) runTimer(now int64) int64 {
	for {
		var timerBucket = th.timers[0]
		var timer = timerBucket.timer
		var status = timer.status.Load()
		switch status {
		case status_Waiting:
			if timer.when > now {
				// Not ready to run.
				return timer.when
			}

			if !timer.status.CompareAndSwap(status, status_Running) {
				continue
			}
			// Note that runOneTimer may temporarily unlock th.timersLock
			th.runOneTimer(timer, now)
			return 0

		case status_Deleted:
			if !timer.status.CompareAndSwap(status, status_Removing) {
				continue
			}
			th.deleteTimer0()
			if !timer.status.CompareAndSwap(status_Removing, status_Removed) {
				badTimer()
			}
			atomic.AddInt32(&th.deletedTimers, -1)
			if len(th.timers) == 0 {
				return -1
			}

		case status_ModifiedEarlier, status_ModifiedLater:
			if !timer.status.CompareAndSwap(status, status_Moving) {
				continue
			}
			th.deleteTimer0()
			th.addTimer(timer)
			if !timer.status.CompareAndSwap(status_Moving, status_Waiting) {
				badTimer()
			}

		case status_Modifying:
			// Wait for modification to complete.
			osyield()
		case status_Unset, status_Removed:
			// Should not see a new or inactive timer on the heap.
			badTimer()
		case status_Running, status_Removing, status_Moving:
			// These should only be set when timers are locked,
			// and we didn't do it.
			badTimer()
		default:
			badTimer()
		}
	}
}

// runOneTimer runs a single timer.
// The caller must have locked the th.timersLock
// This will temporarily unlock the timers while running the timer function.
func (th *TimingHeap) runOneTimer(t *Timer, now int64) {
	if race.DetectorEnabled {
		ppcur := getg().m.p.ptr()
		if ppcur.timerRaceCtx == 0 {
			ppcur.timerRaceCtx = racegostart(abi.FuncPCABIInternal(runtimer) + sys.PCQuantum)
		}
		race.Acquirectx(ppcur.timerRaceCtx, unsafe.Pointer(t))
	}

	if t.period > 0 && t.periodNumber != 0 {
		// Leave in heap but adjust next time to fire.
		var delta = t.when - now
		t.when += t.period * (1 + -delta/t.period)
		if t.when < 0 { // check for overflow.
			t.when = maxWhen
		}
		th.siftDownTimer(0)
		if !t.status.CompareAndSwap(status_Running, status_Waiting) {
			badTimer()
		}
		th.updateTimer0When()
		if t.periodNumber > 0 {
			t.periodNumber--
		}
	} else {
		// Remove from heap.
		th.deleteTimer0()
		if !t.status.CompareAndSwap(status_Running, status_Unset) {
			badTimer()
		}
	}

	if race.DetectorEnabled {
		// Temporarily use the current P's racectx for g0.
		var gp = getg()
		if gp.racectx != 0 {
			panic("timer - runOneTimer: unexpected racectx")
		}
		gp.racectx = gp.m.p.ptr().timerRaceCtx
	}

	var callback = t.callback
	var arg = t.arg
	th.timersLock.Unlock()
	callback(arg)
	th.timersLock.Lock()

	if race.DetectorEnabled {
		var gp = getg()
		gp.racectx = 0
	}
}

// clearDeletedTimers removes all deleted timers from the timers heap.
// This is used to avoid clogging up the heap if the program
// starts a lot of long-running timers and then stops them.
// For example, this can happen via context.WithTimeout.
//
// This is the only function that walks through the entire timer heap,
// other than moveTimers which only runs when the world is stopped.
//
// The caller must have locked the th.timersLock
func (th *TimingHeap) clearDeletedTimers() {
	// We are going to clear all status_ModifiedEarlier timers.
	// Do this now in case new ones show up while we are looping.
	atomic.StoreInt64(&th.timerModifiedEarliest, 0)

	var cdel = int32(0)
	var to = 0
	var changedHeap = false
	var timers = th.timers
	var timersLen = len(timers)
nextTimer:
	for i := 0; i < timersLen; i++ {
		var timerBucket = timers[i]
		var timer = timerBucket.timer
		for {
			var status = timer.status.Load()
			switch status {
			case status_Waiting:
				if changedHeap {
					timers[to] = timerBucket
					th.siftUpTimer(to)
				}
				to++
				continue nextTimer
			case status_ModifiedEarlier, status_ModifiedLater:
				if timer.status.CompareAndSwap(status, status_Moving) {
					timerBucket.when = timer.when
					timers[to] = timerBucket
					th.siftUpTimer(to)
					to++
					changedHeap = true
					if !timer.status.CompareAndSwap(status_Moving, status_Waiting) {
						badTimer()
					}
					continue nextTimer
				}
			case status_Deleted:
				if timer.status.CompareAndSwap(status, status_Removing) {
					timer.timers = nil
					cdel++
					if !timer.status.CompareAndSwap(status_Removing, status_Removed) {
						badTimer()
					}
					changedHeap = true
					continue nextTimer
				}
			case status_Modifying:
				// Loop until modification complete.
				osyield()
			case status_Unset, status_Removed:
				// We should not see these status values in a timer heap.
				badTimer()
			case status_Running, status_Removing, status_Moving:
				// Some other P thinks it owns this timer,
				// which should not happen.
				badTimer()
			default:
				badTimer()
			}
		}
	}

	// Set remaining slots in timers slice to nil,
	// so that the timer values can be garbage collected.
	for i := to; i < len(timers); i++ {
		timers[i].timer = nil
	}

	atomic.AddInt32(&th.deletedTimers, -cdel)
	atomic.AddInt32(&th.numTimers, -cdel)

	timers = timers[:to]
	th.timers = timers
	th.updateTimer0When()

	if verifyTimers {
		th.verifyTimerHeap()
	}
}

// verifyTimerHeap verifies that the timer heap is in a valid state.
// This is only for debugging, and is only called if verifyTimers is true.
// The caller must have locked the th.timersLock
func (th *TimingHeap) verifyTimerHeap() {
	var timers = th.timers
	var timersLen = len(timers)
	// First timer has no parent, so i must be start from 1.
	for i := 1; i < timersLen; i++ {
		var timerBucket = th.timers[0]
		var timer = timerBucket.timer

		// The heap is 4-ary. See siftUpTimer and siftDownTimer.
		var p = (i - 1) / 4
		if timer.when < timers[p].when {
			print("timer: bad timer heap at ", i, ": ", p, ": ", th.timers[p].when, ", ", i, ": ", timer.when, "\n")
			panic("timer: bad timer heap")
		}
	}
	var numTimers = int(atomic.LoadInt32(&th.numTimers))
	if timersLen != numTimers {
		println("timer: heap len", len(th.timers), "!= numTimers", numTimers)
		panic("timer: bad timer heap len")
	}
}

// updateTimer0When sets the timer0When field by check first timer in queue.
// The caller must have locked the th.timersLock
func (th *TimingHeap) updateTimer0When() {
	if len(th.timers) == 0 {
		atomic.StoreInt64(&th.timer0When, 0)
	} else {
		atomic.StoreInt64(&th.timer0When, th.timers[0].when)
	}
}

// updateTimerModifiedEarliest updates the th.timerModifiedEarliest value.
// The timers will not be locked.
func (th *TimingHeap) updateTimerModifiedEarliest(nextWhen int64) {
	for {
		var old = atomic.LoadInt64(&th.timerModifiedEarliest)
		if old != 0 && int64(old) < nextWhen {
			return
		}
		if atomic.CompareAndSwapInt64(&th.timerModifiedEarliest, old, nextWhen) {
			return
		}
	}
}

// sleepUntil returns the time when the next timer should fire.
func (th *TimingHeap) sleepUntil() (until int64) {
	until = int64(maxWhen)

	var timer0When = atomic.LoadInt64(&th.timer0When)
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}

	timer0When = atomic.LoadInt64(&th.timerModifiedEarliest)
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}
	return
}

// noBarrierWakeTime looks at timers and returns the time when we should wake up.
// This function is invoked when dropping a Timers, and must run without any write barriers.
// Unlike th.sleepUntil(), It returns 0 if there are no timers.
func (th *TimingHeap) noBarrierWakeTime() (until int64) {
	until = atomic.LoadInt64(&th.timer0When)
	var nextAdj = atomic.LoadInt64(&th.timerModifiedEarliest)
	if until == 0 || (nextAdj != 0 && nextAdj < until) {
		until = nextAdj
	}
	return
}

// checkTimers runs any timers for the P that are ready.
// If now is not 0 it is the current time.
// It returns the passed time or the current time if now was passed as 0.
// and the time when the next timer should run or 0 if there is no next timer,
// and reports whether it ran any timers.
// If the time when the next timer should run is not 0,
// it is always larger than the returned time.
// We pass now in and out to avoid extra calls of monotonic.RuntimeNano().
func (th *TimingHeap) checkTimers(now int64) (rnow, pollUntil int64, ran bool) {
	// If it's not yet time for the first timer, or the first adjusted
	// timer, then there is nothing to do.
	var next = th.noBarrierWakeTime()
	if next == 0 {
		// No timers to run or adjust.
		return now, 0, false
	}

	if now == 0 {
		now = monotonic.RuntimeNano()
	}
	if now < next {
		// Next timer is not ready to run, but keep going
		// if we would clear deleted timers.
		// This corresponds to the condition below where
		// we decide whether to call clearDeletedTimers.
		if atomic.LoadInt32(&th.deletedTimers) <= atomic.LoadInt32(&th.numTimers)/4 {
			return now, next, false
		}
	}

	th.timersLock.Lock()

	if len(th.timers) > 0 {
		th.adjustTimers(now)
		for len(th.timers) > 0 {
			// Note that th.runTimer may temporarily unlock th.timersLock.
			var tw = th.runTimer(now)
			if tw != 0 {
				if tw > 0 {
					pollUntil = tw
				}
				break
			}
			ran = true
		}
	}

	// If this is the local P, and there are a lot of deleted timers,
	// clear them out. We only do this for the local P to reduce
	// lock contention on timersLock.
	if int(atomic.LoadInt32(&th.deletedTimers)) > len(th.timers)/4 {
		th.clearDeletedTimers()
	}

	th.timersLock.Unlock()

	return now, pollUntil, ran
}

// Check for deadlock situation
func (th *TimingHeap) checkDead() {
	// Maybe jump time forward for playground.
	if faketime != 0 {
		var when = th.sleepUntil()

		faketime = when

		var mp = mget()
		if mp == nil {
			// There should always be a free M since
			// nothing is running.
			throw("timers - checkDead: no m for timer")
		}
		return
	}

	// There are no goroutines running, so we can look at the P's.
	if len(th.timers) > 0 {
		return
	}
}

// badTimer is called if the timer data structures have been corrupted,
// presumably due to racy use by the program. We panic here rather than
// panicing due to invalid slice access while holding locks.
// See issue #25686.
func badTimer() {
	panic("timers: data corruption")
}

// Heap maintenance algorithms.
// These algorithms check for slice index errors manually.
// Slice index error can happen if the program is using racy
// access to timers. We don't want to panic here, because
// it will cause the program to crash with a mysterious
// "panic holding locks" message. Instead, we panic while not
// holding a lock.

// siftUpTimer puts the timer at position i in the right place
// in the heap by moving it up toward the top of the heap.
// It returns the smallest changed index.
func (th *TimingHeap) siftUpTimer(i int) int {
	var timers = th.timers
	var timerWhen = timers[i].when

	var tmp = timers[i]
	for i > 0 {
		var p = (i - 1) / 4 // parent
		if timerWhen >= timers[p].when {
			break
		}
		timers[i] = timers[p]
		i = p
	}
	if tmp != timers[i] {
		timers[i] = tmp
	}
	return i
}

// siftDownTimer puts the timer at position i in the right place
// in the heap by moving it down toward the bottom of the heap.
func (th *TimingHeap) siftDownTimer(i int) {
	var timers = th.timers
	var timersLen = len(timers)
	var timerWhen = timers[i].when

	var tmp = timers[i]
	for {
		var c = i*4 + 1 // left child
		var c3 = c + 2  // mid child
		if c >= timersLen {
			break
		}
		var w = timers[c].when
		if c+1 < timersLen && timers[c+1].when < w {
			w = timers[c+1].when
			c++
		}
		if c3 < timersLen {
			var w3 = timers[c3].when
			if c3+1 < timersLen && timers[c3+1].when < w3 {
				w3 = timers[c3+1].when
				c3++
			}
			if w3 < w {
				w = w3
				c = c3
			}
		}
		if w >= timerWhen {
			break
		}
		timers[i] = timers[c]
		i = c
	}
	if tmp != timers[i] {
		timers[i] = tmp
	}
}
