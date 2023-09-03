/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
	"memar/time/monotonic"
)

// Stream provide some fields to hold stream states.
// Because each stream methods just call by a fixed worker on same CPU core in sync order, don't need to lock or changed atomic any field
type Stream struct {
	// connection      protocol.Connection
	sk  protocol.Socket
	mtu int
	mss int // Max Segment Length

	// just store last send or receive segment not read or write to.
	lastUse monotonic.Time

	// TODO::: Cookie, save stream in nvm

	timing
	send
	recv
	port

	// Stream use to send or receive data on specific connection.
	// It can pass to logic layer to give data access to developer!
	// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
	// If OutcomePayload not present stream is UnidirectionalStream otherwise it is BidirectionalStream!

	id PortNumber

	/* State */
	err protocol.Error // Decode||Encode by ErrorID
	// state        protocol.NetworkStatus      // States locate in const of this file.
	// stateChannel chan protocol.NetworkStatus // States locate in const of this file.
	weight protocol.Weight // 16 queue for priority weight of the streams exist.

	status
	StreamMetrics
}

// Init use to initialize the stream after allocation in both server or client
//
//memar:impl memar/protocol.ObjectLifeCycle
func (s *Stream) Init(timeout protocol.Duration, cca CCA) (err protocol.Error) {
	// TODO:::
	s.mss = CNF_Segment_MaxSize
	s.status.Init(StreamStatus_Listen)

	if timeout == 0 {
		timeout = CNF_KeepAlive_Idle
	}

	err = s.timing.Init(s)
	if err != nil {
		return
	}
	err = s.recv.Init(timeout)
	if err != nil {
		return
	}
	err = s.send.Init(timeout)
	return
}
func (s *Stream) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (s *Stream) Deinit() (err protocol.Error) {
	// TODO:::
	err = s.timing.Deinit()
	if err != nil {
		return
	}
	err = s.recv.Deinit()
	if err != nil {
		return
	}
	err = s.send.Deinit()
	return
}

// Reset use to reset the stream to store in a sync.Pool to reuse in near future before 2 GC period to dealloc forever
func (s *Stream) Reset() (err protocol.Error) {
	// TODO:::
	err = s.Reinit()
	// TODO:::
	return
}

// Open call when a client want to open the stream on the client side.
func (s *Stream) Open() (err protocol.Error) {
	err = s.sendSYN()
	s.status.Store(StreamStatus_SynSent)
	// TODO::: timer, retry, change status, block on status change until StreamStatus_Established
	return
}

// CloseSending close the sending side of a stream. Much like close except that we don't receive shut down
func (s *Stream) CloseSending() (err protocol.Error) {
	return
}

// Receive Don't hold segment, So caller can reuse packet slice for any purpose.
// It must be non blocking and just route packet not to wait for anything else.
// for each stream upper layer must call by same CPU(core), so we don't need implement any locking mechanism.
// https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/net/ipv4/tcp_ipv4.c#n1965
func (s *Stream) Receive(segment Segment) (err protocol.Error) {
	err = segment.CheckSegment()
	if err != nil {
		return
	}

	// TODO:::

	switch s.status.Load() {
	case StreamStatus_Listen:
		err = s.incomeSegmentOnListenState(segment)
	case StreamStatus_SynSent:
		err = s.incomeSegmentOnSynSentState(segment)
	case StreamStatus_SynReceived:
		err = s.incomeSegmentOnSynReceivedState(segment)
	case StreamStatus_Established:
		err = s.incomeSegmentOnEstablishedState(segment)
	case StreamStatus_FinWait1:
		err = s.incomeSegmentOnFinWait1State(segment)
	case StreamStatus_FinWait2:
		err = s.incomeSegmentOnFinWait2State(segment)
	case StreamStatus_Close:
		err = s.incomeSegmentOnCloseState(segment)
	case StreamStatus_CloseWait:
		err = s.incomeSegmentOnCloseWaitState(segment)
	case StreamStatus_Closing:
		err = s.incomeSegmentOnClosingState(segment)
	case StreamStatus_LastAck:
		err = s.incomeSegmentOnLastAckState(segment)
	case StreamStatus_TimeWait:
		err = s.incomeSegmentOnTimeWaitState(segment)
	}
	return
}

// ScheduleProcessingSocket is Non-Blocking means It must not block the caller in any ways.
// Stream must start with NetworkStatus_NeedMoreData if it doesn't need to call the service when the state changed for the first time
func (st *Stream) ScheduleProcessingSocket() {
	// decide by stream odd or even
	// TODO::: check better performance as "streamID%2 == 0" to check odd id
	// if streamID&1 == 0 {
	// 	// TODO::: easily call by "go" or call by workers pool or what??
	// 	go f.callService(conn, stream)
	// } else {
	// 	// income response
	// 	stream.SetState(protocol.NetworkStatus_Ready)
	// }

	// if st.State == protocol.NetworkStatus_Open {
	// TODO::: easily call by "go" or call by workers pool or what??
	// go st.callService()
	// return
	// }
	// st.SetState(protocol.NetworkStatus_ReceivedCompletely)
}
