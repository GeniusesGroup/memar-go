/* For license and copyright information please see the LEGAL file in the code repository */

package sync_p

import (
	error_p "memar/process/error/protocol"
)

type AtomicAccessor[T any] interface {
	Load() T
	Store(new T) (err error_p.Error)
	Swap(new T) (old T, err error_p.Error)
	CompareAndSwap(old, new T) (err error_p.Error) // return more than swapped(bool)
}
