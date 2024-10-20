/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"memar/protocol"
)

// Active timers live in the timers field as heap structure.
// Inactive timers live there too temporarily, until they are removed.
// Balancing a heap is done by th.siftUp or th.siftDown methods
//
// Normally access the timers while running on same CPU core,
// but the scheduler can also do it from a different CPU core,
// Anyway caller MUST decide and MAY use any sync algorithm to sync operations.
//
// https://en.wikipedia.org/wiki/Heap_(data_structure)#Comparison_of_theoretic_bounds_for_variants
type timingHeap struct {
	timers []timerBucketHeap
}

//memar:impl memar/protocol.SoftwareLifeCycle
func (th *timingHeap) Init() (err protocol.Error) {
	// TODO::: let application flow choose timers init cap or force it?
	// th.timers = make([]timerBucketHeap, 1024)
	return
}
func (th *timingHeap) Reinit() (err protocol.Error) {
	// TODO::: Do timers??
	th.timers = th.timers[:0]
	return
}
func (th *timingHeap) Deinit() (err protocol.Error) {
	// th.timers = nil
	return
}

//memar:impl memar/protocol.ADT_LastElementIndex
func (th *timingHeap) LastElementIndex() protocol.ElementIndex {
	return protocol.ElementIndex(th.OccupiedLength() - 1)
}

//memar:impl memar/protocol.OccupiedLength
func (th *timingHeap) OccupiedLength() int { return len(th.timers) }

func (th *timingHeap) Append(b timerBucketHeap) { th.timers = append(th.timers, b) }

// DeleteTimer removes timer i from the timers heap.
// It returns the smallest changed index in the timingHeap
func (th *timingHeap) DeleteTimer(i int) (smallestChanged int) {
	th.timers[i].timer.timing = nil

	var last = int(th.LastElementIndex())
	if i != last {
		th.timers[i] = th.timers[last]
	}
	th.timers[last].timer = nil
	th.timers = th.timers[:last]

	smallestChanged = i
	if i != last {
		// Moving to i may have moved the last timer to a new parent,
		// so sift up to preserve the heap guarantee.
		smallestChanged = th.SiftUpTimer(i)
		th.SiftDownTimer(i)
	}

	return
}

// DeleteTimer0 removes timer 0 from the timers heap.
// It reports whether it saw no problems due to races.
func (th *timingHeap) DeleteTimer0() {
	th.timers[0].timer.timing = nil

	var last = th.LastElementIndex()
	if last > 0 {
		th.timers[0] = th.timers[last]
	}
	th.timers[last].timer = nil
	th.timers = th.timers[:last]
	if last > 0 {
		th.SiftDownTimer(0)
	}
}

// SiftUpTimer puts the timer at position i in the right place
// in the heap by moving it up toward the top of the heap.
// It returns the smallest changed index.
func (th *timingHeap) SiftUpTimer(i int) int {
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

// SiftDownTimer puts the timer at position i in the right place
// in the heap by moving it down toward the bottom of the heap.
func (th *timingHeap) SiftDownTimer(i int) {
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
