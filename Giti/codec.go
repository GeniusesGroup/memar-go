/* For license and copyright information please see LEGAL file in repository */

package giti

import "io"

type Codec interface {
	Decoder
	Encoder

	io.WriterTo
	io.ReaderFrom
}

// ReaderFrom is the interface that wraps the Decode method.
//
// Decode read and decode data from buffer until end of data or error.
type Decoder interface {
	Decode(buf Buffer) (err Error)
}

// Encoder is the interface that wraps the Encode & Len methods.
//
// Encode writes encoded data to buf until there's no more data to write whole encoded data.
// Len return value n is the number of bytes that will written as encode data.
type Encoder interface {
	Encode(buf Buffer)
	Len() int
}
