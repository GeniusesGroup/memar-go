/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// ADT_Clear is an operation
type ADT_Clear interface {
	// Clear will remove all elements.
	Clear() (err Error)
}

type ADT_Empty interface {
	// an operation for testing whether or not a container is empty.
	Empty() bool
}

// ADT_Reversed is a container level operation.
type ADT_Reversed interface {
	// Reversed is an operation that reverse the container.
	// Copy the CONTAINER if you need the original one.
	Reversed() (err Error)
}

// ADT_Sorted is a container level operation.
type ADT_Sorted interface {
	// Reversed is an operation that sort the container elements.
	// Copy the CONTAINER if you need the original one.
	Sorted() (err Error)
}

// Growth factor
type ADT_Resize interface {
	Resize(ln NumberOfElement) Error
	Resized() bool
	// Resizable returns true if the container(buffer, ...) can be resized, or false if not.
	Resizable() bool
}

/*
----------------------------------------
			CONTAINER LEVEL
----------------------------------------
*/

// ADT_Compare is a container level operation.
type ADT_Compare[CONTAINER any] interface {
	// Compare returns a NumberOfElement comparing two CONTAINER lexicographically.
	Compare(with CONTAINER) NumberOfElement
}

// ADT_Concat is a container level operation.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/concat
type ADT_Concat[CONTAINER any] interface {
	// Concat add given containers in order to end of the CONTAINER.
	// Copy the CONTAINER if you need the original one.
	Concat(con ...CONTAINER) (err Error)
}

/*
----------------------------------------
			Container & ELEMENT LEVEL
----------------------------------------
*/

// ADT_Split is a container level operation.
type ADT_Split_Element[CONTAINER, ELEMENT any] interface {
	// Split is an operation that MOVE the container elements after first given ELEMENT index to new container.
	// Copy the CONTAINER if you need the original one.
	SplitByElement(el ELEMENT) (after CONTAINER, err Error)
}

type ADT_Split_Offset[CONTAINER, ELEMENT any] interface {
	// When `Get` returns limit > len(p), it returns a non-nil error explaining why more bytes were not returned.
	SplitByOffset(offset ElementIndex, limit NumberOfElement) (split CONTAINER, err Error)
}

// Trim

/*
----------------------------------------
			ELEMENT LEVEL
----------------------------------------
*/

type ADT_GetElement[ELEMENT any] interface {
	// When `Get` returns limit > len(p), it returns a non-nil error explaining why more bytes were not returned.
	// GetElement like `GetByte()` provides an efficient interface for byte-at-time processing.
	GetElement(offset ElementIndex) (el ELEMENT, err Error)
}

type ADT_SetElements[ELEMENT any] interface {
	// Set will copy element to the container at given offset.
	// Clients can execute parallel `Set` calls on the same destination if the ranges do not overlap.
	SetElements(offset ElementIndex, el ...ELEMENT) (nAdd NumberOfElement, err Error)
}

// ADT_Push is an element operation
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/push
type ADT_Push[ELEMENT any] interface {
	// Push adds an element to the end of the container
	Push(el ELEMENT) (err Error)
}

// ADT_Push is an element operation
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/pop
type ADT_Pop[ELEMENT any] interface {
	// Pop, which removes the most recently added element and return it.
	Pop() (el ELEMENT, err Error)
}

// ADT_Peek is an element operation
// https://en.wikipedia.org/wiki/Peek_(data_type_operation)
type ADT_Peek[ELEMENT any] interface {
	// Peek or `Top()` or `GetLast()` is an operation on certain abstract data types,
	// specifically sequential collections such as stacks and queues,
	// which returns the value of the top ("front") of the collection without removing the element from the collection.
	Peek() (el ELEMENT, err Error)
}

// ADT_Shift is an element operation
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/shift
type ADT_Shift[ELEMENT any] interface {
	// The shift() method removes the element at the zeroth index and shifts the values at consecutive indexes down,
	// then returns the removed value.
	// The Pop() method has similar behavior to shift(), but applied to the last element in a container.
	Shift() (el ELEMENT, err Error)
}

// ADT_Insert is an element operation
type ADT_Insert[ELEMENT any] interface {
	// Insert will insert the given elements in the offset of the container by move elements after offset to `offset+len(el)`.
	Insert(offset ElementIndex, el ...ELEMENT) (nAdd NumberOfElement, err Error)
}

// ADT_Insert is an element operation
type ADT_Add[ELEMENT any] interface {
	// Add will add the given elements to the container in a location decided by the container logic.
	Add(el ELEMENT) (nAdd NumberOfElement, err Error)
}

// ADT_Append is an element operation
type ADT_Append[ELEMENT any] interface {
	// Append will adds the given elements to the end of the container.
	Append(el ...ELEMENT) (nAdd NumberOfElement, err Error)
}

// ADT_Prepend is an element operation
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/unshift
type ADT_Prepend[ELEMENT any] interface {
	// Prepend will adds the given elements to the beginning of the container.
	Prepend(el ...ELEMENT) (nAdd NumberOfElement, err Error)
}

type ADT_Replace[ELEMENT any] interface {
	//
	Replace(old, new ELEMENT, nl NumberOfElement) (err Error)
}

type ADT_Contain[ELEMENT any] interface {
	// reports whether given element exist in the container.
	// Implementation can easy just by call `Index()` and return true if > 0.
	Contain(el ELEMENT) bool
}
