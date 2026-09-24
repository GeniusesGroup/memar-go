/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	capsule_p "memar/computer/capsule/protocol"
	error_p "memar/process/error/protocol"
)

type Connections interface {
	GetConnectionByPath(path []byte) (conn *Connection, err error_p.Error)

	RegisterConnection(conn *Connection) (err error_p.Error)
	DeregisterConnection(conn *Connection) (err error_p.Error)

	capsule_p.LifeCycle
}
