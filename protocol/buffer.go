/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// In computer science, a data buffer (or just buffer) is a region of a memory used to store data temporarily
// while it is being moved from one place to another.
//
// https://en.wikipedia.org/wiki/Data_buffer
type Buffer interface {
	BufferType() BufferType

	ObjectLifeCycle
	// Init(opt BufferOptions)

	ADT_Container[[]byte]

	Buffer_Index
	Buffer_Sizer
	Buffer_Resize

	// Codec

	// DataType_Locker
	DataType_Clone[Buffer]
	DataType_Copy[Buffer]
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

type BufferType uint8

const (
	BufferType_Unset BufferType = iota

	//
	BufferType_Flat

	// In computer science, a circular buffer, circular queue, cyclic buffer or ring buffer is a data structure
	// that uses a single, fixed-size buffer as if it were connected end-to-end.
	// This structure lends itself easily to buffering data streams. There were early circular buffer implementations in hardware.
	//
	// https://en.wikipedia.org/wiki/Circular_buffer
	BufferType_Circular

	// It means buffer not reuse when need copy more data to it and it can't be resize!
	// Use to implement buffer as `zero-copy` feature, But it can cause memory leak.
	// Developers must be aware about this, But we introduce `Memar framework` to fix any problems
	// by introduce new protocol to replace old protocols like TCP, HTTP, ...
	//
	// https://en.wikipedia.org/wiki/Zero-copy
	BufferType_ZeroCopy

	// https://en.wikipedia.org/wiki/Sparse_file
	BufferType_Sparse
)

// Even if ReadAt returns n < len(p), it may use all of p as scratch
// space during the call. If some data is available but not len(p) bytes,
// ReadAt blocks until either all the data is available or an error occurs.
// In this respect ReadAt is different from Read.

type Buffer_Resize interface {
	Resize(ln NumberOfElement) Error
	Resized() bool
	// Resizable returns true if the Buffer can be resized, or false if not.
	Resizable() bool
}

type Buffer_Index interface {
	ReadIndex() ElementIndex
	WriteIndex() ElementIndex

	SetReadIndex(di ElementIndex)
	SetWriteIndex(di ElementIndex)
}

type Buffer_Sizer interface {
	// UnreadLength returns how many bytes are not read(ReadIndex to WriteIndex) in the buffer.
	UnreadLength() NumberOfElement
}

type Buffer_Circular interface {
	Head() ElementIndex
	Tail() ElementIndex
}
