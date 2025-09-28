/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

import (
	error_p "memar/error/protocol"
	time_p "memar/time/protocol"
)

// Deadline is the interface to show how an operation must implement deadline e.g. network socket, ...
type Deadline interface {
	// SetDeadline sets the read and write deadlines associated with the connection.
	// It is equivalent to calling both SetReadDeadline and SetWriteDeadline.
	SetDeadline(t time_p.Time) error_p.Error

	Deadline_Read
	Deadline_Write
}
type Deadline_Read interface {
	SetReadDeadline(t time_p.Time) error_p.Error
}
type Deadline_Write interface {
	SetWriteDeadline(t time_p.Time) error_p.Error
}
