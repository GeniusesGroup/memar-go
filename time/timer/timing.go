/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"fmt"
	"sync"
	"sync/atomic"

	"memar/process/log"
	"memar/hardware/cpu"
	error_p "memar/process/error/protocol"
	race "memar/process/race"
	"memar/computer/runtime/scheduler"
	"memar/time/monotonic"
	timer_errs "memar/time/timer/errors"
)

// TODO::: remove any direct access to self.timingHeap fields
// TODO::: Can remove timingHeapSync by access the each Timing on same CPU core that registered??
// TODO::: Really need race detector with above changes??

// Timing ...
//
// https://github.com/search?l=go&q=timer&type=Repositories
// https://github.com/RussellLuo/timingwheel/blob/master/delayqueue/delayqueue.go
type Timing struct {
	coreID cpu.CoreID // CPU core number this timing run on it
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

	// The caller MUST have locked the timingHeapSync when use timingHeap methods.
	timingHeap
	timingHeapSync sync.Mutex

	race race.Dynamic
}

// Init initialize timing mechanism for the core that call the Init().
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Timing) Init() (err error_p.Error) {
	err = self.timingHeap.Init()
	if err != nil {
		return
	}

	self.coreID.Active()
	// TODO::: Make high priority thread
	// self.thread = scheduler.NewThread()
	
	// TODO::: Why developer need declare race detection manually??
	self.race.Init(self.runTimer)

	// TODO::: change to memar scheduler
	go self.Start()
	return
}

// Reinit releases all of the resources associated with timers in specific CPU core and
// move them to other core that call deinit
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Timing) Reinit() (err error_p.Error) {
	// TODO::: FIX below logic
	self.coreID.Active()
	var newCore = &poolByCores[self.coreID]
	self.moveTimersTo(newCore)

	self.timer0When.Store(0)
	self.timerModifiedEarliest.Store(0)
	self.timersCount.Store(0)
	self.deletedTimersCount.Store(0)

	// TODO::: Why developer need declare race detection manually??
	self.race.Reinit(self.runTimer)

	err = self.timingHeap.Reinit()
	return
}

// Deinit releases all of the resources associated with timers in specific CPU core
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Timing) Deinit() (err error_p.Error) {
	err = self.timingHeap.Deinit()
	// TODO::: Why developer need declare race detection manually??
	err = self.race.Deinit()
	return
}

func (self *Timing) Start() {
	// TODO::: Stop mechanism, new timer added mechanism
	for {
		var now = monotonic.Now()
		var nextWhen, _ = self.checkTimers(now)
		var until = nextWhen.Until(now)
		self.thread.Sleep(until)
	}
}

// MoveToMe releases all of the resources associated with timers in specific CPU core and
// move them to other core that call this method
func (self *Timing) MoveToMe() {
	var callerCoreID cpu.CoreID
	callerCoreID.Active()
	var newCore = &poolByCores[callerCoreID]
	self.moveTimersTo(newCore)
}

// AddTimer adds t to the timers queue.
func (self *Timing) AddTimer(t *Async) {
	self.timingHeapSync.Lock()

	self.cleanTimers()

	var timerWhen = t.when
	t.timing = self
	var i = self.timingHeap.OccupiedLength()
	self.timingHeap.Append(timerBucketHeap{t, timerWhen})

	self.timingHeap.SiftUpTimer(i)
	if t == self.timers[0].timer {
		self.timer0When.Store(timerWhen)
	}
	self.timersCount.Add(1)

	self.timingHeapSync.Unlock()
}

// deleteTimer removes timer i from the timers heap.
// It returns the smallest changed index in self.timingHeap
// The caller must have locked the self.timingHeapSync
func (self *Timing) deleteTimer(i int) (smallestChanged int) {
	smallestChanged = self.timingHeap.DeleteTimer(i)

	if i == 0 {
		self.updateTimer0When()
	}

	var timerRemaining = self.timersCount.Add(-1)
	if timerRemaining == 0 {
		// If there are no timers, then clearly none are modified.
		self.timerModifiedEarliest.Store(0)
	}
	return
}

// deleteTimer0 removes timer 0 from the timers heap.
// It reports whether it saw no problems due to races.
// The caller must have locked the self.timingHeapSync
func (self *Timing) deleteTimer0() {
	self.timingHeap.DeleteTimer0()
	self.updateTimer0When()

	var timerRemaining = self.timersCount.Add(-1)
	if timerRemaining == 0 {
		// If there are no timers, then clearly none are modified.
		self.timerModifiedEarliest.Store(0)
	}
}

// cleanTimers cleans up the head of the timer queue. This speeds up
// programs that create and delete timers; leaving them in the heap
// slows down AddTimer. Reports whether no timer problems were found.
// The caller must have locked the self.timingHeapSync
func (self *Timing) cleanTimers() {
	if self.timingHeap.OccupiedLength() == 0 {
		return
	}

	for {
		// This loop can theoretically run for a while, and because it is holding timers.timingHeapSync.Lock()
		// it cannot be preempted. If someone is trying to preempt us, just return.
		// We can clean the timers later.
		// if gp.preemptStop {
		// 	return
		// }

		var timer = self.timers[0].timer
		var status = timer.status.Load()
		switch status {
		case Status_Deleted:
			if !timer.status.CompareAndSwap(status, Status_Removing) {
				continue
			}
			self.deleteTimer0()
			if !timer.status.CompareAndSwap(Status_Removing, Status_Removed) {
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "cleanTimers: Racy timer access: Removing to Removed")
			}
			self.deletedTimersCount.Add(-1)
		case Status_ModifiedEarlier, Status_ModifiedLater:
			if !timer.status.CompareAndSwap(status, Status_Moving) {
				continue
			}
			// Now we can change the when field of timerBucketHeap.
			self.timers[0].when = timer.when
			// Move timer to the right position.
			self.deleteTimer0()
			self.AddTimer(timer)
			if !timer.status.CompareAndSwap(Status_Moving, Status_Waiting) {
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "cleanTimers: Racy timer access: Moving to Waiting")
			}
		default:
			// Head of timers does not need adjustment.
			return
		}
	}
}

func (self *Timing) moveTimersTo(to *Timing) {
	if self.timingHeap.OccupiedLength() > 0 {
		self.timingHeapSync.Lock()

		to.timingHeapSync.Lock()
		to.moveTimers(self.timers)
		to.timingHeapSync.Unlock()

		self.timingHeapSync.Unlock()
	}
}

// moveTimers moves a slice of timers to the timers heap.
// The slice has been taken from a different Timers.
// This is currently called when the world is stopped, but the caller
// is expected to have locked the self.timingHeapSync
func (self *Timing) moveTimers(timers []timerBucketHeap) {
	for _, timerBucketHeap := range timers {
		var timer = timerBucketHeap.timer
	loop:
		for {
			var status = timer.status.Load()
			switch status {
			case Status_Waiting, Status_ModifiedEarlier, Status_ModifiedLater:
				if !timer.status.CompareAndSwap(status, Status_Moving) {
					continue
				}
				timer.timing = nil
				self.AddTimer(timer)
				if !timer.status.CompareAndSwap(Status_Moving, Status_Waiting) {
					log.Fatal(&timer_errs.ErrTimerRacyAccess, "moveTimers: Racy timer access: Moving to Waiting")
				}
				break loop
			case Status_Deleted:
				if !timer.status.CompareAndSwap(status, Status_Removed) {
					continue
				}
				timer.timing = nil
				// We no longer need this timer in the heap.
				break loop
			case Status_Modifying:
				// Loop until the modification is complete.
				scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
			case Status_Unset, Status_Removed:
				// We should not see these status values in a timers heap.
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "moveTimers: Bad timer status: Unset||Removed")
			case Status_Running, Status_Removing, Status_Moving:
				// Some other P thinks it owns this timer, which should not happen.
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "moveTimers: Bad timer status: Running||Removing||Moving")
			default:
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "moveTimers: Unknown timer status")
			}
		}
	}
}

// adjustTimers looks through the timers for any timers that have been modified to run earlier,
// and puts them in the correct place in the heap. While looking for those timers,
// it also moves timers that have been modified to run later, and removes deleted timers.
// The caller must have locked the self.timingHeapSync
func (self *Timing) adjustTimers(now monotonic.Time) {
	// If we haven't yet reached the time of the first Status_ModifiedEarlier
	// timer, don't do anything. This speeds up programs that adjust
	// a lot of timers back and forth if the timers rarely expire.
	// We'll postpone looking through all the adjusted timers until
	// one would actually expire.
	var first = self.timerModifiedEarliest.Load()
	if first == 0 || first > now {
		if verifyTimers {
			self.verifyTimerHeap()
		}
		return
	}

	// We are going to clear all Status_ModifiedEarlier timers.
	self.timerModifiedEarliest.Store(0)

	var moved []*Async
	var timers = self.timers
	var timersLen = len(timers)
	for i := 0; i < timersLen; i++ {
		var timer = timers[i].timer
		var status = timer.status.Load()
		switch status {
		case Status_Deleted:
			if timer.status.CompareAndSwap(status, Status_Removing) {
				var changed = self.deleteTimer(i)
				if !timer.status.CompareAndSwap(Status_Removing, Status_Removed) {
					log.Fatal(&timer_errs.ErrTimerRacyAccess, "adjustTimers: Racy timer access: Removing to Removed")
				}
				self.deletedTimersCount.Add(-1)
				// Go back to the earliest changed heap entry.
				// "- 1" because the loop will add 1.
				i = changed - 1
			}
		case Status_ModifiedEarlier, Status_ModifiedLater:
			if timer.status.CompareAndSwap(status, Status_Moving) {
				// Take t off the heap, and hold onto it.
				// We don't add it back yet because the
				// heap manipulation could cause our
				// loop to skip some other timer.
				var changed = self.deleteTimer(i)
				moved = append(moved, timer)
				// Go back to the earliest changed heap entry.
				// "- 1" because the loop will add 1.
				i = changed - 1
			}
		case Status_Unset, Status_Running, Status_Removing, Status_Removed, Status_Moving:
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "adjustTimers: Bad timer status: Unset||Running||Removing||Removed||Moving")
		case Status_Waiting:
			// OK, nothing to do.
		case Status_Modifying:
			// Check again after modification is complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
			i--
		default:
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "adjustTimers: Unknown timer status")
		}
	}

	if len(moved) > 0 {
		self.addAdjustedTimers(moved)
	}

	if verifyTimers {
		self.verifyTimerHeap()
	}
}

// addAdjustedTimers adds any timers we adjusted in self.adjustTimers
// back to the timer heap.
func (self *Timing) addAdjustedTimers(moved []*Async) {
	for _, t := range moved {
		self.AddTimer(t)
		if !t.status.CompareAndSwap(Status_Moving, Status_Waiting) {
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "addAdjustedTimers: Racy timer access: Moving to Waiting")
		}
	}
}

// runTimer examines the first timer in timers. If it is ready based on now,
// it runs the timer and removes or updates it.
// Returns 0 if it ran a timer, -1 if there are no more timers, or the time
// when the first timer should run.
// The caller must have locked the self.timingHeapSync
// If a timer is run, this will temporarily unlock the timers.
func (self *Timing) runTimer(now monotonic.Time) monotonic.Time {
	for {
		var timer = self.timers[0].timer
		var status = timer.status.Load()
		switch status {
		case Status_Waiting:
			if timer.when > now {
				// Not ready to run.
				return timer.when
			}

			if !timer.status.CompareAndSwap(status, Status_Running) {
				continue
			}
			// Note that runOneTimer may temporarily unlock self.timers
			self.runOneTimer(timer, now)
			return 0

		case Status_Deleted:
			if !timer.status.CompareAndSwap(status, Status_Removing) {
				continue
			}
			self.deleteTimer0()
			if !timer.status.CompareAndSwap(Status_Removing, Status_Removed) {

				log.Fatal(&timer_errs.ErrTimerRacyAccess, "runTimer: Racy timer access: Removing to Removed")
			}
			self.deletedTimersCount.Add(-1)
			if self.timingHeap.OccupiedLength() == 0 {
				return -1
			}

		case Status_ModifiedEarlier, Status_ModifiedLater:
			if !timer.status.CompareAndSwap(status, Status_Moving) {
				continue
			}
			self.deleteTimer0()
			self.AddTimer(timer)
			if !timer.status.CompareAndSwap(Status_Moving, Status_Waiting) {
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "runTimer: Racy timer access: Moving to Waiting")
			}

		case Status_Modifying:
			// Wait for modification to complete.
			scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
		case Status_Unset, Status_Removed:
			// Should not see a new or inactive timer on the heap.
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "runTimer: Bad timer status: Unset||Removed")
		case Status_Running, Status_Removing, Status_Moving:
			// These should only be set when timers are locked, and we didn't do it.
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "runTimer: Bad timer status: Running||Removing||Moving")
		default:
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "runTimer: Unknown timer status")
		}
	}
}

// runOneTimer runs a single timer.
// The caller must have locked the self.timingHeapSync
// This will temporarily unlock the timers while running the timer function.
func (self *Timing) runOneTimer(t *Async, now monotonic.Time) {
	if race.DetectorEnabled {
		self.race.Acquire(t)
	}

	if t.period > 0 {
		// Leave in heap but adjust next time to fire.
		var delta = t.when.Since(now)
		t.when.Add(t.period * (1 + -delta/t.period))
		if t.when < 0 { // check for overflow.
			t.when = maxWhen
		}
		self.timingHeap.SiftDownTimer(0)
		if !t.status.CompareAndSwap(Status_Running, Status_Waiting) {
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "runOneTimer: Racy timer access: Running to Waiting")
		}
		self.updateTimer0When()
	} else {
		// Remove from heap.
		self.deleteTimer0()
		if !t.status.CompareAndSwap(Status_Running, Status_Unset) {
			log.Fatal(&timer_errs.ErrTimerRacyAccess, "runOneTimer: Racy timer access: Running to Unset")
		}
	}

	if race.DetectorEnabled {
		// Temporarily use the current self.race for thread
		scheduler.SetRaceCtx(self.race)
	}

	var callback = t.callback
	self.timingHeapSync.Unlock()
	callback.TimerHandler()
	self.timingHeapSync.Lock()

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
// The caller must have locked the self.timingHeapSync
func (self *Timing) clearDeletedTimers() {
	// We are going to clear all Status_ModifiedEarlier timers.
	// Do this now in case new ones show up while we are looping.
	self.timerModifiedEarliest.Store(0)

	var cdel = int32(0)
	var to = 0
	var changedHeap = false
	var timers = self.timingHeap.timers
	var timersLen = len(timers)
nextTimer:
	for i := 0; i < timersLen; i++ {
		var timer = timers[i].timer
		for {
			var status = timer.status.Load()
			switch status {
			case Status_Waiting:
				if changedHeap {
					timers[to] = timers[i]
					self.timingHeap.SiftUpTimer(to)
				}
				to++
				continue nextTimer
			case Status_ModifiedEarlier, Status_ModifiedLater:
				if timer.status.CompareAndSwap(status, Status_Moving) {
					timers[i].when = timer.when
					timers[to] = timers[i]
					self.timingHeap.SiftUpTimer(to)
					to++
					changedHeap = true
					if !timer.status.CompareAndSwap(Status_Moving, Status_Waiting) {
						log.Fatal(&timer_errs.ErrTimerRacyAccess, "clearDeletedTimers: Racy timer access: Moving to Waiting")
					}
					continue nextTimer
				}
			case Status_Deleted:
				if timer.status.CompareAndSwap(status, Status_Removing) {
					timer.timing = nil
					cdel++
					if !timer.status.CompareAndSwap(Status_Removing, Status_Removed) {
						log.Fatal(&timer_errs.ErrTimerRacyAccess, "clearDeletedTimers: Racy timer access: Removing to Removed")
					}
					changedHeap = true
					continue nextTimer
				}
			case Status_Modifying:
				// Loop until modification complete.
				scheduler.Yield(scheduler.Thread_WaitReason_Preempted)
			case Status_Unset, Status_Removed:
				// We should not see these status values in a timer heap.
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "clearDeletedTimers: Bad timer status: Unset||Removed")
			case Status_Running, Status_Removing, Status_Moving:
				// Some other P thinks it owns this timer, which should not happen.
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "clearDeletedTimers: Bad timer status: Running||Removing||Moving")
			default:
				log.Fatal(&timer_errs.ErrTimerRacyAccess, "clearDeletedTimers: Unknown timer status")
			}
		}
	}

	// Deinit remaining slots in timers slice,
	// so that the timer values can be garbage collected.
	for i := to; i < len(timers); i++ {
		timers[i].Deinit()
	}

	self.deletedTimersCount.Add(-cdel)
	self.timersCount.Add(-cdel)

	timers = timers[:to]
	self.timingHeap.timers = timers
	self.updateTimer0When()

	if verifyTimers {
		self.verifyTimerHeap()
	}
}

// verifyTimerHeap verifies that the timer heap is in a valid state.
// This is only for debugging, and is only called if verifyTimers is true.
// The caller must have locked the self.timingHeapSync
func (self *Timing) verifyTimerHeap() {
	var timers = self.timingHeap.timers
	var timersLen = len(timers)
	// First timer has no parent, so i must be start from 1.
	for i := 1; i < timersLen; i++ {
		var p = (i - 1) / heapAry
		if timers[i].when < timers[p].when {
			var logMsg = fmt.Sprint("bad timer heap at ", i, ": ", p, ": ", self.timingHeap.timers[p].when, ", ", i, ": ", timers[i].when, "\n")
			log.Fatal(&timer_errs.ErrTimerRacyAccess, logMsg)
		}
	}
	var timersCount = int(self.timersCount.Load())
	if timersLen != timersCount {
		var logMsg = fmt.Sprint("timer: bad timer heap len ", self.timingHeap.OccupiedLength(), "!= timersCount", timersCount)
		log.Fatal(&timer_errs.ErrTimerRacyAccess, logMsg)
	}
}

// updateTimer0When sets the timer0When field by check first timer in queue.
// The caller must have locked the self.timingHeapSync
func (self *Timing) updateTimer0When() {
	if self.timingHeap.OccupiedLength() == 0 {
		self.timer0When.Store(0)
	} else {
		self.timer0When.Store(self.timers[0].when)
	}
}

// updateTimerModifiedEarliest updates the self.timerModifiedEarliest value.
// The self.timingHeapSync will not be locked.
func (self *Timing) updateTimerModifiedEarliest(nextWhen monotonic.Time) {
	for {
		var old = self.timerModifiedEarliest.Load()
		if old != 0 && old < nextWhen {
			return
		}
		if self.timerModifiedEarliest.CompareAndSwap(old, nextWhen) {
			return
		}
	}
}

// sleepUntil returns the time when the next timer should fire.
func (self *Timing) sleepUntil() (until monotonic.Time) {
	until = maxWhen

	var timer0When = self.timer0When.Load()
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}

	timer0When = self.timerModifiedEarliest.Load()
	if timer0When != 0 && timer0When < until {
		until = timer0When
	}
	return
}

// noBarrierWakeTime looks at timers and returns the time when we should wake up.
// This function is invoked when dropping a Timers, and must run without any write barriers.
// Unlike self.sleepUntil(), It returns 0 if there are no timers.
func (self *Timing) noBarrierWakeTime() (until monotonic.Time) {
	until = self.timer0When.Load()
	var nextAdj = self.timerModifiedEarliest.Load()
	if until == 0 || (nextAdj != 0 && nextAdj < until) {
		until = nextAdj
	}
	return
}

// This corresponds to the condition below where we decide whether to call clearDeletedTimers.
// If there are a lot of deleted timers (>25%), clear them out.
func (self *Timing) isCleanNeed() (needClean bool) {
	if self.deletedTimersCount.Load() <= self.timersCount.Load()/4 {
		return false
	}
	return true
}

// checkTimers runs any timers that are ready.
// returns the time when the next timer should run (always larger than the now) or 0 if there is no next timer,
// and reports whether it ran any timers.
// We pass now in to avoid extra calls of monotonic.Now().
func (self *Timing) checkTimers(now monotonic.Time) (nextWhen monotonic.Time, ran bool) {
	// If it's not yet time for the first timer, or the first adjusted
	// timer, then there is nothing to do.
	var next = self.noBarrierWakeTime()
	if next == 0 {
		// No timers to run or adjust.
		return 0, false
	}

	if now < next {
		// Next timer is not ready to run, but keep going
		// if we would clear deleted timers.
		if !self.isCleanNeed() {
			return next, false
		}
	}

	self.timingHeapSync.Lock()

	if self.timingHeap.OccupiedLength() > 0 {
		self.adjustTimers(now)
		for self.timingHeap.OccupiedLength() > 0 {
			// Note that self.runTimer may temporarily unlock self.timingHeap.
			var tw = self.runTimer(now)
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
	if int(self.deletedTimersCount.Load()) > self.timingHeap.OccupiedLength()/4 {
		self.clearDeletedTimers()
	}

	self.timingHeapSync.Unlock()
	return
}

// Check for deadlock situation
func (self *Timing) checkDead() (err error_p.Error) {
	// Maybe jump time forward for playground.
	// if faketime != 0 {
	// 	var when = self.sleepUntil()

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
	if self.timingHeap.OccupiedLength() > 0 {
		return
	}
	return
}
