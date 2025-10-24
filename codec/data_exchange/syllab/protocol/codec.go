/* For license and copyright information please see the LEGAL file in the code repository */

package syllab_p

import (
	buffer_p "memar/buffer/protocol"
	error_p "memar/process/error/protocol"
)

// Codec is 
type Codec interface {
	Encoder
	Decoder
}

type Decoder interface {
	// CheckSyllab usually just check Syllab_StackLength not greater than len of given payload. and call just before decode payload.
	CheckSyllab(source buffer_p.Buffer) (err error_p.Error)

	// FromSyllab ready given payload for get accessors methods.
	// - Due to strongly suggest to use fields get accessors methods, below method just change under hood buffer if it isn't struct.
	// - It can return Syllab_StackLength()-1 as end of stack in payload, but it will be runtime logic. compiler can inline and do -1 in compile time easily.
	FromSyllab(source buffer_p.Buffer, stackIndex Index_Stack) (err error_p.Error)
}

type Encoder interface {
	// ToSyllab encode the struct pointer to Syllab format
	// in non embed struct usually `stackIndex = 0` & `heapIndex = self.Syllab_StackLength()` as heap start index || end of stack size.
	ToSyllab(destination buffer_p.Buffer, stackIndex Index_Stack, heapIndex Index_Heap) (freeHeapIndex Index_Heap, err error_p.Error)

	Field_Lengths
}
