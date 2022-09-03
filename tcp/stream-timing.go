/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/monotonic"
	"github.com/GeniusesGroup/libgo/timer"
)

type timing struct {
	st *Stream
	// TODO::: one timer or many per handler??
	socketTimer timer.Timer

	keepAlive
	delayedAcknowledgment
}

func (t *timing) Init(st *Stream) {
	var now = monotonic.Now()
	var next protocol.Duration

	t.st = st
	t.socketTimer.Init(t)

	if KeepAlive {
		var nxt = t.keepAlive.Init(now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	if DelayedAcknowledgment {
		var nxt = t.delayedAcknowledgment.Init(now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	if next > 0 {
		t.socketTimer.Tick(next, next, -1)
	}
}
func (t *timing) Reinit() {
	if KeepAlive {
		t.keepAlive.Reinit()
	}
	if DelayedAcknowledgment {
		t.delayedAcknowledgment.Reinit()
	}
	t.socketTimer.Stop()
}
func (t *timing) Deinit() {
	if KeepAlive {
		t.keepAlive.Deinit()
	}
	if DelayedAcknowledgment {
		t.delayedAcknowledgment.Deinit()
	}
	t.socketTimer.Stop()
}

// Don't block the caller
func (t *timing) TimerHandler() {
	var next protocol.Duration
	var now = monotonic.Now()
	var st = t.st

	if KeepAlive {
		var nxt = t.keepAlive.CheckInterval(st, now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	if DelayedAcknowledgment {
		var nxt = t.delayedAcknowledgment.CheckInterval(st, now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	// TODO::: add more handler

	if next > 0 {
		t.socketTimer.Reset(protocol.Duration(next))
	}
}
