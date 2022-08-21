/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Codec wraps some other interfaces to act an object as a codec.
// Differences:
// - Marshal() don't think about any other parts and make a byte slice and serialize data to it.
// - MarshalTo() care about the fact that serialized data must wrap with other data and serialize data in the given byte slice.
// - Encode() like MarshalTo() but encode to a buffer not a byte slice by Buffer interface methods. buf can be a temp or final write location.
type Codec interface {
	MediaType() MediaType
	CompressType() CompressType

	Decoder
	Encoder

	Unmarshaler
	Marshaler
}

// Decoder is the interface that wraps the Decode method.
type Decoder interface {
	// Decode read and decode data until end of needed data or occur error.
	// Unlike Buffer.ReadFrom() it isn't read until EOF and just read needed data.
	Decode(reader Codec) (n int, err Error)
}

// Encoder is the interface that wraps the Encode & Len methods.
type Encoder interface {
	// Encode writes serialized(encoded) data to writer until there's no more data to write.
	// It like Buffer.WriteTo() with a very tiny difference that this method know about the serialized data but WriteTo just know marshaled data is a binary data.
	// It use in old OSs fashion or stream writing in large data length. Old OSs do some logic in kernel e.g. IP/TCP packeting, ... that need heavy context switching logic.
	// It will care about how to write data to respect performance. Usually by make temp fixed size buffer like bufio package.
	Encode(writer Codec) (n int, err Error)
	// Len return value n is the number of bytes that will written as encode data. 0 means no data and -1 means can't tell until full write.
	Len() (ln int)
}

// Unmarshaler is the interface that wraps the Unmarshal method.
type Unmarshaler interface {
	// Unmarshal reads and decode data from given slice until end of data or occur error
	Unmarshal(data []byte) (n int, err Error)
	UnmarshalFrom(data []byte) (remaining []byte, err Error)
}

// Marshaler is the interface that wraps the Marshal methods.
// Return any error that occur in logic e.g. timeout error in socket, ...
type Marshaler interface {
	// Marshal serialized(encoded) data and return the byte slice hold serialized data.
	Marshal() (data []byte, err Error)
	// MarshalTo serialized(encoded) data to given slice from len to max cap and save marshal state for future call.
	// It is very similar to Read() in Reader interface but with one difference behavior that this method don't need temporary buffer
	MarshalTo(data []byte) (added []byte, err Error)
	// Len return value n that is the number of bytes that will written by Marshal()||MarshalTo()
	Len() (ln int)
}

type SerializeLen interface {
	ComputeLen() (ln int)
	Len() (ln int)
}
