/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"syscall"

	"github.com/GeniusesGroup/libgo/protocol"
)

func (s *Socket) reinit() {
	// TODO:::
}

func (t *Socket) deinit() {
	// TODO:::
	t.timing.deinit()
}

func (s *Socket) checkSocket() (err protocol.Error) {
	if s != nil {
		err = syscall.EINVAL
	}
	return
}

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

func (s *Socket) close() (err protocol.Error) {
	err = s.sendFIN()
	s.deinit()
	return
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

func (s *Socket) sendPayload(b []byte) (n int, err error) {

	return
}

// BlockInSelect waits for something to happen, which is one of the following conditions in the function body.
func (s *Socket) blockInSelect() (err protocol.Error) {
loop:
	for {
		select {
		case <-s.readTimer.Signal():
			// TODO:::
			// break
		case flag := <-s.recv.flag:
			switch flag {
			case Flag_FIN:
				s.readTimer.Stop()
				// err = TODO:::
				break loop
			case Flag_RST:
				s.readTimer.Stop()
				// err = TODO:::
				break loop
			case Flag_PSH, Flag_URG:
				break loop
			default:
				// TODO::: attack??
				goto loop
			}
		}
	}
	return
}
