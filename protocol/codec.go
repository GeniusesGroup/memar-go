/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Codec wraps some other interfaces that need an data structure be a codec.
// Protocols that have just one specific logic NEED to implement this interface e.g. HTTP, MP3, AVI, ...
// Others can implement other Codec e.g. Syllab, JSON, XML, HTML, CSS, ...
// https://en.wikipedia.org/wiki/Codec
type Codec interface {
	MediaType() MediaType
	CompressType() CompressType

	Decoder
	Encoder
}

// Decoder is the interface that wraps the Decode method.
type Decoder interface {
	// Decode read and decode data until end of needed data or occur error.
	// Unlike io.ReadFrom() it isn't read until EOF and just read needed data.
	Decode(source Buffer) (err Error)
}

// Encoder is the interface that wraps the Encode & Len methods.
type Encoder interface {
	// Encode writes serialized(encoded) data to destination until there's no more data to write.
	// Return any error that occur in buffer logic e.g. timeout error in socket, ...
	Encode(destination Buffer) (err Error)

	CodecLength
}

// In computer science, marshalling or marshaling (US spelling) is the process of
// transforming the memory representation of an object into a data format suitable for storage or transmission.
// https://en.wikipedia.org/wiki/Marshalling_(computer_science)

// Unmarshaler is the interface that wraps the Unmarshal method.
type Unmarshaler interface {
	// Unmarshal reads and decode data from given slice until end of needed data or occur error.
	Unmarshal(source []byte) (n NumberOfElement, err Error)
}

// Marshaler is the interface that wraps the Marshal methods.
type Marshaler interface {
	// Marshal write serialized(encoded) data to given slice from len to max cap and save marshal state for future call.
	Marshal(destination []byte) (n NumberOfElement, err Error)

	CodecLength
}

type CodecLength interface {
	// SerializationLength return value ln, that is the max number of bytes that will written as encode data by Encode()||Marshal()
	// 0 means no data and -1 means can't tell until full write.
	// Due to prevent performance penalty, Implementors can return max number instead of actual number of length.
	SerializationLength() (ln NumberOfElement)
}
