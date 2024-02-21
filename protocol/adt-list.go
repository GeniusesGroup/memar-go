/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// list or sequence is an abstract data type that represents a finite number of ordered values,
// where the same value may occur more than once.
// An instance of a list is a computer representation of the mathematical concept of a tuple or finite sequence;
// the (potentially) infinite analog of a list is a stream.
// Lists are a basic example of containers, as they contain other values.
// If the same value occurs multiple times, each occurrence is considered a distinct item.
// https://en.wikipedia.org/wiki/List_(abstract_data_type)
type List[ELEMENT any] interface {
	ADT_Head[ELEMENT]
	ADT_Tail[ELEMENT]

	Container[ELEMENT]
}

type ADT_Head[ELEMENT any] interface {
	// Head will return first element of the container.
	Head() (el ELEMENT, err Error)
}

type ADT_Tail[ELEMENT any] interface {
	// Tail will return last element of the container.
	// an operation for referring to the list consisting of all the components of a list except for its first (this is called the "tail" of the list.)
	Tail() (el ELEMENT, err Error)
}
