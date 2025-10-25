/* For license and copyright information please see the LEGAL file in the code repository */

package codec_p

import (
	buffer_p "memar/buffer/protocol"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

// Codec wraps some other interfaces that need an data structure be a codec.
// Protocols that have just one specific logic NEED to implement this interface e.g. HTTP, MP3, AVI, ...
// Others can implement other Codec e.g. Syllab, JSON, XML, HTML, CSS, ...
// https://en.wikipedia.org/wiki/Codec
type Codec /*[BUF Buffer]*/ interface {
	Decoder /*[BUF]*/
	Encoder /*[BUF]*/

	datatype_p.DataType
}

// Decoder is the interface that wraps the Decode method.
type Decoder /*[BUF Buffer]*/ interface {
	// Decode read and decode data until end of needed data or occur error.
	// Unlike io.ReadFrom() it isn't read until EOF and just read needed data.
	Decode(source buffer_p.Buffer) (err error_p.Error)
}

// Encoder is the interface that wraps the Encode & Field_Length methods.
type Encoder /*[BUF Buffer]*/ interface {
	// Encode writes serialized(encoded) data to destination until there's no more data to write.
	// Return any error that occur in buffer logic e.g. timeout error in socket, ...
	Encode(destination buffer_p.Buffer) (err error_p.Error)

	Field_Length
}
