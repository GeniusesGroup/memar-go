/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	error_p "memar/process/error/protocol"
	net_p "memar/net/protocol"
)

// Connection indicate the layer 3 OSI model.
type Connection struct {
	localAddr  Addr
	remoteAddr Addr
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (conn *Connection) Init(localAddr, remoteAddr Addr) (err error_p.Error) {
	conn.localAddr = localAddr
	conn.remoteAddr = remoteAddr
	return
}
func (conn *Connection) Reinit() (err error_p.Error) {
	// TODO::: reset metrics
	return
}
func (conn *Connection) Deinit() (err error_p.Error) { return }

//memar:impl memar/net/protocol.Field_FrameType
func (conn *Connection) FrameType() net_p.FrameType { return net_p.FrameType_GP }

//memar:impl memar/net/protocol.Field_NetworkAddresses
func (conn *Connection) LocalAddr() net_p.NetworkAddress  { return &conn.localAddr }
func (conn *Connection) RemoteAddr() net_p.NetworkAddress { return &conn.remoteAddr }

//memar:impl memar/protocol.Network_FrameWriter
func (conn *Connection) WriteFrame(packet []byte) (n int, err error_p.Error) {
	// TODO:::
	return
}
