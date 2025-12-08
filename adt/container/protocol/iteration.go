/* For license and copyright information please see the LEGAL file in the code repository */

package container_p

import (
	error_p "memar/process/error/protocol"
)

// https://github.com/golang/go/discussions/54245

// Iteration or Iterable
// 
// Other protocols:::
// https://www.typescriptlang.org/docs/handbook/symbols.html#symboliterator
type Iteration[ELEMENT Element] interface {
	// It can return `Iterate` errors or internal `Iteration` errors.
	Iteration(startIndex ElementIndex, iterator Iterate[ELEMENT]) (err error_p.Error)

	// TODO::: Stop() or return (breaking bool)??
	// Stop() // break the iterate function
}

type Iterate[ELEMENT Element] interface {
	// Iterate or traverse
	// In each iteration if err != nil, iteration will be stopped
	Iterate(index ElementIndex, el ELEMENT) (err error_p.Error)
}
