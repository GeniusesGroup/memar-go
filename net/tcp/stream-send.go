/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
)

// send as Send Sequence Space
type send struct {
	una  uint32 // send unacknowledged
	next uint32
	wnd  uint16 // send window
	up   bool   // send urgent pointer
	wl1  uint32 // segment sequence number used for last window update
	wl2  uint32 // segment acknowledgment number used for last window update
	iss  uint32 // initial send sequence number
	// buf    []byte Don't need it, because we don't need to copy buffer between kernel and user-space
}

//memar:impl memar/protocol.ObjectLifeCycle
func (s *send) Init(timeout protocol.Duration) (err protocol.Error) {
	// TODO:::
	return
}
func (s *send) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (s *send) Deinit() (err protocol.Error) {
	// TODO:::
	return
}
