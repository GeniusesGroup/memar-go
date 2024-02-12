/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Stack is stack data structure.
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type ADT_Stack[ELEMENT any] interface {
	ADT_Push[ELEMENT]
	ADT_Pop[ELEMENT]
}
