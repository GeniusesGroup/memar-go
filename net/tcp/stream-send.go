/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	error_p "memar/error/protocol"
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

// memar/computer/language/object/protocol.LifeCycle
func (s *send) Init() (err error_p.Error) {
	// TODO:::
	return
}
func (s *send) Reinit() (err error_p.Error) {
	// TODO:::
	return
}
func (s *send) Deinit() (err error_p.Error) {
	// TODO:::
	return
}
