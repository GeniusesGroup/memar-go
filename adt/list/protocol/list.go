/* For license and copyright information please see the LEGAL file in the code repository */

package list_p

import (
	adt_p "memar/adt/protocol"
)

// list or sequence is an abstract data type that represents a finite number of ordered values,
// where the same value may occur more than once.
// An instance of a list is a computer representation of the mathematical concept of a tuple or finite sequence;
// the (potentially) infinite analog of a list is a stream.
// Lists are a basic example of containers, as they contain other values.
// If the same value occurs multiple times, each occurrence is considered a distinct item.
// https://en.wikipedia.org/wiki/List_(abstract_data_type)
// Other Languages:
// - https://docs.python.org/3/tutorial/datastructures.html
type List[ELEMENT adt_p.Element] interface {
	Head[ELEMENT]
	Tail[ELEMENT]

	adt_p.Container[ELEMENT]
}

type Head[ELEMENT adt_p.Element] interface {
	// Head will return first element of the container.
	Head() (el ELEMENT, err error_p.Error)
}

type Tail[ELEMENT adt_p.Element] interface {
	// Tail will return last element of the container.
	// an operation for referring to the list consisting of all the components of a list except for its first (this is called the "tail" of the list.)
	Tail() (el ELEMENT, err error_p.Error)
}
