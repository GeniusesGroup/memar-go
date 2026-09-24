/* For license and copyright information please see the LEGAL file in the code repository */

package syllab_p

import (
	container_p "memar/computer/adt/container/protocol"
)

// Field_Lengths is same as codec_p.Field_Length in many ways
type Field_Lengths interface {
	// Syllab_Length return whole calculated length of Syllab encoded of the struct
	// default is simple as `return (self.Syllab_StackLength() + self.Syllab_HeapLength())`
	Syllab_Length() Length

	// Syllab_StackLength return calculated stack length of Syllab encoded of the struct
	Syllab_StackLength() Length_Stack
	// Syllab_HeapLength return calculated heap length of Syllab encoded of the struct
	Syllab_HeapLength() Length_Heap
}

type Length = container_p.NumberOfElement
type Length_Stack = container_p.NumberOfElement
type Length_Heap = container_p.NumberOfElement
