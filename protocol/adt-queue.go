/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Queue is queue data structure.
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type Queue[ELEMENT any] interface {
	ADT_Dequeue[ELEMENT]
	ADT_Enqueue[ELEMENT]
}

// ADT_Enqueue is an element operation
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type ADT_Enqueue[ELEMENT any] interface {
	// Enqueue adding an element to the rear of the queue.
	// NOT RECOMMENDED but implementation CAN just call `Prepend()`
	Enqueue(el ELEMENT) (err Error)
}

// ADT_Dequeue is an element operation
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type ADT_Dequeue[ELEMENT any] interface {
	// Dequeue removing an element from the front of queue and return it.
	// NOT RECOMMENDED but implementation CAN just call `Pop()`
	Dequeue() (el ELEMENT, err Error)
}
