/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/buffer"
	"memar/protocol"
)

// recv is receive sequence space
type recv struct {
	next uint32 // receive next
	wnd  uint16 // receive window
	up   bool   // receive urgent pointer
	irs  uint32 // initial receive sequence number
	// TODO::: not in order segments
	buf buffer.Queue

	// TODO::: Send more than these flags: push, reset, finish, urgent
	// TODO::: byte is not enough here to distinguish between flags in first byte or second one
	flag chan flag
}

//memar:impl memar/protocol.ObjectLifeCycle
func (r *recv) Init() (err protocol.Error) {
	r.flag = make(chan flag, 1) // 1 buffer slot??


	// TODO:::
	return
}
func (r *recv) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (r *recv) Deinit() (err protocol.Error) {
	// TODO:::
	return
}

// sendFlagSignal use to notify listener in the r.flag channel
func (r *recv) sendFlagSignal(f flag) {
	select {
	case r.flag <- f:
		// nothing to do
	default:
		break
	}
}
