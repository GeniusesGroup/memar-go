/* For license and copyright information please see LEGAL file in repository */

package protocol

import "io"

// Codec wraps some other interfaces!
// Differencess:
// - Marshal() don't think about any other parts and make a byte slice and serialize data to it.
// - MarshalTo() care about the fact that serialized data must wrap with other data and serialize data in the given byte slice.
// - Encode() like MarshalTo() but encode to a buf not a byte slice by Buffer interface methods! buf can be a temp or final write location.
// - WriteTo() is OS changed data location so it will care about how to write data to respect performance. Usually by make temp fixed size buffer like bufio package.
type Codec interface {
	MediaType() string // http://www.iana.org/assignments/media-types/media-types.xhtml
	CompressType() string

	Decoder
	Encoder

	UnMarshaler
	Marshaler

	io.WriterTo
	io.ReaderFrom
	// io.ReaderAt
	// io.WriteAt
}

// Decoder is the interface that wraps the Decode method.
//
// Decode read and decode data from buffer until end of data or occur error.
type Decoder interface {
	Decode(buf Buffer) (err Error)
}

// Encoder is the interface that wraps the Encode & Len methods.
//
// Encode writes serialized(encoded) data to buf until there's no more data to write!
// Len return value n is the number of bytes that will written as encode data.
type Encoder interface {
	Encode(buf Buffer)
	Len() (ln int)
}

// UnMarshaler is the interface that wraps the UnMarshal method.
//
// UnMarshal reads and decode data from given slice until end of data or occur error
type UnMarshaler interface {
	UnMarshal(data []byte) (err Error)
}

// Marshaler is the interface that wraps the Marshal method.
//
// Marshal serialized(encoded) data and return the byte slice
// MarshalTo serialized(encoded) data to given slice. Slice cap-len must >= Len()
// Len return value n that is the number of bytes that will written by Marshal()||MarshalTo()
type Marshaler interface {
	Marshal() (data []byte)
	MarshalTo(data []byte) []byte
	Len() (ln int)
}

type SerializeLen interface {
	ComputeLen() (ln int)
	Len() (ln int)
}
