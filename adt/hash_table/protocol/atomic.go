/* For license and copyright information please see the LEGAL file in the code repository */

package hashTable_p

import (
	error_p "memar/process/error/protocol"
)

type AtomicAccessor[K Key[K], V Value] interface {
	Load(key K) (value V, err error_p.Error)
	Store(key K, value V) (err error_p.Error)
	Swap(key K, value V) (old V, err error_p.Error)
	CompareAndSwap(key K, old, new V) (err error_p.Error) // err return more than swapped(bool)
}
