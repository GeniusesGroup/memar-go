/* For license and copyright information please see the LEGAL file in the code repository */

package reference

type Mutable[T any] struct{}

func (self *Mutable[T]) Mutable() bool { return true }
