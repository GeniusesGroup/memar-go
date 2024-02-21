/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// https://github.com/golang/go/discussions/54245

type ADT_Iteration[ELEMENT any] interface {
	Iterate(startIndex ElementIndex, iterator Iterate[ELEMENT]) (err Error)

	// TODO::: Stop() or return (breaking bool)??
	// Stop() // break the iterate function
}

type Iterate[ELEMENT any] interface {
	// Iterate or traverse
	// In each iteration if err != nil, iteration will be stopped
	Iterate(index ElementIndex, el ELEMENT) (err Error)
}

type Iteration_KV[K, V any] interface {
	Iterate_KV(startIndex int, iterator Iterate_KV[K, V]) (err Error)
}

type Iterate_KV[K, V any] interface {
	// Iterate or traverse
	// In each iteration if err != nil, iteration will be stopped
	Iterate(index int, key K, value V) (err Error)
}
