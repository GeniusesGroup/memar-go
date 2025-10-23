/* For license and copyright information please see the LEGAL file in the code repository */

package hashTable_p

import (
	error_p "memar/process/error/protocol"
)

type Accessor[K Key[K], V Value] interface {
	Get(key K) (value V, err error_p.Error) // err return more than exist(bool)
	Put(key K, value V) (err error_p.Error)
	Remove(key K) (err error_p.Error)
}
