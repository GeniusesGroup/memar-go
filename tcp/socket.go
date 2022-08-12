/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// Socket provide some fields to hold socket state.
// Because each socket methods just call by a fixed worker on same CPU core in sync order, don't need to lock or changed atomic any field
type Socket struct {
	connection      protocol.Connection
	stream          protocol.Stream
	mtu             int
	mss             int    // Max Segment Length
	sourcePort      uint16 // local
	destinationPort uint16 // remote
	status          SocketState
	state           chan SocketState

	// TODO::: Cookie, save socket in nvm

	timing
	send
	recv
}

// Init use to initialize the socket after allocation in both server or client
func (s *Socket) Init(timeout protocol.Duration) {
	// TODO:::
	s.mss = OptionDefault_MSS
	s.setState(SocketState_LISTEN)

	if timeout == 0 {
		timeout = KeepAlive_Idle
	}

	s.timing.init()
	s.recv.init(timeout)
	s.send.init(timeout)
}

func (s *Socket) Connection() protocol.Connection { return s.connection }
func (s *Socket) Stream() protocol.Stream         { return s.stream }

// Reset use to reset the socket to store in a sync.Pool to reuse in near future before 2 GC period to dealloc forever
func (s *Socket) Reset() {
	// TODO:::
}

// Open call when a client want to open the socket on the client side.
func (s *Socket) Open() (err protocol.Error) {
	err = s.sendSYN()
	s.setState(SocketState_SYN_SENT)
	// TODO::: timer, retry, change status, block on status change until SocketState_ESTABLISHED
	return
}

// CloseSending shutdown the sending side of a socket. Much like close except that we don't receive shut down
func (s *Socket) CloseSending() (err protocol.Error) {
	return
}

func (s *Socket) SetTimeout(d protocol.Duration) (err protocol.Error) {
	err = s.SetReadTimeout(d)
	if err != nil {
		return
	}
	err = s.SetWriteTimeout(d)
	return
}
func (s *Socket) SetReadTimeout(d protocol.Duration) (err protocol.Error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	if d < 0 {
		// no timeout
		// TODO::: is it ok??
		s.recv.readTimer.Stop()
		return
	}
	s.recv.readTimer.Reset(d)
	return
}
func (s *Socket) SetWriteTimeout(d protocol.Duration) (err protocol.Error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	if d < 0 {
		// no timeout
		s.send.writeTimer.Stop()
		return
	}
	s.send.writeTimer.Reset(d)
	return
}

// Receive Don't hold segment, So caller can reuse packet slice for any purpose.
// It must be non blocking and just route packet not to wait for anything else.
// for each socket upper layer must call by same CPU(core), so we don't need implement any locking mechanism.
// https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/net/ipv4/tcp_ipv4.c#n1965
func (s *Socket) Receive(segment Packet) (err protocol.Error) {
	err = segment.CheckPacket()
	if err != nil {
		return
	}

	// TODO:::

	switch s.status {
	case SocketState_LISTEN:
		err = s.incomeSegmentOnListenState(segment)
	case SocketState_SYN_SENT:
		err = s.incomeSegmentOnSynSentState(segment)
	case SocketState_SYN_RECEIVED:
		err = s.incomeSegmentOnSynReceivedState(segment)
	case SocketState_ESTABLISHED:
		err = s.incomeSegmentOnEstablishedState(segment)
	case SocketState_FIN_WAIT_1:
		err = s.incomeSegmentOnFinWait1State(segment)
	case SocketState_FIN_WAIT_2:
		err = s.incomeSegmentOnFinWait2State(segment)
	case SocketState_CLOSE:
		err = s.incomeSegmentOnCloseState(segment)
	case SocketState_CLOSE_WAIT:
		err = s.incomeSegmentOnCloseWaitState(segment)
	case SocketState_CLOSING:
		err = s.incomeSegmentOnClosingState(segment)
	case SocketState_LAST_ACK:
		err = s.incomeSegmentOnLastAckState(segment)
	case SocketState_TIME_WAIT:
		err = s.incomeSegmentOnTimeWaitState(segment)
	}
	return
}
