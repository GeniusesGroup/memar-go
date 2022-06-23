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

	pushFlag chan struct{}

	status recvSequenceSpaceState
	state  chan recvSequenceSpaceState
}

func (r *recvSequenceSpace) sendPushFlagSignal() {
	select {
	case r.pushFlag <- struct{}{}:
		// nothing to do
	default:
		break
	}
}
