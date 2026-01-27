/* For license and copyright information please see the LEGAL file in the code repository */

package hashTable_p

import (
	error_p "memar/process/error/protocol"
)

type Iteration[K Key[K], V Value] interface {
	Iteration(iterator Iterator[K, V]) (err error_p.Error)
}

type Iterator[K Key[K], V Value] interface {
	// Iterate or traverse
	// In each iteration if err != nil, iteration will be stopped
	Iterate(key K, value V) (err error_p.Error)
}
