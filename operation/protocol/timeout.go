/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

import (
	error_p "memar/error/protocol"
	"memar/time/duration"
)

// Timeout is the interface to show how an operation must implement timeout e.g. network socket, ...
type Timeout interface {
	// SetTimeout sets the read and write deadlines associated with the connection.
	// It is equivalent to calling both SetReadTimeout and SetWriteTimeout.
	SetTimeout(d duration.NanoSecond) error_p.Error

	Timeout_Read
	Timeout_Write
}
type Timeout_Read interface {
	SetReadTimeout(d duration.NanoSecond) error_p.Error
}
type Timeout_Write interface {
	SetWriteTimeout(d duration.NanoSecond) error_p.Error
}
