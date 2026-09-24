/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	error_p "memar/process/error/protocol"
	"memar/time/monotonic"
)

type timerBucketHeap struct {
	timer *Async
	// Two reason to have timer when here:
	// - hot cache to prevent dereference timer to get when field
	// - It can be difference with timer when filed in timerModifiedXX status.
	when monotonic.Time
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *timerBucketHeap) Init() (err error_p.Error)   { return }
func (self *timerBucketHeap) Reinit() (err error_p.Error) { return }
func (self *timerBucketHeap) Deinit() (err error_p.Error) { self.timer = nil; self.when = 0; return }
