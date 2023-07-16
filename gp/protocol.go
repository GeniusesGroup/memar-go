/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/protocol"
)

type Connections interface {
	GetConnectionByPeerAddr(addr Addr) (conn *Connection, err protocol.Error)
	// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
	GetConnectionByUserIDDelegateUserID(userID, delegateUserID [16]byte) (conn *Connection, err protocol.Error)
	GetConnectionsByUserID(userID protocol.UserUUID) (conns []*Connection, err protocol.Error)
	GetConnectionByDomain(domain string) (conn *Connection, err protocol.Error)

	// state=unregistered -> 'register' -> state=registered -> 'deregister' -> state=unregistered.
	RegisterConnection(conn *Connection) (err protocol.Error)
	DeregisterConnection(conn *Connection) (err protocol.Error)
	RevokeConnection(conn *Connection) (err protocol.Error)

	// Deinit the listener when the application closes or force to closes by not recovered panic.
	protocol.ObjectLifeCycle
}
