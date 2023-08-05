/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"memar/protocol"
)

// Connection indicate the layer 3 OSI model.
type Connection struct {
	localAddr  Addr
	remoteAddr Addr
}

//memar:impl memar/protocol.ObjectLifeCycle
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

//memar:impl memar/protocol.Network_Framer
func (conn *Connection) FrameType() protocol.Network_FrameType { return protocol.Network_FrameType_GP }

//memar:impl memar/protocol.NetworkAddress
func (conn *Connection) LocalAddr() protocol.Stringer  { return &conn.localAddr }
func (conn *Connection) RemoteAddr() protocol.Stringer { return &conn.remoteAddr }

//memar:impl memar/protocol.Network_FrameWriter
func (conn *Connection) WriteFrame(packet []byte) (n int, err protocol.Error) {
	// TODO:::
	return
}
