/* For license and copyright information please see the LEGAL file in the code repository */

package codec_p

import (
	adt_p "memar/adt/protocol"
	buffer_p "memar/buffer/protocol"
	"memar/protocol"
)

// Codec wraps some other interfaces that need an data structure be a codec.
// Protocols that have just one specific logic NEED to implement this interface e.g. HTTP, MP3, AVI, ...
// Others can implement other Codec e.g. Syllab, JSON, XML, HTML, CSS, ...
// https://en.wikipedia.org/wiki/Codec
type Codec /*[BUF Buffer]*/ interface {
	Decoder /*[BUF]*/
	Encoder /*[BUF]*/

	protocol.DataType
}

// Decoder is the interface that wraps the Decode method.
type Decoder /*[BUF Buffer]*/ interface {
	// Decode read and decode data until end of needed data or occur error.
	// Unlike io.ReadFrom() it isn't read until EOF and just read needed data.
	Decode(source buffer_p.Buffer) (err protocol.Error)
}

// Encoder is the interface that wraps the Encode & CodecLength methods.
type Encoder /*[BUF Buffer]*/ interface {
	// Encode writes serialized(encoded) data to destination until there's no more data to write.
	// Return any error that occur in buffer logic e.g. timeout error in socket, ...
	Encode(destination buffer_p.Buffer) (err protocol.Error)

	CodecLength
}

// In computer science, marshalling or marshaling (US spelling) is the process of
// transforming the memory representation of an object into a data format suitable for storage or transmission.
// https://en.wikipedia.org/wiki/Marshalling_(computer_science)

// Unmarshaler is the interface that wraps the Unmarshal method.
type Unmarshaler interface {
	// Unmarshal reads and decode data from given slice until end of needed data or occur error.
	Unmarshal(source []byte) (n adt_p.NumberOfElement, err protocol.Error)
}

// Marshaler is the interface that wraps the Marshal & CodecLength methods.
type Marshaler interface {
	// Marshal write serialized(encoded) data to given slice from len to max cap and save marshal state for future call.
	Marshal(destination []byte) (n adt_p.NumberOfElement, err protocol.Error)

	CodecLength
}

type CodecLength interface {
	// SerializationLength return value ln, that is the max number of bytes that will written as encode data by Encode()||Marshal()
	// 0 means no data and -1 means can't tell until full write.
	// Due to prevent performance penalty, Implementors can return max number instead of actual number of length.
	SerializationLength() (ln adt_p.NumberOfElement)
}
