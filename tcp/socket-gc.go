/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/monotonic"
	"github.com/GeniusesGroup/libgo/timer"
)

type gc struct {
	socketTimer             timer.Timer
	keepAlive_Interval_next int64
}

func (gc *gc) init() {
	gc.socketTimer.Init(gc)
	// TODO:::
	gc.socketTimer.Start()
}

func (gc *gc) checkTimeout() {
	// TODO:::
}

// Don't block the caller
func (gc *gc) TimerHandler() {
	var next int64
	var now = monotonic.RuntimeNano()

	// TODO:::

	if KeepAlive_Interval > 0 {
		var nxt = gc.checkKeepAliveInterval(now)
		if nxt < next {
			next = nxt
		}
	}

	// TODO:::

	gc.socketTimer.Reset(protocol.Duration(next))
}

func (gc *gc) checkKeepAliveInterval(now int64) (next int64) {
	if now-gc.keepAlive_Interval_next > 0 {
		next = now + KeepAlive_Interval
		gc.keepAlive_Interval_next = next
		// TODO::: send keepalive message
	}
	return
}
