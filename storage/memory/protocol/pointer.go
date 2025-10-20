/* For license and copyright information please see the LEGAL file in the code repository */

package memory_p

import (
	logic_p "memar/math/logic/protocol"
	error_p "memar/process/error/protocol"
)

type Field_Pointer interface {
	Pointer() Pointer // uintptr
}

// Pointers are the most primitive type of reference.
// A pointer is a simple, more concrete implementation of the more abstract reference data type.
// Pointer is underlying memory structure for any datatype or capsule.
// But it CAN implement auto by compiler like `Deinit()`, `Drop()`, ...
// 
// https://en.wikipedia.org/wiki/Pointer_(computer_programming)
type Pointer interface {
	Field_MemoryAddress
	Method_SetMemoryAddress

	// container_p.Container[]
	logic_p.Equivalence[Pointer]
}
