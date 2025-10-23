/* For license and copyright information please see the LEGAL file in the code repository */

package hashTable_p

import (
	container_p "memar/adt/container/protocol"
	adt_p "memar/adt/protocol"
	capsule_p "memar/computer/capsule/protocol"
	hash_p "memar/crypto/hash/protocol"
	logic_p "memar/math/logic/protocol"
	error_p "memar/process/error/protocol"
	storage_p "memar/storage/memory/protocol"
)

// func New[K Key, V Value](capacity container_p.NumberOfElement) *HashTable[K, V]

type HashTable[K Key[K], V Value] interface {
	capsule_p.LifeCycle
	Init(capacity container_p.NumberOfElement) (err error_p.Error)

	Accessor[K, V]
	AtomicAccessor[K, V]
	Iteration[K, V]

	adt_p.ADT
	container_p.Capacity
	container_p.OccupiedLength
	container_p.Clear

	memory_p.Copy[HashTable[K, V]]
	memory_p.Clone[HashTable[K, V]]
}

type Key[T any] interface {
	logic_p.Equivalence[T]
	hash_p.Hash64
}

type Value = any
