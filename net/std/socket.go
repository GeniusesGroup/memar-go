/* For license and copyright information please see the LEGAL file in the code repository */

package std

import (
	"memar/net"
	"memar/protocol"
	// "memar/uuid/16byte"
)

type Socket struct {
	net.Socket
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (sk *Socket) Init(timeout protocol.Duration) (err protocol.Error) {
	return
}
func (sk *Socket) Reinit() (err protocol.Error) { return }
func (sk *Socket) Deinit() (err protocol.Error) { return }
