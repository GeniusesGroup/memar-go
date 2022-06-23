/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"time"

	"../protocol"
)

/*
Because each socket methods just call by a fixed worker on same CPU core in sync order, don't need to lock or changed atomic any field

                              +---------+ ---------\      active OPEN
                              |  CLOSED |            \    -----------
                              +---------+<---------\   \   create TCB
                                |     ^              \   \  snd SYN
                   passive OPEN |     |   CLOSE        \   \
                   ------------ |     | ----------       \   \
                    create TCB  |     | delete TCB         \   \
                                V     |                      \   \
                              +---------+            CLOSE    |    \
                              |  LISTEN |          ---------- |     |
                              +---------+          delete TCB |     |
                   rcv SYN      |     |     SEND              |     |
                  -----------   |     |    -------            |     V
 +---------+      snd SYN,ACK  /       \   snd SYN          +---------+
 |         |<-----------------           ------------------>|         |
 |   SYN   |                    rcv SYN                     |   SYN   |
 |   RCVD  |<-----------------------------------------------|   SENT  |
 |         |                    snd ACK                     |         |
 |         |------------------           -------------------|         |
 +---------+   rcv ACK of SYN  \       /  rcv SYN,ACK       +---------+
   |           --------------   |     |   -----------
   |                  x         |     |     snd ACK
   |                            V     V
   |  CLOSE                   +---------+
   | -------                  |  ESTAB  |
   | snd FIN                  +---------+
   |                   CLOSE    |     |    rcv FIN
   V                  -------   |     |    -------
 +---------+          snd FIN  /       \   snd ACK          +---------+
 |  FIN    |<-----------------           ------------------>|  CLOSE  |
 | WAIT-1  |------------------                              |   WAIT  |
 +---------+          rcv FIN  \                            +---------+
   | rcv ACK of FIN   -------   |                            CLOSE  |
   | --------------   snd ACK   |                           ------- |
   V        x                   V                           snd FIN V
 +---------+                  +---------+                   +---------+
 |FINWAIT-2|                  | CLOSING |                   | LAST-ACK|
 +---------+                  +---------+                   +---------+
   |                rcv ACK of FIN |                 rcv ACK of FIN |
   |  rcv FIN       -------------- |    Timeout=2MSL -------------- |
   |  -------              x       V    ------------        x       V
    \ snd ACK                 +---------+delete TCB         +---------+
     ------------------------>|TIME WAIT|------------------>| CLOSED  |
                              +---------+                   +---------+
*/
type Socket struct {
	Connection      protocol.Connection
	Stream          protocol.Stream
	mtu             int
	mss             int // Max Segment Length
	sourcePort      uint16
	destinationPort uint16
	status          SocketState
	state           chan SocketState

	// TODO::: Cookie, save socket in nvm

	socketDeadline time.Time
	readDeadline   time.Timer
	writeDeadline  time.Timer

	// Rx means Receive, and Tx means Transmit
	send sendSequenceSpace
	recv recvSequenceSpace
}

// Init use to initialize the socket after allocation in both server or client
func (s *Socket) Init() {
	// TODO:::
	s.mss = OptionDefault_MSS
	s.setState(SocketState_LISTEN)
	// TODO::: set default timeout
	s.readDeadline = *time.NewTimer()
	s.writeDeadline = *time.NewTimer()
	// s.readDeadline.Init()
	// s.writeDeadline.Init()
	checkTimeout(s)
}

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
	checkTimeout(s)
	return
}

/*
********** Local methods **********
 */

func (s *Socket) incomeSegmentOnListenState(segment Packet) (err protocol.Error) {
	if segment.FlagSYN() {
		// err = s.sendSYNandACK()
		s.setState(SocketState_SYN_RECEIVED)

		// TODO::: set ACK timeout timer

		// TODO::: If we return without any error, caller send the socket to the listeners if any exist.
		// Provide a mechanism to let the listener to decide to accept the socket or refuse it??

		// TODO::: attack?? SYN floods, SYN with payload, ...
		return
	}
	// TODO::: attack??
	return
}

func (s *Socket) incomeSegmentOnSynSentState(segment Packet) (err protocol.Error) {
	if segment.FlagSYN() && segment.FlagACK() {
		s.setState(SocketState_ESTABLISHED)
		// TODO::: anything else??
		return
	}
	// TODO::: if we receive syn from sender? attack??
	// TODO::: attack??
	return
}

func (s *Socket) incomeSegmentOnSynReceivedState(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) incomeSegmentOnEstablishedState(segment Packet) (err protocol.Error) {
	var payload = segment.Payload()
	var sn = segment.SequenceNumber()
	var exceptedNext = s.recv.next
	if sn == exceptedNext {
		s.sendACK()

		err = s.recv.buf.Write(payload)

		// TODO::: Due to CongestionControlAlgorithm, if a segment with push flag not send again
		if segment.FlagPSH() {
			err = s.checkPushFlag()
			if err != nil {
				return
			}
			s.recv.sendPushFlagSignal()
			s.Stream.ScheduleProcessingStream()
		}
	} else {
		err = s.validateSequence(segment)
	}
	return
}

func (s *Socket) incomeSegmentOnFinWait1State(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) incomeSegmentOnFinWait2State(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) incomeSegmentOnCloseState(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) incomeSegmentOnCloseWaitState(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) incomeSegmentOnClosingState(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) incomeSegmentOnLastAckState(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) incomeSegmentOnTimeWaitState(segment Packet) (err protocol.Error) {

	return
}

func (s *Socket) handleOptions(opts []byte) (err protocol.Error) {
	for len(opts) > 0 {
		var options = Options(opts)
		switch options.Kind() {
		case OptionKind_EndList:
			return
		case OptionKind_Nop:
			// var opt Option

			// opts = options.NextOption()
		case OptionKind_MSS:
			var optionMSS = optionMSS(options.Payload())
			err = optionMSS.Process(s)
			if err != nil {
				return
			}
			opts = optionMSS.NextOption()
		default:
			// TODO:::
		}
	}
	return
}

// setState change state of socket and send notification on socket state Channel.
func (s *Socket) setState(state SocketState) {
	s.status = state
	select {
	case s.state <- state:
		// state can be delivered by
	default:
		// nothing to do just drop state because channel is block from other
	}

}

// reset the socket and tell peer about reset
func (s *Socket) reset() {
	// TODO:::
}

// sendSYN sending a segment with SYN flag on
func (s *Socket) sendSYN() (err protocol.Error) {
	// TODO:::
	return
}

// sendACK sending ACKs in SYN-RECV and TIME-WAIT states
func (s *Socket) sendACK() (err protocol.Error) {
	s.recv.next++
	// TODO:::
	if DelayedAcknowledgment && s.delayedACK {
		// go to queue
	} else {
		// send segment
	}
	return
}

// sendRST sending RST flag on segment to other side of the socket
func (s *Socket) sendRST() (err protocol.Error) {
	// TODO:::
	return
}

// sendFIN sending FIN flag on segment to other side of the socket
func (s *Socket) sendFIN() (err protocol.Error) {
	// TODO:::
	return
}

// ValidateSequence: validates sequence number of the segment
// Return: TRUE if acceptable, FALSE if not acceptable
func (s *Socket) validateSequence(segment Packet) (err protocol.Error) {
	// TODO::: Change func args if no more data need
	var payload = segment.Payload()
	var sn = segment.SequenceNumber()
	var exceptedNext = s.recv.next
	// TODO::: Due to CongestionControlAlgorithm make a decision to change next
	err = s.recv.buf.WriteIn(payload, (sn - exceptedNext))
	return
}

// ValidateSequence: validates sequence number of the segment
// Return: TRUE if acceptable, FALSE if not acceptable
func (s *Socket) validateSequenceTemp(cur_ts uint32, p Packet, seq uint32, ack_seq uint32, payloadlen int) bool {
	// https://github.com/mtcp-stack/mtcp/blob/master/mtcp/src/tcp_in.c#L108

	return true
}

// needReset check socket state that need RST on ABORT according to RFC793
func (s *Socket) needReset() bool {
	var ss = s.status
	return ss == SocketState_ESTABLISHED || ss == SocketState_CLOSE_WAIT ||
		ss == SocketState_FIN_WAIT_1 || ss == SocketState_FIN_WAIT_2 || ss == SocketState_SYN_RECEIVED
}
