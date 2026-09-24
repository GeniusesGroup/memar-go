/* For license and copyright information please see the LEGAL file in the code repository */

package reference

// Immutability does not imply that the object as stored in the computer's memory is un-writeable.
// Rather, immutability is a compile-time construct that indicates what a programmer can do
// through the normal interface of the object, not necessarily what they can absolutely do
type Immutable[T any] struct{}

func (self *Immutable[T]) Mutable() bool { return false }
