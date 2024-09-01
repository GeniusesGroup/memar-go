/* For license and copyright information please see the LEGAL file in the code repository */

package ht_p

import (
	adt_p "memar/adt/protocol"
	hash_p "memar/crypto/hash/protocol"
	"memar/protocol"
)

// func New[K Key, V Value](capacity uint64) *HashTable[K, V]

type HashTable[K Key[any], V Value] interface {
	Init(capacity uint64) (err protocol.Error)

	Get(key K) (value V, err protocol.Error) // err return more than exist(bool)
	Put(key K, value V) (err protocol.Error)
	Remove(key K) (err protocol.Error)

	Clear() (err protocol.Error)
	Copy() (new HashTable[K, V], err protocol.Error)
	Iterate(adt_p.Iterate_KV[K, V]) (err protocol.Error)

	adt_p.Capacity
	adt_p.OccupiedLength
}

type AtomicAccessor[K Key[any], V Value] interface {
	Load(key K) (value V, err protocol.Error)
	Store(key K, value V) (err protocol.Error)
	Swap(key K, value V) (old V, err protocol.Error)
	CompareAndSwap(key K, old, new V) (err protocol.Error) // err return more than swapped(bool)
}

type Key[T any] interface {
	protocol.DataType_Equal[T]
	hash_p.Hash64
}

type Value = any
