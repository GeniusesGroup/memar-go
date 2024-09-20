/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	error_p "memar/error/protocol"
	"memar/time/duration"
	"memar/time/monotonic"
	"memar/timer"
)

type timing struct {
	st *Stream
	// TODO::: one timer or many per handler or two timer for high accurate and low one??
	streamTimer timer.Async

	ka timingKeepAlive
	de delayedAcknowledgment
}

// memar/computer/language/object/protocol.LifeCycle
func (t *timing) Init(st *Stream) (err error_p.Error) {
	var now = monotonic.Now()
	var next duration.NanoSecond

	t.st = st

	if CNF_KeepAlive {
		var nxt duration.NanoSecond
		nxt, err = t.ka.Init(now)
		if nxt > 0 && nxt < next {
			next = nxt
		}
	}

	if CNF_DelayedAcknowledgment {
		var nxt duration.NanoSecond
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
func (t *timing) Reinit() (err error_p.Error) {
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
func (t *timing) Deinit() (err error_p.Error) {
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
	var next duration.NanoSecond
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
		t.streamTimer.Reset(duration.NanoSecond(next))
	}
}
