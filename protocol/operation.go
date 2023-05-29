/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Deadline is the interface to show how an operation must implement deadline e.g. network stream, ...
type Deadline interface {
	// SetDeadline sets the read and write deadlines associated with the connection.
	// It is equivalent to calling both SetReadDeadline and SetWriteDeadline.
	SetDeadline(t Time) Error
	SetReadDeadline(t Time) Error
	SetWriteDeadline(t Time) Error
}

// Timeout is the interface to show how an operation must implement timeout e.g. network stream, ...
type Timeout interface {
	// SetTimeout sets the read and write deadlines associated with the connection.
	// It is equivalent to calling both SetReadTimeout and SetWriteTimeout.
	SetTimeout(d Duration) Error
	SetReadTimeout(d Duration) Error
	SetWriteTimeout(d Duration) Error
}

type OperationImportance interface {
	// TODO::: need both of below items??!!
	Priority() Priority // Use to queue requests by its priority
	Weight() Weight     // Use to queue requests by its weights in the same priority
}
