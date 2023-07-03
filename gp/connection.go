/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/net"
	"libgo/protocol"
)

// Connection indicate the layer 3 OSI model.
type Connection struct {
	localAddr  Addr
	remoteAddr Addr

	net.STATUS
	net.Metric
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (conn *Connection) Init() (err protocol.Error)   { return }
func (conn *Connection) Reinit() (err protocol.Error) { return }
func (conn *Connection) Deinit() (err protocol.Error) {
	// first closing open listener for income frame and refuse all new frame,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
	return
}

//libgo:impl libgo/protocol.NetworkAddress
func (conn *Connection) LocalAddr() protocol.Stringer  { return &conn.localAddr }
func (conn *Connection) RemoteAddr() protocol.Stringer { return &conn.remoteAddr }

//libgo:impl libgo/protocol.Network_FrameWriter
func (conn *Connection) WriteFrame(packet []byte) (n int, err protocol.Error) {
	// TODO:::
	return
}

// EstablishNewConnectionByDomainID make new connection by peer domain ID and initialize it.
func EstablishNewConnectionByDomainID(domainID [32]byte) (conn *Connection, err protocol.Error) {
	// TODO::: Get closest domain GP add
	var domainGPAddr = Addr{}
	conn, err = EstablishNewConnectionByPeerAdd(domainGPAddr)
	if err != nil {
		return
	}
	// conn.userID = domainID
	// conn.userType = protocol.UserType_App
	return
}

// EstablishNewConnectionByPeerAdd make new connection by peer GP and initialize it.
func EstablishNewConnectionByPeerAdd(remoteAddr Addr) (conn *Connection, err protocol.Error) {
	// If conn not exist means guest connection.
	if conn == nil {
		conn, err = MakeNewGuestConnection()
		if err == nil {
			conn.remoteAddr = remoteAddr
		}
	}
	return
}

// MakeNewGuestConnection make new connection and register on given stream due to it is first attempt connect to server.
func MakeNewGuestConnection() (conn *Connection, err protocol.Error) {
	// if Server.Manifest.TechnicalInfo.GuestMaxConnections == 0 {
	// 	return nil, ErrGuestConnectionNotAllow
	// } else if Server.Manifest.TechnicalInfo.GuestMaxConnections > 0 && Server.Connections.GuestConnectionCount > Server.Manifest.TechnicalInfo.GuestMaxConnections {
	// 	return nil, ErrGuestConnectionMaxReached
	// }

	conn = &Connection{
		// ID:       uuid.Random32Byte(),
		// status: protocol.NetworkStatus_New,
		// userType: protocol.UserType_Unset,
	}
	return
}
