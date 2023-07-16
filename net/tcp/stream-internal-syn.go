/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/protocol"
)

var maxSynBacklog int

func init() {
	// The default value of 256 is
	// increased to 1024 when the memory present in the system is
	// adequate or greater (>= 128 MB), and reduced to 128 for
	// those systems with very low memory (<= 32 MB).
}

type syn struct {
	maxSynRetry uint8
	retrySyn    uint8
}

func (s *syn) Init() (err protocol.Error) {
	s.maxSynRetry = CNF_SynRetries
	return
}
func (s *syn) Reinit() (err protocol.Error) {
	s.maxSynRetry = CNF_SynRetries
	return
}
func (s *syn) Deinit() (err protocol.Error) {
	return
}

func (s *syn) ReachMaxSyn() (max bool) {
	if s.retrySyn > s.maxSynRetry {
		return true
	}
	return false
}

func (s *syn) SendNewSyn() {
	s.retrySyn++
}
