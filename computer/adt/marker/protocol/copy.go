/* For license and copyright information please see the LEGAL file in the code repository */

package marker_p

import (
	error_p "memar/process/error/protocol"
)

// Copy is implicit, inexpensive, and cannot be re-implemented (memcpy).
// The Copy trait represents values that can be safely duplicated via memcpy:
// things like reassignments and passing an argument by-value to a function are always memcpys, and so for Copy types,
// the compiler understands that it doesn't need to consider those a move.
// every Copy type is also required to be Clone
type Copy[T any] interface {
	// Returns a copy of the itself
	Copy() (c T, err error_p.Error)
	// Performs copy-assignment from source
	CopyFrom(source T) (err error_p.Error)
	// Performs copy-assignment to destination
	CopyTo(destination T) (err error_p.Error)
}
