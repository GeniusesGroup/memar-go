/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"sync"

	"memar/protocol"
)

// Active timers live in the timers field as heap structure.
// Inactive timers live there too temporarily, until they are removed.
// Balancing a heap is done by th.siftUp or th.siftDown methods
//
// Normally access the timers while running on same CPU core,
// but the scheduler can also do it from a different CPU core,
// So in this case caller must hold Lock() to access.
//
// https://en.wikipedia.org/wiki/Heap_(data_structure)#Comparison_of_theoretic_bounds_for_variants
type timingHeap struct {
	sync.Mutex
	timers []timerBucketHeap
}

//memar:impl memar/protocol.SoftwareLifeCycle
func (th *timingHeap) Init() (err protocol.Error) {
	// TODO::: let application flow choose timers init cap or force it?
	// th.timers = make([]timerBucketHeap, 1024)
	return
}
func (th *timingHeap) Reinit() (err protocol.Error) {
	th.timers = nil
	return
}
func (th *timingHeap) Deinit() (err protocol.Error) {
	// th.timers = nil
	return
}

func (th *timingHeap) Len() int                 { return len(th.timers) }
func (th *timingHeap) Append(b timerBucketHeap) { th.timers = append(th.timers, b) }

// deleteTimer removes timer i from the timers heap.
// It returns the smallest changed index in the timingHeap
// The caller MUST have locked the timingHeap
func (th *timingHeap) DeleteTimer(i int) (smallestChanged int) {
	th.timers[i].timer.timing = nil

	var last = th.Len() - 1
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

// deleteTimer0 removes timer 0 from the timers heap.
// It reports whether it saw no problems due to races.
// The caller MUST have locked the timingHeap
func (th *timingHeap) DeleteTimer0() {
	th.timers[0].timer.timing = nil

	var last = th.Len() - 1
	if last > 0 {
		th.timers[0] = th.timers[last]
	}
	th.timers[last].timer = nil
	th.timers = th.timers[:last]
	if last > 0 {
		th.SiftDownTimer(0)
	}
}

// Heap maintenance algorithms.
// These algorithms check for slice index errors manually.
// Slice index error can happen if the program is using racy
// access to timers. We don't want to panic here, because
// it will cause the program to crash with a mysterious
// "panic holding locks" message. Instead, we panic while not
// holding a lock.

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
