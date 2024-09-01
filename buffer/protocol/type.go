/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_p

type Type uint8

const (
	Type_Unset Type = iota

	//
	Type_Flat

	// In computer science, a circular buffer, circular queue, cyclic buffer or ring buffer is a data structure
	// that uses a single, fixed-size buffer as if it were connected end-to-end.
	// This structure lends itself easily to buffering data streams. There were early circular buffer implementations in hardware.
	//
	// https://en.wikipedia.org/wiki/Circular_buffer
	Type_Circular

	// It means buffer not reuse when need copy more data to it and it can't be resize!
	// Use to implement buffer as `zero-copy` feature, But it can cause memory leak.
	// Developers must be aware about this, But we introduce `Memar framework` to fix any problems
	// by introduce new protocol to replace old protocols like TCP, HTTP, ...
	//
	// https://en.wikipedia.org/wiki/Zero-copy
	Type_ZeroCopy

	// https://en.wikipedia.org/wiki/Sparse_file
	Type_Sparse
)
