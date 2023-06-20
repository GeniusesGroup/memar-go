/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/protocol"
	"libgo/time/monotonic"
	"libgo/timer"
)

type timing struct {
	st *Stream
	// TODO::: one timer or many per handler or two timer for high accurate and low one??
	streamTimer timer.Async

	ka timingKeepAlive
	de delayedAcknowledgment
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (t *timing) Init(st *Stream) (err protocol.Error) {
	var now = monotonic.Now()
	var next protocol.Duration

	t.st = st

	if CNF_KeepAlive {
		var nxt protocol.Duration
		nxt, err = t.ka.Init(now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	if CNF_DelayedAcknowledgment {
		var nxt protocol.Duration
		nxt, err = t.de.Init(now)
		if err != nil {
			return
		}
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	if next > 0 {
		err = t.streamTimer.Init(t)
		err = t.streamTimer.Tick(next, next)
	}
	return
}
func (t *timing) Reinit() (err protocol.Error) {
	if CNF_KeepAlive {
		err = t.ka.Reinit()
		if err != nil {
			return
		}
	}
	if CNF_DelayedAcknowledgment {
		err = t.de.Reinit()
		if err != nil {
			return
		}
	}
	err = t.streamTimer.Stop()
	return
}
func (t *timing) Deinit() (err protocol.Error) {
	if CNF_KeepAlive {
		err = t.ka.Deinit()
		if err != nil {
			return
		}
	}
	if CNF_DelayedAcknowledgment {
		err = t.de.Deinit()
		if err != nil {
			return
		}
	}
	err = t.streamTimer.Stop()
	return
}

// Don't block the caller
func (t *timing) TimerHandler() {
	var next protocol.Duration
	var now = monotonic.Now()
	var st = t.st

	if CNF_KeepAlive {
		var nxt = t.ka.CheckInterval(st, now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	if CNF_DelayedAcknowledgment {
		var nxt = t.de.CheckInterval(st, now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	// TODO::: add more handler

	if next > 0 {
		t.streamTimer.Reset(protocol.Duration(next))
	}
}
