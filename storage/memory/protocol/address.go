/* For license and copyright information please see the LEGAL file in the code repository */

package memory_p

import (
	error_p "memar/process/error/protocol"
	logic_p "memar/math/logic/protocol"
)

type Field_MemoryAddress interface {
	MemoryAddress() MemoryAddress
}

type Method_SetMemoryAddress interface {
	// SetMemoryAddress is unsafe method and it is better to use in rare cases.
	// Conversion of a T1 to Pointer to T2.
	SetMemoryAddress(mAdd MemoryAddress) (err error_p.Error)
}

type MemoryAddress interface {
	logic_p.Equivalence[MemoryAddress]

	// string_p.Stringer
}
