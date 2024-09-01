/* For license and copyright information please see the LEGAL file in the code repository */

package syllab_p

import (
	adt_p "memar/adt/protocol"
	buffer_p "memar/buffer/protocol"
	"memar/protocol"
)

// Syllab is the interface that must implement by any struct to be a Syllab object transmittable over networks.
// Standards in https://github.com/GeniusesGroup/memar/blob/main/Syllab.md
type Syllab interface {
	SyllabMarshaler
	SyllabUnmarshaler
}

type SyllabUnmarshaler interface {
	// CheckSyllab usually just check LenOfSyllabStack not greater than len of given payload. and call just before decode payload.
	CheckSyllab(source buffer_p.Buffer) (err protocol.Error)

	// FromSyllab ready given payload for get accessors methods.
	// - Due to strongly suggest to use fields get accessors methods, below method just change under hood buffer if it isn't struct.
	// - It can return LenOfSyllabStack()-1 as end of stack in payload, but it will be runtime logic. compiler can inline and do -1 in compile time easily.
	FromSyllab(source buffer_p.Buffer, stackIndex adt_p.ElementIndex)
}

type SyllabMarshaler interface {
	// ToSyllab encode the struct pointer to Syllab format
	// in non embed struct usually `stackIndex = 0` & `heapIndex = {{rn}}.LenOfSyllabStack()` as heap start index || end of stack size.
	ToSyllab(destination buffer_p.Buffer, stackIndex, heapIndex adt_p.ElementIndex) (freeHeapIndex adt_p.ElementIndex)

	// LenAsSyllab return whole calculated length of Syllab encoded of the struct
	// default is simple as `return ({{RecName}}.LenOfSyllabStack() + {{RecName}}.LenOfSyllabHeap())`
	LenAsSyllab() adt_p.NumberOfElement

	// LenOfSyllabStack return calculated stack length of Syllab encoded of the struct
	LenOfSyllabStack() adt_p.NumberOfElement
	// LenOfSyllabStack return calculated heap length of Syllab encoded of the struct
	LenOfSyllabHeap() adt_p.NumberOfElement
}
