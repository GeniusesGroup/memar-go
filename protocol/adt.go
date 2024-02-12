/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// https://en.wikipedia.org/wiki/Abstract_data_type

// ADT_Reversed is an operation that reverse the container and return new one.
type ADT_Reversed[CONTAINER any] interface {
	Reversed() (con CONTAINER, err Error)
}

// ADT_Sorted is an operation that sort the container elements and return new one.
type ADT_Sorted[CONTAINER any] interface {
	Sorted() (con CONTAINER, err Error)
}

// ADT_Push is an operation
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type ADT_Push[ELEMENT any] interface {
	// Push adds an element to the collection
	Push(el ELEMENT) (err Error)
}

// ADT_Push is an operation
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type ADT_Pop[ELEMENT any] interface {
	// Pop, which removes the most recently added element and return it.
	Pop() (el ELEMENT, err Error)
}

// ADT_Peek is an operation
// https://en.wikipedia.org/wiki/Peek_(data_type_operation)
type ADT_Peek[ELEMENT any] interface {
	// Peek or `Top()` is an operation on certain abstract data types, specifically sequential collections such as stacks and queues,
	// which returns the value of the top ("front") of the collection without removing the element from the collection.
	Peek() (el ELEMENT, err Error)
}

// ADT_Push is an operation
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type ADT_Enqueue[ELEMENT any] interface {
	// Enqueue adding an element to the rear of the queue.
	Enqueue(el ELEMENT) (err Error)
}

// ADT_Dequeue is an operation
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type ADT_Dequeue[ELEMENT any] interface {
	// Dequeue removing an element from the front of queue and return it.
	Dequeue() (el ELEMENT, err Error)
}

// ADT_Insert is an operation
type ADT_Insert[ELEMENT any] interface {
	// Insert will add the `el` in the offset of the container by move elements after offset to `offset+len(el)`.
	Insert(el ELEMENT, offset ElementIndex) (err Error)
}

// ADT_Append is an operation
type ADT_Append[ELEMENT any] interface {
	// Append will add the `el` to the end of the container.
	Append(el ELEMENT) (err Error)
}

// ADT_Prepend is an operation
type ADT_Prepend[ELEMENT any] interface {
	// Prepend will add the `el` to the beginning of the container.
	Prepend(el ELEMENT) (err Error)
}

// ADT_Clear is an operation
type ADT_Clear interface {
	// Clear will remove all elements.
	Clear() (err Error)
}
