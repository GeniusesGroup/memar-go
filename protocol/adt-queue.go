/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Queue is queue data structure.
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type ADT_Queue[ELEMENT any] interface {
	ADT_Dequeue[ELEMENT]
	ADT_Enqueue[ELEMENT]
}
