/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	error_p "memar/error/protocol"
	"memar/time/duration"
	"memar/time/monotonic"
)

type delayedAcknowledgment struct {
	enable bool
	// The duration of the delayed-acknowledgement (and persist timers) depends
	// on the measured round trip time of the connection
	interval  duration.NanoSecond
	nextCheck monotonic.Time
}

// memar/computer/language/object/protocol.LifeCycle
func (da *delayedAcknowledgment) Init(now monotonic.Time) (next duration.NanoSecond, err error_p.Error) {
	da.enable = true
	da.interval = CNF_DelayedAcknowledgment_Timeout
	now.Add(CNF_DelayedAcknowledgment_Timeout)
	da.nextCheck = now
	return
}
func (da *delayedAcknowledgment) Reinit() (err error_p.Error) {
	return
}
func (da *delayedAcknowledgment) Deinit() (err error_p.Error) {
	return
}

func (da *delayedAcknowledgment) Enabled() bool { return da.enable }

// Don't block the caller
func (da *delayedAcknowledgment) CheckInterval(s *Stream, now monotonic.Time) (next duration.NanoSecond) {
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
func (da *delayedAcknowledgment) next(now monotonic.Time) (d duration.NanoSecond) {
	d = da.nextCheck.Until(now)
	return
}
