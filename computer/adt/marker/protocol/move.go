/* For license and copyright information please see the LEGAL file in the code repository */

package marker_p

import (
	error_p "memar/process/error/protocol"
)

// Move is `Clone`>>`Delete` operation
type Move[T any] interface {
	// Performs move-assignment from source
	MoveFrom(source T) (err error_p.Error)
	// Performs move-assignment to destination
	MoveTo(destination T) (err error_p.Error)
}
