/* For license and copyright information please see the LEGAL file in the code repository */

package array_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/error/protocol"
)

// In computer science, an associative array, map, symbol table, or dictionary is an abstract data type that
// stores a collection of (key, value) pairs, such that each possible key appears at most once in the collection.
// In mathematical terms, an associative array is a function with finite domain.
// https://en.wikipedia.org/wiki/Associative_array
type Associative[KEY comparable, VALUE any] interface {
	Put[KEY, VALUE]
	Lookup[KEY, VALUE]
	Remove[KEY]
	Keys[KEY]
	Values[VALUE]

	Iteration_KV[KEY, VALUE]

	container_p.Capacity
	container_p.OccupiedLength
}

type Iteration_KV[K, V any] interface {
	Iterate_KV(startIndex container_p.ElementIndex, iterator Iterate_KV[K, V]) (err error_p.Error)
}

type Iterate_KV[K, V any] interface {
	// Iterate or traverse
	// In each iteration if err != nil, iteration will be stopped
	Iterate(index container_p.ElementIndex, key K, value V) (err error_p.Error)
}

type Put[KEY, VALUE any] interface {
	// Put or Insert()
	Put(key KEY, value VALUE) (err error_p.Error)
}

type Lookup[KEY comparable, VALUE any] interface {
	Lookup(key KEY) (value VALUE, err error_p.Error)
}

type Remove[KEY comparable] interface {
	Remove(key KEY) (err error_p.Error)
}

type Keys[KEY any] interface {
	// Keys return all keys
	Keys() (keys []KEY, err error_p.Error)
}

type Values[VALUE any] interface {
	// Values return all values
	Values() (values []VALUE, err error_p.Error)
}
