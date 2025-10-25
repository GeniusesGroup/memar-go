/* For license and copyright information please see the LEGAL file in the code repository */

package codec_p

import (
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

type Marshalling interface {
	Unmarshaler
	Marshaler

	datatype_p.DataType
}

// In computer science, marshalling or marshaling (US spelling) is the process of
// transforming the memory representation of an object into a data format suitable for storage or transmission.
// https://en.wikipedia.org/wiki/Marshalling_(computer_science)

// Unmarshaler is the interface that wraps the Unmarshal method.
type Unmarshaler interface {
	// Unmarshal reads and decode data from given slice until end of needed data or occur error.
	Unmarshal(source []byte) (n Length, err error_p.Error)
}

// Marshaler is the interface that wraps the Marshal & Field_Length methods.
type Marshaler interface {
	// Marshal write serialized(encoded) data to given slice from len to max cap and save marshal state for future call.
	Marshal(destination []byte) (n Length, err error_p.Error)

	Field_Length
}
