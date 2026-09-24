/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	capsule_p "memar/computer/capsule/protocol"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

type Connections interface {
	GetConnectionByPeerAddr(addr Addr) (conn *Connection, err error_p.Error)
	// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
	GetConnectionBySessionID(sID net_p.SessionID) (conn *Connection, err error_p.Error)

	// state=unregistered -> 'register' -> state=registered -> 'deregister' -> state=unregistered.
	RegisterConnection(conn *Connection) (err error_p.Error)
	DeregisterConnection(conn *Connection) (err error_p.Error)
	RevokeConnection(conn *Connection) (err error_p.Error)

	// Deinit the listener when the application closes or force to closes by not recovered panic.
	capsule_p.LifeCycle
}
