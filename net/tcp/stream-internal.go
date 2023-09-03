/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	// "syscall"

	"memar/protocol"
)

func (s *Stream) checkStream() (err protocol.Error) {
	if s != nil {
		// err = syscall.EINVAL
	}
	return
}

func (s *Stream) incomeSegmentOnListenState(segment Segment) (err protocol.Error) {
	if segment.FlagSYN() {
		// err = s.sendSYNandACK()
		s.status.Store(StreamStatus_SynReceived)

		// TODO::: set ACK timeout timer

		// TODO::: If we return without any error, caller send the stream to the listeners if any exist.
		// Provide a mechanism to let the listener to decide to accept the stream or refuse it??

		// TODO::: attack?? SYN floods, SYN with payload, ...
		return
	}
	// TODO::: attack??
	return
}

func (s *Stream) incomeSegmentOnSynSentState(segment Segment) (err protocol.Error) {
	if segment.FlagSYN() && segment.FlagACK() {
		s.status.Store(StreamStatus_Established)
		// TODO::: anything else??
		return
	}
	// TODO::: if we receive syn from sender? attack??
	// TODO::: attack??
	return
}

func (s *Stream) incomeSegmentOnSynReceivedState(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) incomeSegmentOnEstablishedState(segment Segment) (err protocol.Error) {
	var payload = segment.Payload()
	var sn = segment.SequenceNumber()
	var exceptedNext = s.recv.next
	if sn == exceptedNext {
		s.sendACK()

		_, err = s.recv.buf.Write(payload)

		// TODO::: Due to CongestionControlAlgorithm, if a segment with push flag not send again
		if segment.FlagPSH() {
			err = s.checkPushFlag()
			if err != nil {
				return
			}

			// TODO:::
			s.recv.sendFlagSignal(flag_PSH)
			s.ScheduleProcessingSocket()
		}
	} else {
		err = s.validateSequence(segment)
	}
	return
}

func (s *Stream) incomeSegmentOnFinWait1State(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) incomeSegmentOnFinWait2State(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) incomeSegmentOnCloseState(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) incomeSegmentOnCloseWaitState(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) incomeSegmentOnClosingState(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) incomeSegmentOnLastAckState(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) incomeSegmentOnTimeWaitState(segment Segment) (err protocol.Error) {

	return
}

func (s *Stream) handleOptions(opts []byte) (err protocol.Error) {
	for len(opts) > 0 {
		var options = Options(opts)
		var optionsKind = options.Kind()
		switch optionsKind {
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

// reset the stream and tell peer about reset
func (s *Stream) reset() {
	// TODO:::
}

func (s *Stream) close() (err protocol.Error) {
	// TODO:::
	err = s.sendFIN()
	if err != nil {
		return
	}
	err = s.Deinit()
	return
}

// sendSYN sending a segment with SYN flag on
func (s *Stream) sendSYN() (err protocol.Error) {
	// TODO:::
	return
}

// sendACK sending ACKs in SYN-RECV and TIME-WAIT states
func (s *Stream) sendACK() (err protocol.Error) {
	s.recv.next++
	// TODO:::
	if CNF_DelayedAcknowledgment && s.timing.de.Enabled() {
		// go to queue
	} else {
		s.sendQuickACK()
	}
	return
}

// sendQuickACK sending ACKs in SYN-RECV and TIME-WAIT states without respect CNF_DelayedAcknowledgment.
func (s *Stream) sendQuickACK() (err protocol.Error) {
	// TODO:::
	return
}

// sendRST sending RST flag on segment to other side of the stream
func (s *Stream) sendRST() (err protocol.Error) {
	// TODO:::
	return
}

// sendFIN sending FIN flag on segment to other side of the stream
func (s *Stream) sendFIN() (err protocol.Error) {
	// TODO:::
	return
}

// ValidateSequence: validates sequence number of the segment
// Return: TRUE if acceptable, FALSE if not acceptable
func (s *Stream) validateSequence(segment Segment) (err protocol.Error) {
	// TODO::: Change func args if no more data need
	var payload = segment.Payload()
	var sn = segment.SequenceNumber()
	var exceptedNext = s.recv.next
	// TODO::: Due to CongestionControlAlgorithm make a decision to change next
	_, err = s.recv.buf.WriteIn(payload, (sn - exceptedNext))
	return
}

// ValidateSequence: validates sequence number of the segment
// Return: TRUE if acceptable, FALSE if not acceptable
func (s *Stream) validateSequenceTemp(cur_ts uint32, p Segment, seq uint32, ack_seq uint32, payloadlen int) bool {
	// https://github.com/mtcp-stack/mtcp/blob/master/mtcp/src/tcp_in.c#L108

	return true
}

// needReset check stream state that need RST on ABORT according to RFC793
func (s *Stream) needReset() bool {
	var ss = s.status.Load()
	return ss == StreamStatus_Established || ss == StreamStatus_CloseWait ||
		ss == StreamStatus_FinWait1 || ss == StreamStatus_FinWait2 || ss == StreamStatus_SynReceived
}

func (s *Stream) sendPayload(b []byte) (n int, err protocol.Error) {

	return
}

// BlockInSelect waits for something to happen, which is one of the following conditions in the function body.
func (s *Stream) blockInSelect() (err protocol.Error) {
	// TODO::: check auto scheduling or block??

loop:
	for {
		select {
		// TODO::: if buffer not full but before get push flag go to full state??
		// I think we must send custom package level flag here when process last segment change buffer state to full.
		case flag := <-s.recv.flag:
			switch flag {
			case flag_FIN:
				// s.readTimer.Stop()
				// err = TODO:::
				break loop
			case flag_RST:
				// s.readTimer.Stop()
				// err = TODO:::
				break loop
			case flag_PSH, flag_URG:
				break loop
			default:
				// TODO::: attack??
				goto loop
			}
		}
	}
	return
}

// func (st *Stream) callService() {
// 	var err = st.Handler().HandleIncomeRequest(st)
// 	if err == nil {
// 		st.connection.StreamSucceed()
// 	} else {
// 		st.connection.StreamFailed()
// 	}
// }
