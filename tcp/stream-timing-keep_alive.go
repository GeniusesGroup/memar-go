/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/monotonic"
)

type keepAlive struct {
	enable     bool
	idle       protocol.Duration
	interval   protocol.Duration
	lastUse    monotonic.Time
	nextCheck  monotonic.Time
	retryCount int
}

func (ka *keepAlive) Init(now monotonic.Time) (next protocol.Duration) {
	ka.enable = true
	ka.idle = KeepAlive_Idle
	ka.interval = KeepAlive_Interval
	now.Add(KeepAlive_Interval)
	ka.nextCheck = now
	return KeepAlive_Interval
}
func (ka *keepAlive) Reinit() {
	ka.enable = false
	ka.interval = 0
	ka.nextCheck = 0
	ka.retryCount = 0
}
func (ka *keepAlive) Deinit() {}

// Don't block the caller
func (ka *keepAlive) CheckInterval(st *Stream, now monotonic.Time) (next protocol.Duration) {
	if !ka.enable {
		return -1
	}

	// TODO::: check stream state

	// check last use of stream and compare with our state
	if ka.lastUse != st.lastUse {
		ka.lastUse = st.lastUse
		now.Add(ka.idle)
		ka.nextCheck = now
		ka.retryCount = 0
		next = ka.idle
		return
	}

	next = ka.next(now)
	if next > 0 {
		return
	}
	if next < 0 {
		// calculate timer delta because timer start after nextCheck.
		if ka.retryCount == 0 {
			next = ka.idle + next // + because we have negative next
		} else {
			next = ka.interval + next // + because we have negative next
		}
	}
	ka.nextCheck.Add(next)

	// if (tp->packets_out || !tcp_write_queue_empty(sk))

	if ka.retryCount <= KeepAlive_Probes {
		// TODO::: send ack segment (keepalive message)
		ka.retryCount++
	} else {
		// TODO::: kill the stream
		// return
	}
	return
}

// d can be negative that show ka.CheckInterval called with some delay
func (ka *keepAlive) next(now monotonic.Time) (d protocol.Duration) {
	d = ka.nextCheck.Until(now)
	return
}
