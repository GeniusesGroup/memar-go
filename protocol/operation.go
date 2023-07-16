/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Deadline is the interface to show how an operation must implement deadline e.g. network socket, ...
type Deadline interface {
	// SetDeadline sets the read and write deadlines associated with the connection.
	// It is equivalent to calling both SetReadDeadline and SetWriteDeadline.
	SetDeadline(t Time) Error

	Deadline_Read
	Deadline_Write
}
type Deadline_Read interface {
	SetReadDeadline(t Time) Error
}
type Deadline_Write interface {
	SetWriteDeadline(t Time) Error
}

// Timeout is the interface to show how an operation must implement timeout e.g. network socket, ...
type Timeout interface {
	// SetTimeout sets the read and write deadlines associated with the connection.
	// It is equivalent to calling both SetReadTimeout and SetWriteTimeout.
	SetTimeout(d Duration) Error

	Timeout_Read
	Timeout_Write
}
type Timeout_Read interface {
	SetReadTimeout(d Duration) Error
}
type Timeout_Write interface {
	SetWriteTimeout(d Duration) Error
}

type OperationImportance interface {
	// TODO::: need both of below items??!!
	Priority() Priority // Use to queue requests by its priority
	Weight() Weight     // Use to queue requests by its weights in the same priority
}
