/* For license and copyright information please see the LEGAL file in the code repository */

package stack_p

import (
	adt_p "memar/adt/protocol"
)

// Stack is stack data structure.
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack[ELEMENT adt_p.Element] interface {
	adt_p.Push[ELEMENT]
	adt_p.Pop[ELEMENT]
}
