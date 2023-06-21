/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"libgo/protocol"
	"libgo/time/monotonic"
)

type timerBucketHeap struct {
	timer *Async
	// Two reason to have timer when here:
	// - hot cache to prevent dereference timer to get when field
	// - It can be difference with timer when filed in timerModifiedXX status.
	when monotonic.Time
}

//libgo:impl libgo/protocol.SoftwareLifeCycle
func (tb *timerBucketHeap) Init() (err protocol.Error)   { return }
func (tb *timerBucketHeap) Reinit() (err protocol.Error) { return }
func (tb *timerBucketHeap) Deinit() (err protocol.Error) { tb.timer = nil; tb.when = 0; return }
