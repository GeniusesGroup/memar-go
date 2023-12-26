/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"

	"memar/cpu"
	"memar/log"
	"memar/protocol"
	"memar/race"
	"memar/scheduler"
	"memar/time/monotonic"
	errs "memar/timer/errors"
)

// TODO::: remove any direct access to tg.timingHeap fields

// Timing ...
//
// https://github.com/search?l=go&q=timer&type=Repositories
// https://github.com/RussellLuo/timingwheel/blob/master/delayqueue/delayqueue.go
type Timing struct {
	coreID uint64 // CPU core number this timing run on it
	thread *scheduler.Thread

	// The when field of the first entry on the timer heap.
	// This is 0 if the timer heap is empty.
	timer0When monotonic.Atomic

	// The earliest known when field of a timer with
	// timerModifiedEarlier status. Because the timer may have been
	// modified again, there need not be any timer with this value.
	// This is 0 if there are no timerModifiedEarlier timers.
	timerModifiedEarliest monotonic.Atomic

	// Number of timers in this timing.
	timersCount atomic.Int32
	// Number of deleted timers in this timing.
	deletedTimersCount atomic.Int32

	// Race context used while executing timer functions.
	timerRaceCtx uintptr

	// The caller MUST have locked the timingHeapSync when use timingHeap methods.
	timingHeap
	timingHeapSync sync.Mutex
}

// Init initialize timing mechanism for the core that call the Init().
//
//memar:impl memar/protocol.SoftwareLifeCycle
func (tg *Timing) Init() (err protocol.Error) {
	err = tg.timingHeap.Init()
	if err != nil {
		return
	}

	tg.coreID = cpu.ActiveCoreID()
	// tg.thread = scheduler.NewThread()
	// tg.timerRaceCtx = racegostart(abi.FuncPCABIInternal(tg.runTimer) + sys.PCQuantum)

	// TODO::: change to memar scheduler
	go tg.Start()
	return
}

// Reinit releases all of the resources associated with timers in specific CPU core and
// move them to other core that call deinit
//
//memar:impl memar/protocol.SoftwareLifeCycle
func (tg *Timing) Reinit() (err protocol.Error) {
	var callerCoreID = cpu.ActiveCoreID()
	var newCore = &poolByCores[callerCoreID]
	tg.moveTimersTo(newCore)

	tg.coreID = callerCoreID
	tg.timer0When.Store(0)
	tg.timerModifiedEarliest.Store(0)
	tg.timersCount.Store(0)
	tg.deletedTimersCount.Store(0)
	tg.timerRaceCtx = 0

	err = tg.timingHeap.Reinit()
	return
}

// Deinit releases all of the resources associated with timers in specific CPU core
//
//memar:impl memar/protocol.SoftwareLifeCycle
func (tg *Timing) Deinit() (err protocol.Error) {
	err = tg.timingHeap.Deinit()
	return
}

func (tg *Timing) Start() {
	// TODO::: Stop mechanism, new timer added mechanism
	for {
		var now = monotonic.Now()
		var nextWhen, _ = tg.checkTimers(now)
		var until = nextWhen.Until(now)
		tg.thread.Sleep(until)
	}
}

// MoveToMe releases all of the resources associated with timers in specific CPU core and
// move them to other core that call this method
func (tg *Timing) MoveToMe() {
	var callerCoreID = cpu.ActiveCoreID()
	var newCore = &poolByCores[callerCoreID]
	tg.moveTimersTo(newCore)
}

// AddTimer adds t to the timers queue.
func (tg *Timing) AddTimer(t *Async) {
	tg.timingHeapSync.Lock()

	tg.cleanTimers()

	var timerWhen = t.when
	t.timing = tg
	var i = tg.timingHeap.Len()
	tg.timingHeap.Append(timerBucketHeap{t, timerWhen})

	tg.timingHeap.SiftUpTimer(i)
	if t == tg.timers[0].timer {
		tg.timer0When.Store(timerWhen)
	}
	tg.timersCount.Add(1)

	tg.timingHeapSync.Unlock()
}

// deleteTimer removes timer i from the timers heap.
// It returns the smallest changed index in tg.timingHeap
// The caller must have locked the tg.timingHeapSync
func (tg *Timing) deleteTimer(i int) (smallestChanged int) {
	smallestChanged = tg.timingHeap.DeleteTimer(i)

	if i == 0 {
		tg.updateTimer0When()
	}

	var timerRemaining = tg.timersCount.Add(-1)
	if timerRemaining == 0 {
		// If there are no timers, then clearly none are modified.
		tg.timerModifiedEarliest.Store(0)
	}
	return
}

// deleteTimer0 removes timer 0 from the timers heap.
// It reports whether it saw no problems due to races.
// The caller must have locked the tg.timingHeapSync
func (tg *Timing) deleteTimer0() {
	tg.timingHeap.DeleteTimer0()
	tg.updateTimer0When()

	var timerRemaining = tg.timersCount.Add(-1)
	if timerRemaining == 0 {
		// If there are no timers, then clearly none are modified.
		tg.timerModifiedEarliest.Store(0)
	}
}

// cleanTimers cleans up the head of the timer queue. This speeds up
// programs that create and delete timers; leaving them in the heap
// slows down AddTimer. Reports whether no timer problems were found.
// The caller must have locked the tg.timingHeapSync
func (tg *Timing) cleanTimers() {
	if tg.timingHeap.Len() == 0 {
		return
	}

	for {
		// This loop can theoretically run for a while, and because it is holding timers.timingHeapSync.Lock()
		// it cannot be preempted. If someone is trying to preempt us, just return.
		// We can clean the timers later.
		// if gp.preemptStop {
		// 	return
		// }

		var timer = tg.timers[0].timer
		var status = timer.status.Load()
		switch status {
		case protocol.TimerStatus_Deleted:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
				continue
			}
			tg.deleteTimer0()
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {
				log.Fatal(&errs.ErrTimerRacyAccess, "cleanTimers: Racy timer access: Removing to Removed")
			}
			tg.deletedTimersCount.Add(-1)
		case protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
				continue
			}
			// Now we can change the when field of timerBucketHeap.
			tg.timers[0].when = timer.when
			// Move timer to the right position.
			tg.deleteTimer0()
			tg.AddTimer(timer)
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
				log.Fatal(&errs.ErrTimerRacyAccess, "cleanTimers: Racy timer access: Moving to Waiting")
			}
		default:
			// Head of timers does not need adjustment.
			return
		}
	}
}

func (tg *Timing) moveTimersTo(to *Timing) {
	if tg.timingHeap.Len() > 0 {
		tg.timingHeapSync.Lock()

		to.timingHeapSync.Lock()
		to.moveTimers(tg.timers)
		to.timingHeapSync.Unlock()

		tg.timingHeapSync.Unlock()
	}
}

// moveTimers moves a slice of timers to the timers heap.
// The slice has been taken from a different Timers.
// This is currently called when the world is stopped, but the caller
// is expected to have locked the tg.timingHeapSync
func (tg *Timing) moveTimers(timers []timerBucketHeap) {
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
				tg.AddTimer(timer)
				if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
					log.Fatal(&errs.ErrTimerRacyAccess, "moveTimers: Racy timer access: Moving to Waiting")
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
				log.Fatal(&errs.ErrTimerRacyAccess, "moveTimers: Bad timer status: Unset||Removed")
			case protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Moving:
				// Some other P thinks it owns this timer, which should not happen.
				log.Fatal(&errs.ErrTimerRacyAccess, "moveTimers: Bad timer status: Running||Removing||Moving")
			default:
				log.Fatal(&errs.ErrTimerRacyAccess, "moveTimers: Unknown timer status")
			}
		}
	}
}

// adjustTimers looks through the timers for any timers that have been modified to run earlier,
// and puts them in the correct place in the heap. While looking for those timers,
// it also moves timers that have been modified to run later, and removes deleted timers.
// The caller must have locked the tg.timingHeapSync
func (tg *Timing) adjustTimers(now monotonic.Time) {
	// If we haven't yet reached the time of the first protocol.TimerStatus_ModifiedEarlier
	// timer, don't do anything. This speeds up programs that adjust
	// a lot of timers back and forth if the timers rarely expire.
	// We'll postpone looking through all the adjusted timers until
	// one would actually expire.
	var first = tg.timerModifiedEarliest.Load()
	if first == 0 || first > now {
		if verifyTimers {
			tg.verifyTimerHeap()
		}
		return
	}

	// We are going to clear all protocol.TimerStatus_ModifiedEarlier timers.
	tg.timerModifiedEarliest.Store(0)

	var moved []*Async
	var timers = tg.timers
	var timersLen = len(timers)
	for i := 0; i < timersLen; i++ {
		var timer = timers[i].timer
		var status = timer.status.Load()
		switch status {
		case protocol.TimerStatus_Deleted:
			if timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
				var changed = tg.deleteTimer(i)
				if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {
					log.Fatal(&errs.ErrTimerRacyAccess, "adjustTimers: Racy timer access: Removing to Removed")
				}
				tg.deletedTimersCount.Add(-1)
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
				var changed = tg.deleteTimer(i)
				moved = append(moved, timer)
				// Go back to the earliest changed heap entry.
				// "- 1" because the loop will add 1.
				i = changed - 1
			}
		case protocol.TimerStatus_Unset, protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Removed, protocol.TimerStatus_Moving:
			log.Fatal(&errs.ErrTimerRacyAccess, "adjustTimers: Bad timer status: Unset||Running||Removing||Removed||Moving")
		case protocol.TimerStatus_Waiting:
			// OK, nothing to do.
		case protocol.TimerStatus_Modifying:
			// Check again after modification is complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
			i--
		default:
			log.Fatal(&errs.ErrTimerRacyAccess, "adjustTimers: Unknown timer status")
		}
	}

	if len(moved) > 0 {
		tg.addAdjustedTimers(moved)
	}

	if verifyTimers {
		tg.verifyTimerHeap()
	}
}

// addAdjustedTimers adds any timers we adjusted in tg.adjustTimers
// back to the timer heap.
func (tg *Timing) addAdjustedTimers(moved []*Async) {
	for _, t := range moved {
		tg.AddTimer(t)
		if !t.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
			log.Fatal(&errs.ErrTimerRacyAccess, "addAdjustedTimers: Racy timer access: Moving to Waiting")
		}
	}
}

// runTimer examines the first timer in timers. If it is ready based on now,
// it runs the timer and removes or updates it.
// Returns 0 if it ran a timer, -1 if there are no more timers, or the time
// when the first timer should run.
// The caller must have locked the tg.timingHeapSync
// If a timer is run, this will temporarily unlock the timers.
func (tg *Timing) runTimer(now monotonic.Time) monotonic.Time {
	for {
		var timer = tg.timers[0].timer
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
			// Note that runOneTimer may temporarily unlock tg.timers
			tg.runOneTimer(timer, now)
			return 0

		case protocol.TimerStatus_Deleted:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
				continue
			}
			tg.deleteTimer0()
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {

				log.Fatal(&errs.ErrTimerRacyAccess, "runTimer: Racy timer access: Removing to Removed")
			}
			tg.deletedTimersCount.Add(-1)
			if tg.timingHeap.Len() == 0 {
				return -1
			}

		case protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
			if !timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
				continue
			}
			tg.deleteTimer0()
			tg.AddTimer(timer)
			if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
				log.Fatal(&errs.ErrTimerRacyAccess, "runTimer: Racy timer access: Moving to Waiting")
			}

		case protocol.TimerStatus_Modifying:
			// Wait for modification to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case protocol.TimerStatus_Unset, protocol.TimerStatus_Removed:
			// Should not see a new or inactive timer on the heap.
			log.Fatal(&errs.ErrTimerRacyAccess, "runTimer: Bad timer status: Unset||Removed")
		case protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Moving:
			// These should only be set when timers are locked, and we didn't do it.
			log.Fatal(&errs.ErrTimerRacyAccess, "runTimer: Bad timer status: Running||Removing||Moving")
		default:
			log.Fatal(&errs.ErrTimerRacyAccess, "runTimer: Unknown timer status")
		}
	}
}

// runOneTimer runs a single timer.
// The caller must have locked the tg.timingHeapSync
// This will temporarily unlock the timers while running the timer function.
func (tg *Timing) runOneTimer(t *Async, now monotonic.Time) {
	if race.DetectorEnabled {
		race.AcquireCTX(tg.timerRaceCtx, unsafe.Pointer(t))
	}

	if t.period > 0 {
		// Leave in heap but adjust next time to fire.
		var delta = t.when.Since(now)
		t.when.Add(t.period * (1 + -delta/t.period))
		if t.when < 0 { // check for overflow.
			t.when = maxWhen
		}
		tg.timingHeap.SiftDownTimer(0)
		if !t.status.CompareAndSwap(protocol.TimerStatus_Running, protocol.TimerStatus_Waiting) {
			log.Fatal(&errs.ErrTimerRacyAccess, "runOneTimer: Racy timer access: Running to Waiting")
		}
		tg.updateTimer0When()
	} else {
		// Remove from heap.
		tg.deleteTimer0()
		if !t.status.CompareAndSwap(protocol.TimerStatus_Running, protocol.TimerStatus_Unset) {
			log.Fatal(&errs.ErrTimerRacyAccess, "runOneTimer: Racy timer access: Running to Unset")
		}
	}

	if race.DetectorEnabled {
		// Temporarily use the current tg.timerRaceCtx for thread
		scheduler.SetRaceCtx(tg.timerRaceCtx)
	}

	var callback = t.callback
	tg.timingHeapSync.Unlock()
	callback.TimerHandler()
	tg.timingHeapSync.Lock()

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
// The caller must have locked the tg.timingHeapSync
func (tg *Timing) clearDeletedTimers() {
	// We are going to clear all protocol.TimerStatus_ModifiedEarlier timers.
	// Do this now in case new ones show up while we are looping.
	tg.timerModifiedEarliest.Store(0)

	var cdel = int32(0)
	var to = 0
	var changedHeap = false
	var timers = tg.timingHeap.timers
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
					tg.timingHeap.SiftUpTimer(to)
				}
				to++
				continue nextTimer
			case protocol.TimerStatus_ModifiedEarlier, protocol.TimerStatus_ModifiedLater:
				if timer.status.CompareAndSwap(status, protocol.TimerStatus_Moving) {
					timers[i].when = timer.when
					timers[to] = timers[i]
					tg.timingHeap.SiftUpTimer(to)
					to++
					changedHeap = true
					if !timer.status.CompareAndSwap(protocol.TimerStatus_Moving, protocol.TimerStatus_Waiting) {
						log.Fatal(&errs.ErrTimerRacyAccess, "clearDeletedTimers: Racy timer access: Moving to Waiting")
					}
					continue nextTimer
				}
			case protocol.TimerStatus_Deleted:
				if timer.status.CompareAndSwap(status, protocol.TimerStatus_Removing) {
					timer.timing = nil
					cdel++
					if !timer.status.CompareAndSwap(protocol.TimerStatus_Removing, protocol.TimerStatus_Removed) {
						log.Fatal(&errs.ErrTimerRacyAccess, "clearDeletedTimers: Racy timer access: Removing to Removed")
					}
					changedHeap = true
					continue nextTimer
				}
			case protocol.TimerStatus_Modifying:
				// Loop until modification complete.
				scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
			case protocol.TimerStatus_Unset, protocol.TimerStatus_Removed:
				// We should not see these status values in a timer heap.
				log.Fatal(&errs.ErrTimerRacyAccess, "clearDeletedTimers: Bad timer status: Unset||Removed")
			case protocol.TimerStatus_Running, protocol.TimerStatus_Removing, protocol.TimerStatus_Moving:
				// Some other P thinks it owns this timer, which should not happen.
				log.Fatal(&errs.ErrTimerRacyAccess, "clearDeletedTimers: Bad timer status: Running||Removing||Moving")
			default:
				log.Fatal(&errs.ErrTimerRacyAccess, "clearDeletedTimers: Unknown timer status")
			}
		}
	}

	// Deinit remaining slots in timers slice,
	// so that the timer values can be garbage collected.
	for i := to; i < len(timers); i++ {
		timers[i].Deinit()
	}

	tg.deletedTimersCount.Add(-cdel)
	tg.timersCount.Add(-cdel)

	timers = timers[:to]
	tg.timingHeap.timers = timers
	tg.updateTimer0When()

	if verifyTimers {
		tg.verifyTimerHeap()
	}
}

// verifyTimerHeap verifies that the timer heap is in a valid state.
// This is only for debugging, and is only called if verifyTimers is true.
// The caller must have locked the tg.timingHeapSync
func (tg *Timing) verifyTimerHeap() {
	var timers = tg.timingHeap.timers
	var timersLen = len(timers)
	// First timer has no parent, so i must be start from 1.
	for i := 1; i < timersLen; i++ {
		var p = (i - 1) / heapAry
		if timers[i].when < timers[p].when {
			var logMsg = fmt.Sprint("bad timer heap at ", i, ": ", p, ": ", tg.timingHeap.timers[p].when, ", ", i, ": ", timers[i].when, "\n")
			log.Fatal(&errs.ErrTimerRacyAccess, logMsg)
		}
	}
	var timersCount = int(tg.timersCount.Load())
	if timersLen != timersCount {
		var logMsg = fmt.Sprint("timer: bad timer heap len ", tg.timingHeap.Len(), "!= timersCount", timersCount)
		log.Fatal(&errs.ErrTimerRacyAccess, logMsg)
	}
}

// updateTimer0When sets the timer0When field by check first timer in queue.
// The caller must have locked the tg.timingHeapSync
func (tg *Timing) updateTimer0When() {
	if tg.timingHeap.Len() == 0 {
		tg.timer0When.Store(0)
	} else {
		tg.timer0When.Store(tg.timers[0].when)
	}
}

// updateTimerModifiedEarliest updates the tg.timerModifiedEarliest value.
// The tg.timingHeapSync will not be locked.
func (tg *Timing) updateTimerModifiedEarliest(nextWhen monotonic.Time) {
	for {
		var old = tg.timerModifiedEarliest.Load()
		if old != 0 && old < nextWhen {
			return
		}
		if tg.timerModifiedEarliest.CompareAndSwap(old, nextWhen) {
			return
		}
	}
}

// sleepUntil returns the time when the next timer should fire.
func (tg *Timing) sleepUntil() (until monotonic.Time) {
	until = maxWhen

	var timer0When = tg.timer0When.Load()
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}

	timer0When = tg.timerModifiedEarliest.Load()
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}
	return
}

// noBarrierWakeTime looks at timers and returns the time when we should wake up.
// This function is invoked when dropping a Timers, and must run without any write barriers.
// Unlike tg.sleepUntil(), It returns 0 if there are no timers.
func (tg *Timing) noBarrierWakeTime() (until monotonic.Time) {
	until = tg.timer0When.Load()
	var nextAdj = tg.timerModifiedEarliest.Load()
	if until == 0 || (nextAdj != 0 && nextAdj < until) {
		until = nextAdj
	}
	return
}

// This corresponds to the condition below where we decide whether to call clearDeletedTimers.
// If there are a lot of deleted timers (>25%), clear them out.
func (tg *Timing) isCleanNeed() (needClean bool) {
	if tg.deletedTimersCount.Load() <= tg.timersCount.Load()/4 {
		return false
	}
	return true
}

// checkTimers runs any timers that are ready.
// returns the time when the next timer should run (always larger than the now) or 0 if there is no next timer,
// and reports whether it ran any timers.
// We pass now in to avoid extra calls of monotonic.Now().
func (tg *Timing) checkTimers(now monotonic.Time) (nextWhen monotonic.Time, ran bool) {
	// If it's not yet time for the first timer, or the first adjusted
	// timer, then there is nothing to do.
	var next = tg.noBarrierWakeTime()
	if next == 0 {
		// No timers to run or adjust.
		return 0, false
	}

	if now < next {
		// Next timer is not ready to run, but keep going
		// if we would clear deleted timers.
		if !tg.isCleanNeed() {
			return next, false
		}
	}

	tg.timingHeapSync.Lock()

	if tg.timingHeap.Len() > 0 {
		tg.adjustTimers(now)
		for tg.timingHeap.Len() > 0 {
			// Note that tg.runTimer may temporarily unlock tg.timingHeap.
			var tw = tg.runTimer(now)
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
	if int(tg.deletedTimersCount.Load()) > tg.timingHeap.Len()/4 {
		tg.clearDeletedTimers()
	}

	tg.timingHeapSync.Unlock()
	return
}

// Check for deadlock situation
func (tg *Timing) checkDead() (err protocol.Error) {
	// Maybe jump time forward for playground.
	// if faketime != 0 {
	// 	var when = tg.sleepUntil()

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
	if tg.timingHeap.Len() > 0 {
		return
	}
	return
}
