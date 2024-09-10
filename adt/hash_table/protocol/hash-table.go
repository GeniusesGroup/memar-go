/* For license and copyright information please see the LEGAL file in the code repository */

package ht_p

import (
	adt_p "memar/adt/protocol"
	primitive_p "memar/computer/language/primitive/protocol"
	hash_p "memar/crypto/hash/protocol"
	error_p "memar/error/protocol"
)

// func New[K Key, V Value](capacity uint64) *HashTable[K, V]

type HashTable[K Key[K], V Value] interface {
	Init(capacity uint64) (err error_p.Error)

	Get(key K) (value V, err error_p.Error) // err return more than exist(bool)
	Put(key K, value V) (err error_p.Error)
	Remove(key K) (err error_p.Error)

	Clear() (err error_p.Error)
	Copy() (new HashTable[K, V], err error_p.Error)
	Iterate(adt_p.Iterate_KV[K, V]) (err error_p.Error)

	adt_p.Capacity
	adt_p.OccupiedLength
}

type AtomicAccessor[K Key[any], V Value] interface {
	Load(key K) (value V, err error_p.Error)
	Store(key K, value V) (err error_p.Error)
	Swap(key K, value V) (old V, err error_p.Error)
	CompareAndSwap(key K, old, new V) (err error_p.Error) // err return more than swapped(bool)
}

type Key[T any] interface {
	primitive_p.Equivalence[T]
	hash_p.Hash64
}

type Value = any
