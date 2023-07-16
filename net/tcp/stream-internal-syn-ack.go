/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/protocol"
)

type synAck struct {
	maxSynAckRetry uint8
	retrySynAck    uint8
}

func (s *synAck) Init() (err protocol.Error) {
	s.maxSynAckRetry = CNF_SynAck_Retries
	return
}
func (s *synAck) Reinit() (err protocol.Error) {
	s.maxSynAckRetry = CNF_SynAck_Retries
	return
}
func (s *synAck) Deinit() (err protocol.Error) {
	return
}

func (s *synAck) ReachMaxSynACK() (max bool) {
	if s.retrySynAck > s.maxSynAckRetry {
		return true
	}
	return false
}

func (s *synAck) ReceiveNewSynAck() {
	s.retrySynAck++
}
