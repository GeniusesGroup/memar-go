/* For license and copyright information please see the LEGAL file in the code repository */

package marker_p

import (
	error_p "memar/process/error/protocol"
)

// Clone is explicit, may be expensive, and may be re-implement arbitrarily.
// Clone is designed for arbitrary duplications:
// a Clone implementation for a type T can do arbitrarily complicated operations required to create a new T.
// It is a normal trait (other than being in the prelude), and so requires being used like a normal trait, with method calls, etc.
type Clone[T any] interface {
	// Returns a clone of the itself
	Clone() (c T, err error_p.Error)
	// Performs clone-assignment from source
	CloneFrom(source T) (err error_p.Error)
	// Performs clone-assignment to destination
	CloneTo(destination T) (err error_p.Error)
}
