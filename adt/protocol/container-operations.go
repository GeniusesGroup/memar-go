/* For license and copyright information please see the LEGAL file in the code repository */

package adt_p

import (
	"memar/protocol"
)

// Clear is an operation
type Clear interface {
	// Clear will remove all elements.
	Clear() (err protocol.Error)
}

// Reversed is a container level operation.
type Reversed interface {
	// Reversed is an operation that reverse the container.
	// Copy the CONTAINER if you need the original one.
	Reversed() (err protocol.Error)
}

// Sorted is a container level operation.
type Sorted interface {
	// Reversed is an operation that sort the container elements.
	// Copy the CONTAINER if you need the original one.
	Sorted() (err protocol.Error)
}

// Growth factor
type Resize interface {
	Resize(ln NumberOfElement) protocol.Error
	Resized() bool
	// Resizable returns true if the container(buffer, ...) can be resized, or false if not.
	Resizable() bool
}

/*
----------------------------------------
			CONTAINER LEVEL
----------------------------------------
*/

// Compare is a container level operation.
type Compare[CONTAINER any] interface {
	// Compare returns a NumberOfElement comparing two CONTAINER lexicographically.
	Compare(with CONTAINER) NumberOfElement
}

// Concat is a container level operation.
// https://en.wikipedia.org/wiki/Concatenation
// https://en.wikipedia.org/wiki/Alternation_(formal_language_theory)
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/concat
type Concat[CONTAINER any] interface {
	// Concat is like `Append()` and `Prepend()` add given containers in order to end of the CONTAINER,
	// and return new CONTAINER.
	Concat(con ...CONTAINER) (new CONTAINER, err protocol.Error)
}

type Replace_Elements[CONTAINER any] interface {
	// Replace_Elements is like `Replace()` but in container level.
	// Copy the CONTAINER if you need the original one.
	Replace_Elements(old, new CONTAINER, nl NumberOfElement) (err protocol.Error)
}

/*
----------------------------------------
			Container & ELEMENT LEVEL
----------------------------------------
*/

// Split_Element is a container level operation.
type Split_Element[CONTAINER, ELEMENT Element] interface {
	// Split is an operation that MOVE the container elements after first given ELEMENT index to new container.
	// Copy the CONTAINER if you need the original one.
	SplitByElement(el ELEMENT) (after CONTAINER, err protocol.Error)
}

type Split_Offset[CONTAINER, ELEMENT Element] interface {
	// When `Get` returns limit > len(p), it returns a non-nil error explaining why more bytes were not returned.
	SplitByOffset(offset ElementIndex, limit NumberOfElement) (split CONTAINER, err protocol.Error)
}

// Trim

/*
----------------------------------------
			ELEMENT LEVEL
----------------------------------------
*/

type GetElement[ELEMENT Element] interface {
	// When `Get` returns limit > len(p), it returns a non-nil error explaining why more bytes were not returned.
	// GetElement like `GetByte()` provides an efficient interface for byte-at-time processing.
	GetElement(offset ElementIndex) (el ELEMENT, err protocol.Error)
}

type SetElements[ELEMENT Element] interface {
	// Set will copy element to the container at given offset.
	// Clients can execute parallel `Set` calls on the same destination if the ranges do not overlap.
	SetElements(offset ElementIndex, el ...ELEMENT) (nAdd NumberOfElement, err protocol.Error)
}

// Push is an element operation
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/push
type Push[ELEMENT Element] interface {
	// Push adds an element to the end of the container
	Push(el ELEMENT) (err protocol.Error)
}

// Pop is an element operation
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/pop
type Pop[ELEMENT Element] interface {
	// Pop, which removes the most recently added element and return it.
	Pop() (el ELEMENT, err protocol.Error)
}

// Peek is an element operation
// https://en.wikipedia.org/wiki/Peek_(data_type_operation)
type Peek[ELEMENT Element] interface {
	// Peek or `Top()` or `GetLast()` is an operation on certain abstract data types,
	// specifically sequential collections such as stacks and queues,
	// which returns the value of the top ("front") of the collection without removing the element from the collection.
	Peek() (el ELEMENT, err protocol.Error)
}

// Shift is an element operation
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/shift
type Shift[ELEMENT Element] interface {
	// The shift() method removes the element at the zeroth index and shifts the values at consecutive indexes down,
	// then returns the removed value.
	// The Pop() method has similar behavior to shift(), but applied to the last element in a container.
	Shift() (el ELEMENT, err protocol.Error)
}

// Insert is an element operation
type Insert[ELEMENT Element] interface {
	// Insert will insert the given elements in the offset of the container by move elements after offset to `offset+len(el)`.
	Insert(offset ElementIndex, el ...ELEMENT) (nAdd NumberOfElement, err protocol.Error)
}

// Add is an element operation
type Add[ELEMENT Element] interface {
	// Add will add the given elements to the container in a location decided by the container logic.
	Add(el ELEMENT) (nAdd NumberOfElement, err protocol.Error)
}

// Append is an element operation
type Append[ELEMENT Element] interface {
	// Append will adds the given elements to the end of the container.
	Append(el ...ELEMENT) (nAdd NumberOfElement, err protocol.Error)
}

// Prepend is an element operation
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/unshift
type Prepend[ELEMENT Element] interface {
	// Prepend will adds the given elements to the beginning of the container.
	Prepend(el ...ELEMENT) (nAdd NumberOfElement, err protocol.Error)
}

type Replace[ELEMENT Element] interface {
	//
	Replace(old, new ELEMENT, nl NumberOfElement) (err protocol.Error)
}

type Contain[ELEMENT Element] interface {
	// reports whether given element exist in the container.
	// Implementation can easy just by call `Index()` and return true if > 0.
	Contain(el ELEMENT) bool
}
