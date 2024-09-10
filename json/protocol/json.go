/* For license and copyright information please see the LEGAL file in the code repository */

package json_p

import (
	adt_p "memar/adt/protocol"
	buffer_p "memar/buffer/protocol"
	error_p "memar/error/protocol"
)

// JSON is the interface that must implement by any struct that can be a JSON object.
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
type JSON interface {
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

	Length
}

// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify
type Unmarshaler interface {
	// UnmarshalFromJSON like `FromJSON()` decode JSON to the desire structure. API is same as `codec.Unmarshal()`
	UnmarshalFromJSON(source []byte) (n adt_p.NumberOfElement, err error_p.Error)
}

// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/parse
type Marshaler interface {
	// MarshalToJSON is same as `ToJSON()` encode the data to JSON format. API is same as `codec.Marshal()`
	MarshalToJSON(destination []byte) (n adt_p.NumberOfElement, err error_p.Error)

	Length
}

// Length is same as CodecLength
type Length interface {
	// LengthAsJSON return whole calculated length of JSON encoded of the struct
	LengthAsJSON() adt_p.NumberOfElement
}
