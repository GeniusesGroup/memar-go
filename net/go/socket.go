/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/net"
	"memar/protocol"
)

type Socket struct{
	net.Socket
}

//memar:impl memar/protocol.ObjectLifeCycle
func (sk *Socket) Init() (err protocol.Error) {
	return
}
func (sk *Socket) Reinit() (err protocol.Error) {
	return
}
func (sk *Socket) Deinit() (err protocol.Error) {
	return
}
