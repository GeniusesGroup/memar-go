/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/protocol"
	"libgo/time/monotonic"
)

type delayedAcknowledgment struct {
	enable bool
	// The duration of the delayed-acknowledgement (and persist timers) depends
	// on the measured round trip time of the connection
	interval  protocol.Duration
	nextCheck monotonic.Time
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (da *delayedAcknowledgment) Init(now monotonic.Time) (next protocol.Duration, err protocol.Error) {
	da.enable = true
	da.interval = CNF_DelayedAcknowledgment_Timeout
	now.Add(CNF_DelayedAcknowledgment_Timeout)
	da.nextCheck = now
	return
}
func (da *delayedAcknowledgment) Reinit() (err protocol.Error) {
	return
}
func (da *delayedAcknowledgment) Deinit() (err protocol.Error) {
	return
}

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
