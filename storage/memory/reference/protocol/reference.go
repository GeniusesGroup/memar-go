/* For license and copyright information please see the LEGAL file in the code repository */

package reference_p

import (
	capsule_p "memar/computer/capsule/protocol"
	error_p "memar/process/error/protocol"
)

// A reference is an abstract data type and may be implemented in many ways 
// In computer programming, a reference is a value that enables a program to indirectly access a particular datum,
// such as a variable's value or a record, in the computer's memory or in some other storage device.
// The reference is said to refer to the datum, and accessing the datum is called dereferencing the reference. 
// A reference is distinct from the datum itself.
//
// https://en.wikipedia.org/wiki/Reference_(computer_science)
type Reference[T any] interface {
	// Returns a direct access of the type itself
	// Devs CAN made a `Copy` from return value.
	// It is very similar to capsule_p.Accessor but with some differences
	Dereference() (t T, err error_p.Error)

	// TODO::: we need out of scope mechanism like RUST, how to implement this in the library instead of language syntax??
	capsule_p.LifeCycle
}
