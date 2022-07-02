/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"../buffer"
)

// Receive Sequence Space
//
// 1          2          3
// ----------|----------|----------
//    RCV.NXT    RCV.NXT
// 			 +RCV.WND
//
// 1 - old sequence numbers which have been acknowledged
// 2 - sequence numbers allowed for new reception
// 3 - future sequence numbers which are not yet allowed
type recvSequenceSpace struct {
	next uint32 // receive next
	wnd  uint16 // receive window
	up   bool   // receive urgent pointer
	irs  uint32 // initial receive sequence number
	// TODO::: not in order segments
	buf buffer.Queue

	// TODO::: Send more than these flags: push, reset, finish, urgent
	flag chan flag
}

func (r *recvSequenceSpace) init() {
	r.flag = make(chan flag, 1)
	// TODO:::
}

// sendFlagSignal use to notify listener in the r.flag channel
func (r *recvSequenceSpace) sendFlagSignal(f flag) {
	select {
	case r.flag <- f:
		// nothing to do
	default:
		break
	}
}
