/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/process/error/protocol"
)

// Active timers live in the timers field as heap structure.
// Inactive timers live there too temporarily, until they are removed.
// Balancing a heap is done by timingHeap.siftUp() or timingHeap.siftDown() methods
//
// Normally access the timers while running on same CPU core,
// but the scheduler can also do it from a different CPU core,
// Anyway caller MUST decide and MAY use any sync algorithm to sync operations.
//
// https://en.wikipedia.org/wiki/Heap_(data_structure)#Comparison_of_theoretic_bounds_for_variants
type timingHeap struct {
	timers []timerBucketHeap
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *timingHeap) Init() (err error_p.Error) {
	// TODO::: let application flow choose timers init cap or force it?
	// self.timers = make([]timerBucketHeap, 1024)
	return
}
func (self *timingHeap) Reinit() (err error_p.Error) {
	// TODO::: Do timers??
	self.timers = self.timers[:0]
	return
}
func (self *timingHeap) Deinit() (err error_p.Error) {
	// self.timers = nil
	return
}

//memar:impl memar/adt/container/protocol.LastElementIndex
func (self *timingHeap) LastElementIndex() container_p.ElementIndex {
	return container_p.ElementIndex(self.OccupiedLength() - 1)
}

//memar:impl memar/adt/protocol.OccupiedLength
func (self *timingHeap) OccupiedLength() int /* container_p.NumberOfElement */ { return len(self.timers) }

func (self *timingHeap) Append(b timerBucketHeap) { self.timers = append(self.timers, b) }

// DeleteTimer removes timer i from the timers heap.
// It returns the smallest changed index in the timingHeap
func (self *timingHeap) DeleteTimer(i int) (smallestChanged int) {
	self.timers[i].timer.timing = nil

	var last = int(self.LastElementIndex())
	if i != last {
		self.timers[i] = self.timers[last]
	}
	self.timers[last].timer = nil
	self.timers = self.timers[:last]

	smallestChanged = i
	if i != last {
		// Moving to i may have moved the last timer to a new parent,
		// so sift up to preserve the heap guarantee.
		smallestChanged = self.SiftUpTimer(i)
		self.SiftDownTimer(i)
	}

	return
}

// DeleteTimer0 removes timer 0 from the timers heap.
// It reports whether it saw no problems due to races.
func (self *timingHeap) DeleteTimer0() {
	self.timers[0].timer.timing = nil

	var last = self.LastElementIndex()
	if last > 0 {
		self.timers[0] = self.timers[last]
	}
	self.timers[last].timer = nil
	self.timers = self.timers[:last]
	if last > 0 {
		self.SiftDownTimer(0)
	}
}

// SiftUpTimer puts the timer at position i in the right place
// in the heap by moving it up toward the top of the heap.
// It returns the smallest changed index.
func (self *timingHeap) SiftUpTimer(i int) int {
	var timers = self.timers
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
func (self *timingHeap) SiftDownTimer(i int) {
	var timers = self.timers
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
