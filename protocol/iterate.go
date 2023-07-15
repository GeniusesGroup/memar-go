/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Iterate[T any] interface {
	// Iterate or traverse
	Iterate(t T) (breaking bool)
}

type Iterate_KV[T any] interface {
	// Iterate or traverse
	Iterate(key, value T) (breaking bool)
}
