/* For license and copyright information please see the LEGAL file in the code repository */

package queue_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/error/protocol"
)

// Queue is queue data structure.
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type Queue[ELEMENT container_p.Element] interface {
	Dequeue[ELEMENT]
	Enqueue[ELEMENT]
}

// Enqueue is an element operation
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type Enqueue[ELEMENT container_p.Element] interface {
	// Enqueue adding an element to the rear of the queue.
	// NOT RECOMMENDED but implementation CAN just call `Prepend()`
	Enqueue(el ELEMENT) (err error_p.Error)
}

// Dequeue is an element operation
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type Dequeue[ELEMENT container_p.Element] interface {
	// Dequeue removing an element from the front of queue and return it.
	// NOT RECOMMENDED but implementation CAN just call `Pop()`
	Dequeue() (el ELEMENT, err error_p.Error)
}
