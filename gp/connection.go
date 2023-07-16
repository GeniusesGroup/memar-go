/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/protocol"
)

// Connection indicate the layer 3 OSI model.
type Connection struct {
	localAddr  Addr
	remoteAddr Addr
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (conn *Connection) Init(localAddr, remoteAddr Addr) (err protocol.Error) {
	conn.localAddr = localAddr
	conn.remoteAddr = remoteAddr
	return
}
func (conn *Connection) Reinit() (err protocol.Error) {
	// TODO::: reset metrics
	return
}
func (conn *Connection) Deinit() (err protocol.Error) { return }

//libgo:impl libgo/protocol.Network_Framer
func (conn *Connection) FrameID() protocol.Network_FrameID { return protocol.Network_FrameID_GP }

//libgo:impl libgo/protocol.NetworkAddress
func (conn *Connection) LocalAddr() protocol.Stringer  { return &conn.localAddr }
func (conn *Connection) RemoteAddr() protocol.Stringer { return &conn.remoteAddr }

//libgo:impl libgo/protocol.Network_FrameWriter
func (conn *Connection) WriteFrame(packet []byte) (n int, err protocol.Error) {
	// TODO:::
	return
}
