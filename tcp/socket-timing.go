/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/monotonic"
	"github.com/GeniusesGroup/libgo/timer"
)

type timing struct {
	socketTimer timer.Timer

	keepAlive_Interval_next    monotonic.Time
	delayedAcknowledgment_next monotonic.Time

	config
}

func (t *timing) init() {
	var next = t.config.init()
	t.socketTimer.Init(t)
	t.socketTimer.Tick(next, next, -1)
}

func (t *timing) deinit() {
	t.socketTimer.Stop()
}

type config struct {
	keepAlive_Interval            protocol.Duration
	delayedAcknowledgment_Timeout protocol.Duration
}

func (c *config) init() (next protocol.Duration) {
	c.keepAlive_Interval = KeepAlive_Interval
	// TODO::: next?
	return
}

// Don't block the caller
func (t *timing) TimerHandler() {
	var next protocol.Duration
	var now = monotonic.Now()

	// TODO:::

	if t.keepAlive_Interval > 0 {
		t.checkKeepAliveInterval(now)
		var nxt = t.keepAlive_Interval
		if nxt < next {
			next = nxt
		}
	}

	if t.delayedAcknowledgment_Timeout > 0 {
		var nxt = t.checkDelayedAcknowledgmentInterval(now)
		if nxt < next {
			next = nxt
		}
	}

	// TODO:::

	t.socketTimer.Reset(protocol.Duration(next))
}

func (t *timing) checkKeepAliveInterval(now monotonic.Time) {
	if t.keepAlive_Interval_next.Pass(now) {
		next = now + monotonic.Time(t.keepAlive_Interval)
		t.keepAlive_Interval_next = next
		// TODO::: send keepalive message
	}
	return
}

func (t *timing) checkDelayedAcknowledgmentInterval(now monotonic.Time) (next monotonic.Time) {

	return
}
