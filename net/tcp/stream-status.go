/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"sync/atomic"

	error_p "memar/error/protocol"
)

// streamStatus use to indicate stream state.
type streamStatus uint32

const (
	StreamStatus_Unset streamStatus = iota
	// represents waiting for a connection request from any remote TCP and port
	StreamStatus_Listen
	// represents waiting for a matching connection request after having sent a connection request.
	StreamStatus_SynSent
	// represents waiting for a confirming connection request acknowledgment
	// after having both received and sent a connection request.
	StreamStatus_SynReceived
	/* represents an open connection, data received can be
	delivered to the user.  The normal state for the data transfer phase
	of the connection. */
	StreamStatus_Established
	// represents no connection state at all.
	StreamStatus_Close
	// represents waiting for a connection termination request from the local user.
	StreamStatus_CloseWait
	// represents waiting for a connection termination request acknowledgment from the remote TCP.
	StreamStatus_Closing
	// represents waiting for an acknowledgment of the connection termination request previously sent to the remote TCP
	// which includes an acknowledgment of its connection termination request
	StreamStatus_LastAck
	// represents waiting for enough time to pass to be sure the remote TCP received the acknowledgment
	// of its connection termination request.
	StreamStatus_TimeWait
	// represents waiting for a connection termination request from the remote TCP,
	// or an acknowledgment of the connection termination request previously sent.
	StreamStatus_FinWait1
	// represents waiting for a connection termination request from the remote TCP.
	StreamStatus_FinWait2

	// StreamStatus_NEW_SYN_RECV
)

type status struct {
	ss     atomic.Uint32
	ssChan chan streamStatus
}

// memar/computer/language/object/protocol.LifeCycle
func (s *status) Init(is streamStatus) (err error_p.Error) {
	s.ssChan = make(chan streamStatus)
	s.ss.Store(uint32(is))
	// s.stateChan <- is
	return
}

func (s *status) Load() streamStatus {
	return streamStatus(s.ss.Load())
}
func (s *status) Store(ss streamStatus) {
	s.ss.Store(uint32(ss))
	s.notify(ss)
}
func (s *status) CompareAndSwap(old, new streamStatus) (swapped bool) {
	swapped = s.ss.CompareAndSwap(uint32(old), uint32(new))
	if swapped {
		s.notify(new)
	}
	return
}

func (s *status) notify(ss streamStatus) {
	select {
	case s.ssChan <- ss:
		// state can be delivered by
	default:
		// nothing to do just drop state because channel is block from other
	}
}
