/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	// "syscall"

	adt_p "memar/adt/protocol"
	error_p "memar/error/protocol"
)

func (s *Stream) checkStream() (err error_p.Error) {
	if s != nil {
		// err = syscall.EINVAL
	}
	return
}

func (s *Stream) checkSegmentFlags(segment Segment) (err error_p.Error) {
	// TODO:::
	return
}

func (s *Stream) incomeSegmentOnListenState(segment Segment) (err error_p.Error) {
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

func (s *Stream) incomeSegmentOnSynSentState(segment Segment) (err error_p.Error) {
	if segment.FlagSYN() && segment.FlagACK() {
		s.status.Store(StreamStatus_Established)
		// TODO::: anything else??
		return
	}
	// TODO::: if we receive syn from sender? attack??
	// TODO::: attack??
	return
}

func (s *Stream) incomeSegmentOnSynReceivedState(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) incomeSegmentOnEstablishedState(segment Segment) (err error_p.Error) {
	var payload = segment.Payload()
	var sn = segment.SequenceNumber()
	var exceptedNext = s.recv.next
	if sn == exceptedNext {
		s.sendACK()

		_, err = s.sk.ReceiveBuffer().Append(payload...)

		// TODO::: Due to CongestionControlAlgorithm, if a segment with push flag not send again
		if segment.FlagPSH() {
			err = s.checkPushFlag()
			if err != nil {
				return
			}

			// TODO:::
			// s.recv.sendFlagSignal(flag_PSH)
			s.sk.ApplicationLayer().ScheduleProcessingSocket()
		}
	} else {
		err = s.validateSequence(segment)
	}
	return
}

func (s *Stream) incomeSegmentOnFinWait1State(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) incomeSegmentOnFinWait2State(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) incomeSegmentOnCloseState(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) incomeSegmentOnCloseWaitState(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) incomeSegmentOnClosingState(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) incomeSegmentOnLastAckState(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) incomeSegmentOnTimeWaitState(segment Segment) (err error_p.Error) {

	return
}

func (s *Stream) handleOptions(opts []byte) (err error_p.Error) {
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

func (s *Stream) close() (err error_p.Error) {
	// TODO:::
	err = s.sendFIN()
	if err != nil {
		return
	}
	err = s.Deinit()
	return
}

// sendSYN sending a segment with SYN flag on
func (s *Stream) sendSYN() (err error_p.Error) {
	// TODO:::
	return
}

// sendACK sending ACKs in SYN-RECV and TIME-WAIT states
func (s *Stream) sendACK() (err error_p.Error) {
	s.recv.next++
	if CNF_DelayedAcknowledgment && s.timing.de.Enabled() {
		// TODO::: go to queue
	} else {
		s.sendQuickACK()
	}
	return
}

// sendQuickACK sending ACKs in SYN-RECV and TIME-WAIT states without respect CNF_DelayedAcknowledgment.
func (s *Stream) sendQuickACK() (err error_p.Error) {
	// TODO:::
	return
}

// sendRST sending RST flag on segment to other side of the stream
func (s *Stream) sendRST() (err error_p.Error) {
	// TODO:::
	return
}

// sendFIN sending FIN flag on segment to other side of the stream
func (s *Stream) sendFIN() (err error_p.Error) {
	// TODO:::
	return
}

// ValidateSequence: validates sequence number of the segment
// Return: TRUE if acceptable, FALSE if not acceptable
func (s *Stream) validateSequence(segment Segment) (err error_p.Error) {
	// TODO::: Change func args if no more data need
	var payload = segment.Payload()
	var sn = segment.SequenceNumber()
	var exceptedNext = s.recv.next
	// TODO::: Due to CongestionControlAlgorithm make a decision to change next
	_, err = s.sk.ReceiveBuffer().Insert(adt_p.ElementIndex(sn-exceptedNext), payload...)
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

// func (st *Stream) callService() {
// 	var err = st.Handler().HandleIncomeRequest(st)
// 	if err == nil {
// 		st.connection.StreamSucceed()
// 	} else {
// 		st.connection.StreamFailed()
// 	}
// }
