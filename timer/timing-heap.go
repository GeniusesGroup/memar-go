/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"

	"libgo/cpu"
	"libgo/protocol"
	"libgo/race"
	"libgo/scheduler"
	"libgo/time/monotonic"

	"github.com/GeniusesGroup/libgo/log"
)

// TimingHeap ...
// Active timers live in the timers field as heap structure.
// Inactive timers live there too temporarily, until they are removed.
//
// https://github.com/search?l=go&q=timer&type=Repositories
// https://github.com/RussellLuo/timingwheel/blob/master/delayqueue/delayqueue.go
type TimingHeap struct {
	coreID uint64 // CPU core number this timing run on it

	// The when field of the first entry on the timer heap.
	// This is 0 if the timer heap is empty.
	timer0When monotonic.Atomic

	// The earliest known when field of a timer with
	// timerModifiedEarlier status. Because the timer may have been
	// modified again, there need not be any timer with this value.
	// This is 0 if there are no timerModifiedEarlier timers.
	timerModifiedEarliest monotonic.Atomic

	// Number of timers in P's heap.
	numTimers atomic.Int32

	// Number of timerDeleted timers in P's heap.
	deletedTimers atomic.Int32

	// Race context used while executing timer functions.
	timerRaceCtx uintptr

	// Lock for timers. We normally access the timers while running
	// on this TimingHeap, but the scheduler can also do it from a different P.
	timersLock sync.Mutex
	// Must hold timersLock to access.
	// https://en.wikipedia.org/wiki/Heap_(data_structure)#Comparison_of_theoretic_bounds_for_variants
	// Balancing a heap is done by th.siftUp or th.siftDown methods
	timers []timerBucketHeap
}

type timerBucketHeap struct {
	timer *Async
	// Two reason to have timer when here:
	// - hot cache to prevent dereference timer to get when field
	// - It can be difference with timer when filed in timerModifiedXX status.
	when monotonic.Time
}

// Init initialize timing mechanism for th core that call the Init().
//
//libgo:impl /libgo/protocol.SoftwareLifeCycle
func (th *TimingHeap) Init() {
	// TODO::: let application flow choose timers init cap or force it?
	// th.timers = make([]timerBucketHeap, 1024)

	th.coreID = cpu.ActiveCoreID()
	// th.timerRaceCtx = racegostart(abi.FuncPCABIInternal(th.runTimer) + sys.PCQuantum)

	// TODO::: change to libgo scheduler
	go th.Start()
}

// Reinit releases all of the resources associated with timers in specific CPU core and
// move them to other core that call deinit
//
//libgo:impl /libgo/protocol.SoftwareLifeCycle
func (th *TimingHeap) Reinit() {
	var callerCoreID = cpu.ActiveCoreID()
	var newCore = &poolByCores[callerCoreID]
	th.moveTimersTo(newCore)

	th.coreID = callerCoreID
	th.timer0When.Store(0)
	th.timerModifiedEarliest.Store(0)
	th.numTimers.Store(0)
	th.deletedTimers.Store(0)
	th.timerRaceCtx = 0
	th.timers = nil
}

// Deinit releases all of the resources associated with timers in specific CPU core
//
//libgo:impl /libgo/protocol.SoftwareLifeCycle
func (th *TimingHeap) Deinit() {
	// TODO::: call all timers TimerHandler or what??
	// th.timers = nil
}

func (th *TimingHeap) Start() {
	// TODO::: Stop mechanism, new timer added mechanism
	for {
		var now = monotonic.Now()
		var nextWhen, _ = th.checkTimers(now)
		var until = nextWhen.Until(nextWhen)
		scheduler.Sleep(until)
	}
}

// MoveToMe releases all of the resources associated with timers in specific CPU core and
// move them to other core that call this method
func (th *TimingHeap) MoveToMe() {
	var callerCoreID = cpu.ActiveCoreID()
	var newCore = &poolByCores[callerCoreID]
	th.moveTimersTo(newCore)
}

// AddTimer adds t to the timers queue.
func (th *TimingHeap) AddTimer(t *Async) {
	th.timersLock.Lock()

	th.cleanTimers()

	var timerWhen = t.when
	t.timing = th
	var i = len(th.timers)
	th.timers = append(th.timers, timerBucketHeap{t, timerWhen})

	th.siftUpTimer(i)
	if t == th.timers[0].timer {
		th.timer0When.Store(timerWhen)
	}
	th.numTimers.Add(1)

	th.timersLock.Unlock()
}

// deleteTimer removes timer i from the timers heap.
// It returns the smallest changed index in th.timers
// The caller must have locked the th.timersLock
func (th *TimingHeap) deleteTimer(i int) int {
	th.timers[i].timer.timing = nil

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

	var timerRemaining = th.numTimers.Add(-1)
	if timerRemaining == 0 {
		// If there are no timers, then clearly none are modified.
		th.timerModifiedEarliest.Store(0)
	}
	return smallestChanged
}

// deleteTimer0 removes timer 0 from the timers heap.
// It reports whether it saw no problems due to races.
// The caller must have locked the th.timersLock
func (th *TimingHeap) deleteTimer0() {
	th.timers[0].timer.timing = nil

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

	var timerRemaining = th.numTimers.Add(-1)
	if timerRemaining == 0 {
		// If there are no timers, then clearly none are modified.
		th.timerModifiedEarliest.Store(0)
	}
}

// cleanTimers cleans up the head of the timer queue. This speeds up
// programs that create and delete timers; leaving them in the heap
// slows down AddTimer. Reports whether no timer problems were found.
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

		var timer = th.timers[0].timer
		var status = timer.status.Load()
		switch status {
		case protocol.TimerStatus_Deleted:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
				continue
			}
			th.deleteTimer0()
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {
				badTimer()
			}
			th.deletedTimers.Add(-1)
		case protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
				continue
			}
			// Now we can change the when field of timerBucketHeap.
			th.timers[0].when = timer.when
			// Move timer to the right position.
			th.deleteTimer0()
			th.AddTimer(timer)
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
				badTimer()
			}
		default:
			// Head of timers does not need adjustment.
			return
		}
	}
}

func (th *TimingHeap) moveTimersTo(to *TimingHeap) {
	if len(th.timers) > 0 {
		th.timersLock.Lock()

		to.timersLock.Lock()
		to.moveTimers(th.timers)
		to.timersLock.Unlock()

		th.timersLock.Unlock()
	}
}

// moveTimers moves a slice of timers to the timers heap.
// The slice has been taken from a different Timers.
// This is currently called when the world is stopped, but the caller
// is expected to have locked the th.timersLock
func (th *TimingHeap) moveTimers(timers []timerBucketHeap) {
	for _, timerBucketHeap := range timers {
		var timer = timerBucketHeap.timer
	loop:
		for {
			var status = timer.status.Load()
			switch status {
			case protocol.TimerStatus_Waiting, protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
				if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
					continue
				}
				timer.timing = nil
				th.AddTimer(timer)
				if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
					badTimer()
				}
				break loop
			case protocol.TimerStatus_Deleted:
				if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Removed) {
					continue
				}
				timer.timing = nil
				// We no longer need this timer in the heap.
				break loop
			case protocol.TimerStatus_Modifying:
				// Loop until the modification is complete.
				scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
			case protocol.TimerStatus_Unset, protocol.TimerStatus_Removed:
				// We should not see these status values in a timers heap.
				badTimer()
			case protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Moving:
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
func (th *TimingHeap) adjustTimers(now monotonic.Time) {
	// If we haven't yet reached the time of the first protocol.TimerStatus_ModifiedEarlier
	// timer, don't do anything. This speeds up programs that adjust
	// a lot of timers back and forth if the timers rarely expire.
	// We'll postpone looking through all the adjusted timers until
	// one would actually expire.
	var first = th.timerModifiedEarliest.Load()
	if first == 0 || first > now {
		if verifyTimers {
			th.verifyTimerHeap()
		}
		return
	}

	// We are going to clear all protocol.TimerStatus_ModifiedEarlier timers.
	th.timerModifiedEarliest.Store(0)

	var moved []*Async
	var timers = th.timers
	var timersLen = len(timers)
	for i := 0; i < timersLen; i++ {
		var timer = timers[i].timer
		var status = timer.status.Load()
		switch status {
		case protocol.TimerStatus_Deleted:
			if timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
				var changed = th.deleteTimer(i)
				if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {
					badTimer()
				}
				th.deletedTimers.Add(-1)
				// Go back to the earliest changed heap entry.
				// "- 1" because the loop will add 1.
				i = changed - 1
			}
		case protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
			if timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
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
		case protocol.TimerStatus_Unset, protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Removed, protocol.TimerStatus_Moving:
			badTimer()
		case protocol.TimerStatus_Waiting:
			// OK, nothing to do.
		case protocol.TimerStatus_Modifying:
			// Check again after modification is complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
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
func (th *TimingHeap) addAdjustedTimers(moved []*Async) {
	for _, t := range moved {
		th.AddTimer(t)
		if !t.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
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
func (th *TimingHeap) runTimer(now monotonic.Time) monotonic.Time {
	for {
		var timer = th.timers[0].timer
		var status = timer.status.Load()
		switch status {
		case protocol.TimerStatus_Waiting:
			if timer.when > now {
				// Not ready to run.
				return timer.when
			}

			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Running) {
				continue
			}
			// Note that runOneTimer may temporarily unlock th.timersLock
			th.runOneTimer(timer, now)
			return 0

		case protocol.TimerStatus_Deleted:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
				continue
			}
			th.deleteTimer0()
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {
				badTimer()
			}
			th.deletedTimers.Add(-1)
			if len(th.timers) == 0 {
				return -1
			}

		case protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
				continue
			}
			th.deleteTimer0()
			th.AddTimer(timer)
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
				badTimer()
			}

		case protocol.TimerStatus_Modifying:
			// Wait for modification to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case protocol.TimerStatus_Unset, protocol.TimerStatus_Removed:
			// Should not see a new or inactive timer on the heap.
			badTimer()
		case protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Moving:
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
func (th *TimingHeap) runOneTimer(t *Async, now monotonic.Time) {
	if race.DetectorEnabled {
		race.AcquireCTX(th.timerRaceCtx, unsafe.Pointer(t))
	}

	if t.period > 0 {
		// Leave in heap but adjust next time to fire.
		var delta = t.when.Since(now)
		t.when.Add(t.period * (1 + -delta/t.period))
		if t.when < 0 { // check for overflow.
			t.when = maxWhen
		}
		th.siftDownTimer(0)
		if !t.status.CompareAndSwap(protocol.TimerStatus_Running, protocol.TimerStatus_Waiting) {
			badTimer()
		}
		th.updateTimer0When()
	} else {
		// Remove from heap.
		th.deleteTimer0()
		if !t.status.CompareAndSwap(protocol.TimerStatus_Running, protocol.TimerStatus_Unset) {
			badTimer()
		}
	}

	if race.DetectorEnabled {
		// Temporarily use the current th.timerRaceCtx for thread
		scheduler.SetRaceCtx(th.timerRaceCtx)
	}

	var callback = t.callback
	th.timersLock.Unlock()
	callback.TimerHandler()
	th.timersLock.Lock()

	if race.DetectorEnabled {
		scheduler.ReleaseRaceCtx()
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
	// We are going to clear all protocol.TimerStatus_ModifiedEarlier timers.
	// Do this now in case new ones show up while we are looping.
	th.timerModifiedEarliest.Store(0)

	var cdel = int32(0)
	var to = 0
	var changedHeap = false
	var timers = th.timers
	var timersLen = len(timers)
nextTimer:
	for i := 0; i < timersLen; i++ {
		var timer = timers[i].timer
		for {
			var status = timer.status.Load()
			switch status {
			case protocol.TimerStatus_Waiting:
				if changedHeap {
					timers[to] = timers[i]
					th.siftUpTimer(to)
				}
				to++
				continue nextTimer
			case protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
				if timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
					timers[i].when = timer.when
					timers[to] = timers[i]
					th.siftUpTimer(to)
					to++
					changedHeap = true
					if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
						badTimer()
					}
					continue nextTimer
				}
			case protocol.TimerStatus_Deleted:
				if timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
					timer.timing = nil
					cdel++
					if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {
						badTimer()
					}
					changedHeap = true
					continue nextTimer
				}
			case protocol.TimerStatus_Modifying:
				// Loop until modification complete.
				scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
			case protocol.TimerStatus_Unset, protocol.TimerStatus_Removed:
				// We should not see these status values in a timer heap.
				badTimer()
			case protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Moving:
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

	th.deletedTimers.Add(-cdel)
	th.numTimers.Add(-cdel)

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
		var p = (i - 1) / heapAry
		if timers[i].when < timers[p].when {
			var logMsg = fmt.Sprint("bad timer heap at ", i, ": ", p, ": ", th.timers[p].when, ", ", i, ": ", timers[i].when, "\n")
			protocol.App.Log(log.FatalEvent("libgo/timer", logMsg))
		}
	}
	var numTimers = int(th.numTimers.Load())
	if timersLen != numTimers {
		var logMsg = fmt.Sprint("timer: bad timer heap len ", len(th.timers), "!= numTimers", numTimers)
		protocol.App.Log(log.FatalEvent("libgo/timer", logMsg))
	}
}

// updateTimer0When sets the timer0When field by check first timer in queue.
// The caller must have locked the th.timersLock
func (th *TimingHeap) updateTimer0When() {
	if len(th.timers) == 0 {
		th.timer0When.Store(0)
	} else {
		th.timer0When.Store(th.timers[0].when)
	}
}

// updateTimerModifiedEarliest updates the th.timerModifiedEarliest value.
// The timers will not be locked.
func (th *TimingHeap) updateTimerModifiedEarliest(nextWhen monotonic.Time) {
	for {
		var old = th.timerModifiedEarliest.Load()
		if old != 0 && old < nextWhen {
			return
		}
		if th.timerModifiedEarliest.CompareAndSwap(old, nextWhen) {
			return
		}
	}
}

// sleepUntil returns the time when the next timer should fire.
func (th *TimingHeap) sleepUntil() (until monotonic.Time) {
	until = maxWhen

	var timer0When = th.timer0When.Load()
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}

	timer0When = th.timerModifiedEarliest.Load()
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}
	return
}

// noBarrierWakeTime looks at timers and returns the time when we should wake up.
// This function is invoked when dropping a Timers, and must run without any write barriers.
// Unlike th.sleepUntil(), It returns 0 if there are no timers.
func (th *TimingHeap) noBarrierWakeTime() (until monotonic.Time) {
	until = th.timer0When.Load()
	var nextAdj = th.timerModifiedEarliest.Load()
	if until == 0 || (nextAdj != 0 && nextAdj < until) {
		until = nextAdj
	}
	return
}

// checkTimers runs any timers that are ready.
// returns the time when the next timer should run (always larger than the now) or 0 if there is no next timer,
// and reports whether it ran any timers.
// We pass now in to avoid extra calls of monotonic.Now().
func (th *TimingHeap) checkTimers(now monotonic.Time) (nextWhen monotonic.Time, ran bool) {
	// If it's not yet time for the first timer, or the first adjusted
	// timer, then there is nothing to do.
	var next = th.noBarrierWakeTime()
	if next == 0 {
		// No timers to run or adjust.
		return 0, false
	}

	if now < next {
		// Next timer is not ready to run, but keep going
		// if we would clear deleted timers.
		// This corresponds to the condition below where
		// we decide whether to call clearDeletedTimers.
		if th.deletedTimers.Load() <= th.numTimers.Load()/4 {
			return next, false
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
					nextWhen = tw
				}
				break
			}
			ran = true
		}
	}

	// If there are a lot of deleted timers (>25%), clear them out.
	if int(th.deletedTimers.Load()) > len(th.timers)/4 {
		th.clearDeletedTimers()
	}

	th.timersLock.Unlock()
	return
}

// Check for deadlock situation
func (th *TimingHeap) checkDead() {
	// Maybe jump time forward for playground.
	// if faketime != 0 {
	// 	var when = th.sleepUntil()

	// 	faketime = when

	// 	var mp = mget()
	// 	if mp == nil {
	// 		// There should always be a free M since
	// 		// nothing is running.
	// 		panic("timers - checkDead: no m for timer")
	// 	}
	// 	return
	// }

	// There are no goroutines running, so we can look at the P's.
	if len(th.timers) > 0 {
		return
	}
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
		var p = (i - 1) / heapAry // parent
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
		var c = i*heapAry + 1      // left child
		var c3 = c + (heapAry / 2) // mid child
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
