/* For license and copyright information please see the LEGAL file in the code repository */

package json_p

import (
	buffer_p "memar/buffer/protocol"
	error_p "memar/process/error/protocol"
)

// Codec JSON is the interface that must implement by any struct that can be a JSON object.
// Standards by https://www.json.org/json-en.html
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON
//
// Decode:
// - The input can be assumed to be a valid encoding of a JSON value.
// - FromJSON must declare as comment that it copy the JSON data or not, if it wishes to retain the data after returning.
// - By convention, to approximate the behavior of Unmarshal itself, Unmarshaler implement FromJSON([]byte("null")) as a no-op.
//
// Encode:
// - Some types can't encode to JSON like large integers, ... and get error on runtime or code generator phase.
type Codec interface {
	Decoder
	Encoder
}

type Decoder interface {
	// FromJSON decode JSON to the desire structure. API is same as `codec.Decode()`
	FromJSON(source buffer_p.Buffer) (err error_p.Error)
}
type Encoder interface {
	// FromJSON encode the structure to JSON format. API is same as `codec.Encoder()`
	ToJSON(destination buffer_p.Buffer) (err error_p.Error)

	Field_Length
}
