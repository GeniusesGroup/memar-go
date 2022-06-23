/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"internal/abi"
	"runtime/internal/atomic"
	"runtime/internal/sys"
	"sync"
	"unsafe"
)

//
// Active timers live in heaps attached to P, in the timers field.
// Inactive timers live there too temporarily, until they are removed.
//
// https://github.com/search?l=go&q=timer&type=Repositories
// https://github.com/RussellLuo/timingwheel/blob/master/delayqueue/delayqueue.go
type Timers struct {
	// The when field of the first entry on the timer heap.
	// This is updated using atomic functions.
	// This is 0 if the timer heap is empty.
	timer0When uint64

	// The earliest known nextwhen field of a timer with
	// timerModifiedEarlier status. Because the timer may have been
	// modified again, there need not be any timer with this value.
	// This is updated using atomic functions.
	// This is 0 if there are no timerModifiedEarlier timers.
	timerModifiedEarliest uint64

	// Lock for timers. We normally access the timers while running
	// on this P, but the scheduler can also do it from a different P.
	timersLock sync.Mutex

	// Actions to take at some time. This is used to implement the
	// standard library's time package.
	// Must hold timersLock to access.
	timers []*Timer
	// tree
	head *Timer

	// Number of timers in P's heap.
	// Modified using atomic instructions.
	numTimers uint32

	// Number of timerDeleted timers in P's heap.
	// Modified using atomic instructions.
	deletedTimers uint32

	// Race context used while executing timer functions.
	timerRaceCtx uintptr
}

// doaddtimer adds t to the current P's heap.
// The caller must have locked the timers for pp.
func (ts *Timers) addTimer(t *timer) {
	// Timers rely on the network poller, so make sure the poller
	// has started.
	if netpollInited == 0 {
		netpollGenericInit()
	}

	if t.pp != 0 {
		throw("doaddtimer: P already set in timer")
	}
	t.pp.set(pp)
	i := len(pp.timers)
	pp.timers = append(pp.timers, t)
	siftupTimer(pp.timers, i)
	if t == pp.timers[0] {
		atomic.Store64(&pp.timer0When, uint64(t.when))
	}
	atomic.Xadd(&pp.numTimers, 1)
}

// dodeltimer removes timer i from the current P's heap.
// We are locked on the P when this is called.
// It returns the smallest changed index in pp.timers.
// The caller must have locked the timers for pp.
func (ts *Timers) delete(i int) int {
	if t := pp.timers[i]; t.pp.ptr() != pp {
		throw("dodeltimer: wrong P")
	} else {
		t.pp = 0
	}
	last := len(pp.timers) - 1
	if i != last {
		pp.timers[i] = pp.timers[last]
	}
	pp.timers[last] = nil
	pp.timers = pp.timers[:last]
	smallestChanged := i
	if i != last {
		// Moving to i may have moved the last timer to a new parent,
		// so sift up to preserve the heap guarantee.
		smallestChanged = siftupTimer(pp.timers, i)
		siftdownTimer(pp.timers, i)
	}
	if i == 0 {
		updateTimer0When(pp)
	}
	atomic.Xadd(&pp.numTimers, -1)
	return smallestChanged
}

// dodeltimer0 removes timer 0 from the current P's heap.
// We are locked on the P when this is called.
// It reports whether it saw no problems due to races.
// The caller must have locked the timers for pp.
func dodeltimer0(pp *p) {
	if t := pp.timers[0]; t.pp.ptr() != pp {
		throw("dodeltimer0: wrong P")
	} else {
		t.pp = 0
	}
	last := len(pp.timers) - 1
	if last > 0 {
		pp.timers[0] = pp.timers[last]
	}
	pp.timers[last] = nil
	pp.timers = pp.timers[:last]
	if last > 0 {
		siftdownTimer(pp.timers, 0)
	}
	updateTimer0When(pp)
	atomic.Xadd(&pp.numTimers, -1)
}

// cleantimers cleans up the head of the timer queue. This speeds up
// programs that create and delete timers; leaving them in the heap
// slows down addtimer. Reports whether no timer problems were found.
// The caller must have locked the timers for pp.
func cleantimers(pp *p) {
	gp := getg()
	for {
		if len(pp.timers) == 0 {
			return
		}

		// This loop can theoretically run for a while, and because
		// it is holding timersLock it cannot be preempted.
		// If someone is trying to preempt us, just return.
		// We can clean the timers later.
		if gp.preemptStop {
			return
		}

		t := pp.timers[0]
		if t.pp.ptr() != pp {
			throw("cleantimers: bad p")
		}
		switch s := atomic.Load(&t.status); s {
		case status_Deleted:
			if !atomic.Cas(&t.status, s, status_Removing) {
				continue
			}
			dodeltimer0(pp)
			if !atomic.Cas(&t.status, status_Removing, status_Removed) {
				badTimer()
			}
			atomic.Xadd(&pp.deletedTimers, -1)
		case status_ModifiedEarlier, status_ModifiedLater:
			if !atomic.Cas(&t.status, s, status_Moving) {
				continue
			}
			// Now we can change the when field.
			t.when = t.nextwhen
			// Move t to the right position.
			dodeltimer0(pp)
			doaddtimer(pp, t)
			if !atomic.Cas(&t.status, status_Moving, status_Waiting) {
				badTimer()
			}
		default:
			// Head of timers does not need adjustment.
			return
		}
	}
}

// moveTimers moves a slice of timers to pp. The slice has been taken
// from a different P.
// This is currently called when the world is stopped, but the caller
// is expected to have locked the timers for pp.
func moveTimers(pp *p, timers []*timer) {
	for _, t := range timers {
	loop:
		for {
			switch s := atomic.Load(&t.status); s {
			case status_Waiting:
				if !atomic.Cas(&t.status, s, status_Moving) {
					continue
				}
				t.pp = 0
				doaddtimer(pp, t)
				if !atomic.Cas(&t.status, status_Moving, status_Waiting) {
					badTimer()
				}
				break loop
			case status_ModifiedEarlier, status_ModifiedLater:
				if !atomic.Cas(&t.status, s, status_Moving) {
					continue
				}
				t.when = t.nextwhen
				t.pp = 0
				doaddtimer(pp, t)
				if !atomic.Cas(&t.status, status_Moving, status_Waiting) {
					badTimer()
				}
				break loop
			case status_Deleted:
				if !atomic.Cas(&t.status, s, status_Removed) {
					continue
				}
				t.pp = 0
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

// adjusttimers looks through the timers in the current P's heap for
// any timers that have been modified to run earlier, and puts them in
// the correct place in the heap. While looking for those timers,
// it also moves timers that have been modified to run later,
// and removes deleted timers. The caller must have locked the timers for pp.
func adjusttimers(pp *p, now int64) {
	// If we haven't yet reached the time of the first status_ModifiedEarlier
	// timer, don't do anything. This speeds up programs that adjust
	// a lot of timers back and forth if the timers rarely expire.
	// We'll postpone looking through all the adjusted timers until
	// one would actually expire.
	first := atomic.Load64(&pp.timerModifiedEarliest)
	if first == 0 || int64(first) > now {
		if verifyTimers {
			verifyTimerHeap(pp)
		}
		return
	}

	// We are going to clear all status_ModifiedEarlier timers.
	atomic.Store64(&pp.timerModifiedEarliest, 0)

	var moved []*timer
	for i := 0; i < len(pp.timers); i++ {
		t := pp.timers[i]
		if t.pp.ptr() != pp {
			throw("adjusttimers: bad p")
		}
		switch s := atomic.Load(&t.status); s {
		case status_Deleted:
			if atomic.Cas(&t.status, s, status_Removing) {
				changed := dodeltimer(pp, i)
				if !atomic.Cas(&t.status, status_Removing, status_Removed) {
					badTimer()
				}
				atomic.Xadd(&pp.deletedTimers, -1)
				// Go back to the earliest changed heap entry.
				// "- 1" because the loop will add 1.
				i = changed - 1
			}
		case status_ModifiedEarlier, status_ModifiedLater:
			if atomic.Cas(&t.status, s, status_Moving) {
				// Now we can change the when field.
				t.when = t.nextwhen
				// Take t off the heap, and hold onto it.
				// We don't add it back yet because the
				// heap manipulation could cause our
				// loop to skip some other timer.
				changed := dodeltimer(pp, i)
				moved = append(moved, t)
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
		addAdjustedTimers(pp, moved)
	}

	if verifyTimers {
		verifyTimerHeap(pp)
	}
}

// addAdjustedTimers adds any timers we adjusted in adjusttimers
// back to the timer heap.
func addAdjustedTimers(pp *p, moved []*timer) {
	for _, t := range moved {
		doaddtimer(pp, t)
		if !atomic.Cas(&t.status, status_Moving, status_Waiting) {
			badTimer()
		}
	}
}

// nobarrierWakeTime looks at P's timers and returns the time when we
// should wake up the netpoller. It returns 0 if there are no timers.
// This function is invoked when dropping a P, and must run without
// any write barriers.
//go:nowritebarrierrec
func nobarrierWakeTime(pp *p) int64 {
	next := int64(atomic.Load64(&pp.timer0When))
	nextAdj := int64(atomic.Load64(&pp.timerModifiedEarliest))
	if next == 0 || (nextAdj != 0 && nextAdj < next) {
		next = nextAdj
	}
	return next
}

// runtimer examines the first timer in timers. If it is ready based on now,
// it runs the timer and removes or updates it.
// Returns 0 if it ran a timer, -1 if there are no more timers, or the time
// when the first timer should run.
// The caller must have locked the timers for pp.
// If a timer is run, this will temporarily unlock the timers.
//go:systemstack
func runtimer(pp *p, now int64) int64 {
	for {
		t := pp.timers[0]
		if t.pp.ptr() != pp {
			throw("runtimer: bad p")
		}
		switch s := atomic.Load(&t.status); s {
		case status_Waiting:
			if t.when > now {
				// Not ready to run.
				return t.when
			}

			if !atomic.Cas(&t.status, s, status_Running) {
				continue
			}
			// Note that runOneTimer may temporarily unlock
			// pp.timersLock.
			runOneTimer(pp, t, now)
			return 0

		case status_Deleted:
			if !atomic.Cas(&t.status, s, status_Removing) {
				continue
			}
			dodeltimer0(pp)
			if !atomic.Cas(&t.status, status_Removing, status_Removed) {
				badTimer()
			}
			atomic.Xadd(&pp.deletedTimers, -1)
			if len(pp.timers) == 0 {
				return -1
			}

		case status_ModifiedEarlier, status_ModifiedLater:
			if !atomic.Cas(&t.status, s, status_Moving) {
				continue
			}
			t.when = t.nextwhen
			dodeltimer0(pp)
			doaddtimer(pp, t)
			if !atomic.Cas(&t.status, status_Moving, status_Waiting) {
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
// The caller must have locked the timers for pp.
// This will temporarily unlock the timers while running the timer function.
//go:systemstack
func runOneTimer(pp *p, t *timer, now int64) {
	if raceenabled {
		ppcur := getg().m.p.ptr()
		if ppcur.timerRaceCtx == 0 {
			ppcur.timerRaceCtx = racegostart(abi.FuncPCABIInternal(runtimer) + sys.PCQuantum)
		}
		raceacquirectx(ppcur.timerRaceCtx, unsafe.Pointer(t))
	}

	f := t.f
	arg := t.arg
	seq := t.seq

	if t.period > 0 {
		// Leave in heap but adjust next time to fire.
		delta := t.when - now
		t.when += t.period * (1 + -delta/t.period)
		if t.when < 0 { // check for overflow.
			t.when = maxWhen
		}
		siftdownTimer(pp.timers, 0)
		if !atomic.Cas(&t.status, status_Running, status_Waiting) {
			badTimer()
		}
		updateTimer0When(pp)
	} else {
		// Remove from heap.
		dodeltimer0(pp)
		if !atomic.Cas(&t.status, status_Running, status_Unset) {
			badTimer()
		}
	}

	if raceenabled {
		// Temporarily use the current P's racectx for g0.
		gp := getg()
		if gp.racectx != 0 {
			throw("runOneTimer: unexpected racectx")
		}
		gp.racectx = gp.m.p.ptr().timerRaceCtx
	}

	unlock(&pp.timersLock)

	f(arg, seq)

	lock(&pp.timersLock)

	if raceenabled {
		gp := getg()
		gp.racectx = 0
	}
}

// clearDeletedTimers removes all deleted timers from the P's timer heap.
// This is used to avoid clogging up the heap if the program
// starts a lot of long-running timers and then stops them.
// For example, this can happen via context.WithTimeout.
//
// This is the only function that walks through the entire timer heap,
// other than moveTimers which only runs when the world is stopped.
//
// The caller must have locked the timers for pp.
func clearDeletedTimers(pp *p) {
	// We are going to clear all status_ModifiedEarlier timers.
	// Do this now in case new ones show up while we are looping.
	atomic.Store64(&pp.timerModifiedEarliest, 0)

	cdel := int32(0)
	to := 0
	changedHeap := false
	timers := pp.timers
nextTimer:
	for _, t := range timers {
		for {
			switch s := atomic.Load(&t.status); s {
			case status_Waiting:
				if changedHeap {
					timers[to] = t
					siftupTimer(timers, to)
				}
				to++
				continue nextTimer
			case status_ModifiedEarlier, status_ModifiedLater:
				if atomic.Cas(&t.status, s, status_Moving) {
					t.when = t.nextwhen
					timers[to] = t
					siftupTimer(timers, to)
					to++
					changedHeap = true
					if !atomic.Cas(&t.status, status_Moving, status_Waiting) {
						badTimer()
					}
					continue nextTimer
				}
			case status_Deleted:
				if atomic.Cas(&t.status, s, status_Removing) {
					t.pp = 0
					cdel++
					if !atomic.Cas(&t.status, status_Removing, status_Removed) {
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
		timers[i] = nil
	}

	atomic.Xadd(&pp.deletedTimers, -cdel)
	atomic.Xadd(&pp.numTimers, -cdel)

	timers = timers[:to]
	pp.timers = timers
	updateTimer0When(pp)

	if verifyTimers {
		verifyTimerHeap(pp)
	}
}

// verifyTimerHeap verifies that the timer heap is in a valid state.
// This is only for debugging, and is only called if verifyTimers is true.
// The caller must have locked the timers.
func verifyTimerHeap(pp *p) {
	for i, t := range pp.timers {
		if i == 0 {
			// First timer has no parent.
			continue
		}

		// The heap is 4-ary. See siftupTimer and siftdownTimer.
		p := (i - 1) / 4
		if t.when < pp.timers[p].when {
			print("bad timer heap at ", i, ": ", p, ": ", pp.timers[p].when, ", ", i, ": ", t.when, "\n")
			throw("bad timer heap")
		}
	}
	if numTimers := int(atomic.Load(&pp.numTimers)); len(pp.timers) != numTimers {
		println("timer heap len", len(pp.timers), "!= numTimers", numTimers)
		throw("bad timer heap len")
	}
}

// updateTimer0When sets the P's timer0When field.
// The caller must have locked the timers for pp.
func updateTimer0When(pp *p) {
	if len(pp.timers) == 0 {
		atomic.Store64(&pp.timer0When, 0)
	} else {
		atomic.Store64(&pp.timer0When, uint64(pp.timers[0].when))
	}
}

// updateTimerModifiedEarliest updates the recorded nextwhen field of the
// earlier timerModifiedEarier value.
// The timers for pp will not be locked.
func updateTimerModifiedEarliest(pp *p, nextwhen int64) {
	for {
		old := atomic.Load64(&pp.timerModifiedEarliest)
		if old != 0 && int64(old) < nextwhen {
			return
		}
		if atomic.Cas64(&pp.timerModifiedEarliest, old, uint64(nextwhen)) {
			return
		}
	}
}

// timeSleepUntil returns the time when the next timer should fire,
// and the P that holds the timer heap that that timer is on.
// This is only called by sysmon and checkdead.
func timeSleepUntil() (int64, *p) {
	next := int64(maxWhen)
	var pret *p

	// Prevent allp slice changes. This is like retake.
	lock(&allpLock)
	for _, pp := range allp {
		if pp == nil {
			// This can happen if procresize has grown
			// allp but not yet created new Ps.
			continue
		}

		w := int64(atomic.Load64(&pp.timer0When))
		if w != 0 && w < next {
			next = w
			pret = pp
		}

		w = int64(atomic.Load64(&pp.timerModifiedEarliest))
		if w != 0 && w < next {
			next = w
			pret = pp
		}
	}
	unlock(&allpLock)

	return next, pret
}

// Heap maintenance algorithms.
// These algorithms check for slice index errors manually.
// Slice index error can happen if the program is using racy
// access to timers. We don't want to panic here, because
// it will cause the program to crash with a mysterious
// "panic holding locks" message. Instead, we panic while not
// holding a lock.

// siftupTimer puts the timer at position i in the right place
// in the heap by moving it up toward the top of the heap.
// It returns the smallest changed index.
func siftupTimer(t []*timer, i int) int {
	if i >= len(t) {
		badTimer()
	}
	when := t[i].when
	if when <= 0 {
		badTimer()
	}
	tmp := t[i]
	for i > 0 {
		p := (i - 1) / 4 // parent
		if when >= t[p].when {
			break
		}
		t[i] = t[p]
		i = p
	}
	if tmp != t[i] {
		t[i] = tmp
	}
	return i
}

// siftdownTimer puts the timer at position i in the right place
// in the heap by moving it down toward the bottom of the heap.
func siftdownTimer(t []*timer, i int) {
	n := len(t)
	if i >= n {
		badTimer()
	}
	when := t[i].when
	if when <= 0 {
		badTimer()
	}
	tmp := t[i]
	for {
		c := i*4 + 1 // left child
		c3 := c + 2  // mid child
		if c >= n {
			break
		}
		w := t[c].when
		if c+1 < n && t[c+1].when < w {
			w = t[c+1].when
			c++
		}
		if c3 < n {
			w3 := t[c3].when
			if c3+1 < n && t[c3+1].when < w3 {
				w3 = t[c3+1].when
				c3++
			}
			if w3 < w {
				w = w3
				c = c3
			}
		}
		if w >= when {
			break
		}
		t[i] = t[c]
		i = c
	}
	if tmp != t[i] {
		t[i] = tmp
	}
}

// badTimer is called if the timer data structures have been corrupted,
// presumably due to racy use by the program. We panic here rather than
// panicing due to invalid slice access while holding locks.
// See issue #25686.
func badTimer() {
	panic("timer data corruption")
}
