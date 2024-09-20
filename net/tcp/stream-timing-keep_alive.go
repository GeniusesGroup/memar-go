/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	error_p "memar/error/protocol"
	"memar/time/duration"
	"memar/time/monotonic"
)

type timingKeepAlive struct {
	enable     bool
	idle       duration.NanoSecond
	interval   duration.NanoSecond
	lastUse    monotonic.Time
	nextCheck  monotonic.Time
	retryCount int
}

// memar/computer/language/object/protocol.LifeCycle
func (ka *timingKeepAlive) Init(now monotonic.Time) (next duration.NanoSecond, err error_p.Error) {
	ka.enable = CNF_KeepAlive_PerStream
	ka.idle = CNF_KeepAlive_Idle
	ka.interval = CNF_KeepAlive_Interval
	now.Add(CNF_KeepAlive_Idle)
	ka.nextCheck = now

	next = CNF_KeepAlive_Idle
	return
}
func (ka *timingKeepAlive) Reinit() (err error_p.Error) {
	ka.enable = CNF_KeepAlive_PerStream
	ka.interval = 0
	ka.nextCheck = 0
	ka.retryCount = 0
	return
}
func (ka *timingKeepAlive) Deinit() (err error_p.Error) {
	return
}

func (ka *timingKeepAlive) Enable() bool                  { return ka.enable }
func (ka *timingKeepAlive) Idle() duration.NanoSecond     { return ka.idle }
func (ka *timingKeepAlive) Interval() duration.NanoSecond { return ka.interval }

func (ka *timingKeepAlive) SetEnable(keepalive bool) {
	// TODO::: rfc: uncomment below??
	// var now = monotonic.Now()
	// now.Add(CNF_KeepAlive_Idle)
	// ka.nextCheck = now
	ka.enable = keepalive
}
func (ka *timingKeepAlive) SetIdle(d duration.NanoSecond) {
	// TODO::: atomic or normal access??
}
func (ka *timingKeepAlive) SetInterval(d duration.NanoSecond) {
	// TODO::: atomic or normal access??
}

// Don't block the caller
func (ka *timingKeepAlive) CheckInterval(st *Stream, now monotonic.Time) (next duration.NanoSecond) {
	if !ka.enable {
		return -1
	}

	next = ka.nextCheck.Until(now)
	if next > 0 {
		return
	}

	var streamStatus = st.status.Load()
	// TODO::: check other stream status??
	switch streamStatus {
	case StreamStatus_SynSent, StreamStatus_Close:
		// TODO::: check rfc.
		return -1
	case StreamStatus_Established:
		// Nothing to do, just continue.
	default:
		return -1
	}

	// check last use of stream and compare with our state
	if ka.lastUse != st.lastUse {
		ka.lastUse = st.lastUse
		now.Add(ka.idle)
		ka.nextCheck = now
		ka.retryCount = 0
		next = ka.idle
		return
	}

	// next can be negative that show ka.CheckInterval called with some delay
	if next < 0 {
		// calculate timer delta because timer start after nextCheck.
		// TODO::: why need add delta??
		next = ka.interval + next // + because we have negative next
	}
	ka.nextCheck.Add(next)

	// if (tp->packets_out || !tcp_write_queue_empty(sk))

	if ka.retryCount <= CNF_KeepAlive_Probes {
		// TODO::: send ack segment (keepalive message)
		var err = st.sendQuickACK()
		if err != nil {
			// TODO:::
		}
		ka.retryCount++
	} else {
		var err = st.close()
		if err != nil {
			// TODO:::
		}
		return
	}
	return
}
