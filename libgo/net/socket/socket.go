/* For license and copyright information please see the LEGAL file in the code repository */

package std

import (
	error_p "memar/process/error/protocol"
	"memar/net"
	"memar/time/duration"
	// "memar/uuid/16byte"
)

type Socket struct {
	net.Socket
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (sk *Socket) Init(timeout duration.NanoSecond) (err error_p.Error) {
	return
}
func (sk *Socket) Reinit() (err error_p.Error) { return }
func (sk *Socket) Deinit() (err error_p.Error) { return }
