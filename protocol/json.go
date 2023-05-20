/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// JSON is the interface that must implement by any struct that can be a JSON object.
// Standards by https://www.json.org/json-en.html
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON
type JSON interface {
	JsonUnmarshaler
	JsonMarshaler
}

// JsonUnmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify
// - The input can be assumed to be a valid encoding of a JSON value.
// - FromJSON must declare as comment that it copy the JSON data or not, if it wishes to retain the data after returning.
// - By convention, to approximate the behavior of Unmarshal itself, Unmarshalers implement FromJSON([]byte("null")) as a no-op.
type JsonUnmarshaler interface {
	// FromJSON or UnmarshalFromJSON decode JSON to the struct. API is same as codec.UnmarshalFrom
	// actually payload is a byte slice buffer interface but due to prevent unnecessary memory allocation use simple []byte
	FromJSON(payload []byte) (remaining []byte, err Error)
}

// JsonMarshaler is the interface implemented by types that can marshal themselves into valid JSON.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/parse
// - Some types can't encode to JSON like large integers, ... and get error on runtime or code generator phase.
type JsonMarshaler interface {
	// ToJSON or MarshalToJSON encode the data to JSON format. API is same as codec.MarshalTo
	// actually payload is a byte slice buffer interface but due to prevent unnecessary memory allocation use simple []byte
	ToJSON(payload []byte) (added []byte, err Error)

	// LenAsJSON return whole calculated length of JSON encoded of the struct
	LenAsJSON() int
}
