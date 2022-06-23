/* For license and copyright information please see LEGAL file in repository */

package tcp

import "sync/atomic"

const (
	SocketState_Unset SocketState = iota
	// represents waiting for a connection request from any remote TCP and port
	SocketState_LISTEN
	// represents waiting for a matching connection request after having sent a connection request.
	SocketState_SYN_SENT
	// represents waiting for a confirming connection request acknowledgment
	// after having both received and sent a connection request.
	SocketState_SYN_RECEIVED
	/* represents an open connection, data received can be
	delivered to the user.  The normal state for the data transfer phase
	of the connection. */
	SocketState_ESTABLISHED
	// represents waiting for a connection termination request from the remote TCP,
	// or an acknowledgment of the connection termination request previously sent.
	SocketState_FIN_WAIT_1
	// represents waiting for a connection termination request from the remote TCP.
	SocketState_FIN_WAIT_2
	// represents no connection state at all.
	SocketState_CLOSE
	// represents waiting for a connection termination request from the local user.
	SocketState_CLOSE_WAIT
	// represents waiting for a connection termination request acknowledgment from the remote TCP.
	SocketState_CLOSING
	// represents waiting for an acknowledgment of the connection termination request previously sent to the remote TCP
	// which includes an acknowledgment of its connection termination request
	SocketState_LAST_ACK
	// represents waiting for enough time to pass to be sure the remote TCP received the acknowledgment
	// of its connection termination request.
	SocketState_TIME_WAIT

	// SocketState_NEW_SYN_RECV
)

// SocketState use to indicate socket state.
type SocketState uint32

func (s *SocketState) Load() SocketState {
	return SocketState(atomic.LoadUint32((*uint32)(s)))
}
func (s *SocketState) Store(SocketState SocketState) {
	atomic.StoreUint32((*uint32)(s), uint32(SocketState))
}
func (s *SocketState) CompareAndSwap(old, new SocketState) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(s), uint32(old), uint32(new))
}
