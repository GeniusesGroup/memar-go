/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_p

import (
	adt_p "memar/adt/protocol"
	"memar/protocol"
)

// In computer science, a data buffer (or just buffer) is a region of a memory used to store data temporarily
// while it is being moved from one place to another.
//
// https://en.wikipedia.org/wiki/Data_buffer
type Buffer interface {
	BufferType() Type

	protocol.ObjectLifeCycle
	// Init(opt BufferOptions)

	Buffer_Index
	Buffer_Sizer

	adt_p.Container[byte]

	adt_p.Compare[Buffer]
	adt_p.Concat[Buffer]
	adt_p.Replace_Elements[Buffer]

	adt_p.Split_Element[Buffer, byte]
	adt_p.Split_Offset[Buffer, byte]

	// If source is a `Split` result, no copy action need and just increase buffer write index.
	protocol.DataType_Clone[Buffer]
	protocol.DataType_Copy[Buffer]
}

// BufferOptions:
// 	`Blocking bool`:
// 		All `Get` methods block the caller if desire payload in buffer not exist.
// 		If caller don't want to block, It MUST call `Length` and check data already exist.
// 		OR:::
// 		If some data is available but not as `limit`, `Get` conventionally
// 		returns what is available instead of waiting for more.
// `ln NumberOfElement`:
//
// `Resizable bool`:
//
//  ...
//

// Even if ReadAt returns n < len(p), it may use all of p as scratch
// space during the call. If some data is available but not len(p) bytes,
// ReadAt blocks until either all the data is available or an error occurs.
// In this respect ReadAt is different from Read.

type Buffer_Index interface {
	ReadIndex() adt_p.ElementIndex
	WriteIndex() adt_p.ElementIndex

	SetReadIndex(di adt_p.ElementIndex)
	SetWriteIndex(di adt_p.ElementIndex)
}

type Buffer_Sizer interface {
	// UnreadLength returns how many bytes are not read(ReadIndex to WriteIndex) in the buffer.
	UnreadLength() adt_p.NumberOfElement
}
