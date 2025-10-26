/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_p

import (
	container_p "memar/adt/container/protocol"
	capsule_p "memar/computer/capsule/protocol"
	datatype_p "memar/computer/datatype/protocol"
	memory_p "memar/storage/memory/protocol"
)

// In computer science, a data buffer (or just buffer) is a region of a memory used to store data temporarily
// while it is being moved from one place to another.
//
// https://en.wikipedia.org/wiki/Data_buffer
//
// Other frameworks:
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer
type Buffer interface {
	datatype_p.Field_ID

	capsule_p.LifeCycle
	// Init(opt BufferOptions)

	Field_Indexes
	Method_Indexes
	Field_Lengths

	container_p.Container[byte]

	container_p.Compare[Buffer]
	container_p.Concat[Buffer]
	container_p.Replace_Elements[Buffer]

	container_p.Split_Element[Buffer, byte]
	container_p.Split_Offset[Buffer, byte]

	// If source is a `Split` result, no copy action need and just increase buffer write index.
	memory_p.Clone[Buffer]
	memory_p.Copy[Buffer]
}

// BufferOptions:
// 	`Blocking bool`:
// 		All `Get` methods block the caller if desire payload in buffer not exist.
// 		If caller don't want to block, It MUST call `Length` and check data already exist.
// 		OR:::
// 		If some data is available but not as `limit`, `Get` conventionally
// 		returns what is available instead of waiting for more.
// `ln Length`:
//
// `Resizable bool`:
//
//  ...
//

// Even if ReadAt returns n < len(p), it may use all of p as scratch
// space during the call. If some data is available but not len(p) bytes,
// ReadAt blocks until either all the data is available or an error occurs.
// In this respect ReadAt is different from Read.
