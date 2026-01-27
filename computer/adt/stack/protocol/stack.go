/* For license and copyright information please see the LEGAL file in the code repository */

package stack_p

import (
	container_p "memar/adt/container/protocol"
)

// Stack is stack data structure.
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack[ELEMENT container_p.Element] interface {
	container_p.Push[ELEMENT]
	container_p.Pop[ELEMENT]
}
