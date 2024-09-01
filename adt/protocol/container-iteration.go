/* For license and copyright information please see the LEGAL file in the code repository */

package adt_p

import (
	"memar/protocol"
)

// https://github.com/golang/go/discussions/54245

type Iteration[ELEMENT Element] interface {
	Iterate(startIndex ElementIndex, iterator Iterate[ELEMENT]) (err protocol.Error)

	// TODO::: Stop() or return (breaking bool)??
	// Stop() // break the iterate function
}

type Iterate[ELEMENT Element] interface {
	// Iterate or traverse
	// In each iteration if err != nil, iteration will be stopped
	Iterate(index ElementIndex, el ELEMENT) (err protocol.Error)
}
