/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/monotonic"
)

type delayedAcknowledgment struct {
	enable    bool
	interval  protocol.Duration
	nextCheck monotonic.Time
}

func (da *delayedAcknowledgment) Init(now monotonic.Time) (next protocol.Duration) {
	da.enable = true
	now.Add(DelayedAcknowledgment_Timeout)
	da.nextCheck = now
	return
}
func (da *delayedAcknowledgment) Reinit() {
}
func (da *delayedAcknowledgment) Deinit() {}

// Don't block the caller
func (da *delayedAcknowledgment) CheckInterval(s *Stream, now monotonic.Time) (next protocol.Duration) {
	if !da.enable {
		return -1
	}

	// TODO::: check stream state

	next = da.next(now)
	if next > 0 {
		return
	}

	// TODO:::

	return
}

// d can be negative that show ka.CheckInterval called with some delay
func (da *delayedAcknowledgment) next(now monotonic.Time) (d protocol.Duration) {
	d = da.nextCheck.Until(now)
	return
}
