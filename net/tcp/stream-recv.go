/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/buffer"
	"libgo/protocol"
	"libgo/timer"
)

// recv is receive sequence space
type recv struct {
	readTimer timer.Sync // read deadline timer

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

//libgo:impl libgo/protocol.ObjectLifeCycle
func (r *recv) Init(timeout protocol.Duration) (err protocol.Error) {
	r.flag = make(chan flag, 1) // 1 buffer slot??

	err = r.readTimer.Init()
	err = r.readTimer.Start(timeout)

	// TODO:::
	return
}
func (r *recv) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (r *recv) Deinit() (err protocol.Error) {
	// TODO:::
	err = r.readTimer.Deinit()
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
