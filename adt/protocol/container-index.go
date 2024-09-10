/* For license and copyright information please see the LEGAL file in the code repository */

package adt_p

import (
	error_p "memar/error/protocol"
)

// ElementIndex is related to `NumberOfElement`
// ElementIndex can refer to any location of memory blocks in byte or 8bit number.
type ElementIndex int

type LastElementIndex interface {
	// LastElementIndex return the location of last element in the container.
	LastElementIndex() ElementIndex
}

type Index[ELEMENT Element] interface {
	// Index return the location of given element in the container.
	Index(el ELEMENT) (ei ElementIndex, err error_p.Error)
}

type LastIndex[ELEMENT Element] interface {
	// LastIndex return the location of given element in the container from end of container.
	LastIndex(el ELEMENT) (ei ElementIndex, err error_p.Error)
}
