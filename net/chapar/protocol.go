/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

type Connections interface {
	GetConnectionByPath(path []byte) (conn *Connection, err protocol.Error)

	RegisterConnection(conn *Connection) (err protocol.Error)
	DeregisterConnection(conn *Connection) (err protocol.Error)

	protocol.ObjectLifeCycle
}
