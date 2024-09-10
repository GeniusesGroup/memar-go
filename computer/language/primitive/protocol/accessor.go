/* For license and copyright information please see the LEGAL file in the code repository */

package primitive_p

import (
	error_p "memar/error/protocol"
)

type Accessor[T any] interface {
	Get() T

	// It will check(validate) given value and return proper error for
	Set(new T) (err error_p.Error)
}

type AtomicAccessor[T any] interface {
	Load() T
	Store(new T) (err error_p.Error)
	Swap(new T) (old T, err error_p.Error)
	CompareAndSwap(old, new T) (err error_p.Error) // return more than swapped(bool)
}

type DefaultValue[T any] interface {
	Default() T
	SetDefault() // default value
}

type OptionalValue interface {
	// false means data required and must be exist.
	Optional() bool
}
