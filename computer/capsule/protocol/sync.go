/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	error_p "memar/process/error/protocol"
)

type Locker interface {
	// Be aware that the requested `thread` which call `Lock()` CAN `yield`!
	Lock() (err error_p.Error)
	Unlock() (err error_p.Error)
}

type AtomicAccessor[T any] interface {
	Load() T
	Store(new T) (err error_p.Error)
	Swap(new T) (old T, err error_p.Error)
	CompareAndSwap(old, new T) (err error_p.Error) // return more than swapped(bool)
}
