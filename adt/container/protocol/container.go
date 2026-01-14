/* For license and copyright information please see the LEGAL file in the code repository */

package container_p

import (
	adt_p "memar/adt/protocol"
)

type Container[ELEMENT Element] interface {
	adt_p.ADT

	Accessor[ELEMENT]
	Assignment[ELEMENT]

	Clear
	Reversed
	Sorted
	Resize

	Field_Capacity
	Field_OccupiedLength
	Field_AvailableLength
	// Field_ExpectedLength
}

type Container_READONLY[ELEMENT Element] interface {
	adt_p.ADT

	Accessor[ELEMENT]

	Field_ExpectedLength
}
