/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"../protocol"
	"../timer"
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

	socketTimer timer.Timer
	readTimer   timer.Timer // read deadline timer
	writeTimer  timer.Timer // write deadline timer

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
	s.readTimer.Init()
	s.writeTimer.Init()
	checkTimeout(s)

	s.recv.init()
	s.send.init()
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
		return
	}
	s.readTimer.Reset(d)
	return
}
func (s *Socket) SetWriteTimeout(d protocol.Duration) (err protocol.Error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	if d < 0 {
		// no timeout
		return
	}
	s.writeTimer.Reset(d)
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
