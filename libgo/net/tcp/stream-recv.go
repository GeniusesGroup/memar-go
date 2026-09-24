/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	container_p "memar/computer/adt/container/protocol"
	error_p "memar/process/error/protocol"
)

// recv is receive sequence space
type recv struct {
	next         container_p.ElementIndex // receive next
	wnd          uint16                   // receive window
	up           bool                     // receive urgent pointer
	irs          uint32                   // initial receive sequence number
	recvPushFlag bool
	recvUrgFlag  bool
	// TODO::: not in order segments
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *recv) Init() (err error_p.Error) {
	// TODO:::
	return
}
func (r *recv) Reinit() (err error_p.Error) {
	// TODO:::
	return
}
func (r *recv) Deinit() (err error_p.Error) {
	// TODO:::
	return
}
