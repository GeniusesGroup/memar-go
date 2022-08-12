/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/buffer"
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/timer"
)

// recv is receive sequence space
type recv struct {
	readTimer timer.Timer // read deadline timer

	next uint32 // receive next
	wnd  uint16 // receive window
	up   bool   // receive urgent pointer
	irs  uint32 // initial receive sequence number
	// TODO::: not in order segments
	buf buffer.Queue

	// TODO::: Send more than these flags: push, reset, finish, urgent
	flag chan flag
}

func (r *recv) init(timeout protocol.Duration) {
	r.flag = make(chan flag, 1) // 1 buffer slot??

	r.readTimer.Init(nil)
	r.readTimer.Start(timeout)

	// TODO:::
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
