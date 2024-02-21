/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// func New[K HashTable_Key, V HashTable_Value](capacity uint64) *HashTable[K, V]

type HashTable[K HashTable_Key[any], V HashTable_Value] interface {
	Init(capacity uint64) (err Error)

	Get(key K) (value V, err Error) // err return more than exist(bool)
	Put(key K, value V) (err Error)
	Remove(key K) (err Error)

	Clear() (err Error)
	Copy() (new HashTable[K, V], err Error)
	Iterate(Iterate_KV[K, V]) (err Error)

	Capacity
	OccupiedLength
}

type HashTable_AtomicAccessor[K HashTable_Key[any], V HashTable_Value] interface {
	Load(key K) (value V, err Error)
	Store(key K, value V) (err Error)
	Swap(key K, value V) (old V, err Error)
	CompareAndSwap(key K, old, new V) (err Error) // err return more than swapped(bool)
}

type HashTable_Key[T any] interface {
	DataType_Equal[T]
	Hash64
}

type HashTable_Value = any
