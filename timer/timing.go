/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"sync/atomic"

	"../cpu"
)

// Timing ...
type Timing struct {
	coreID uint64 // CPU core number this heap run on it

	// The when field of the first entry on the timer heap.
	// This is updated using atomic functions.
	// This is 0 if the timer heap is empty.
	timer0When atomic.Int64

	// The earliest known when field of a timer with
	// timerModifiedEarlier status. Because the timer may have been
	// modified again, there need not be any timer with this value.
	// This is updated using atomic functions.
	// This is 0 if there are no timerModifiedEarlier timers.
	timerModifiedEarliest atomic.Int64

	// Number of timers in P's heap.
	// Modified using atomic instructions.
	numTimers atomic.Int32

	// Number of timerDeleted timers in P's heap.
	// Modified using atomic instructions.
	deletedTimers atomic.Int32

	// Race context used while executing timer functions.
	timerRaceCtx uintptr
}

func (th *Timing) init() {
	th.coreID = cpu.ActiveCoreID()
}

// updateTimerModifiedEarliest updates the th.timerModifiedEarliest value.
// The timers will not be locked.
func (th *Timing) updateTimerModifiedEarliest(nextWhen int64) {
	for {
		var old = th.timerModifiedEarliest.Load()
		if old != 0 && int64(old) < nextWhen {
			return
		}
		if th.timerModifiedEarliest.CompareAndSwap(old, nextWhen) {
			return
		}
	}
}

// sleepUntil returns the time when the next timer should fire.
func (th *Timing) sleepUntil() (until int64) {
	until = int64(maxWhen)

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
func (th *Timing) noBarrierWakeTime() (until int64) {
	until = th.timer0When.Load()
	var nextAdj = th.timerModifiedEarliest.Load()
	if until == 0 || (nextAdj != 0 && nextAdj < until) {
		until = nextAdj
	}
	return
}
