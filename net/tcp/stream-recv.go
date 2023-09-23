/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
)

// recv is receive sequence space
type recv struct {
	next uint32 // receive next
	wnd  uint16 // receive window
	up   bool   // receive urgent pointer
	irs  uint32 // initial receive sequence number
	// TODO::: not in order segments
}

//memar:impl memar/protocol.ObjectLifeCycle
func (r *recv) Init() (err protocol.Error) {
	// TODO:::
	return
}
func (r *recv) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (r *recv) Deinit() (err protocol.Error) {
	// TODO:::
	return
}
