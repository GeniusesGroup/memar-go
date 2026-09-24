/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	array_p "memar/computer/adt/array/protocol"
	marker_p "memar/computer/adt/marker/protocol"
	container_p "memar/computer/adt/container/protocol"
	// buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
	logic_p "memar/math/logic/protocol"
)

type String interface {
	CharacterEncoding() CharacterEncoding
	// Buffer() buffer_p.Buffer

	array_p.Dynamic[Character]

	marker_p.Clone[String]
	marker_p.Copy[String]

	container_p.Compare[String]
	container_p.Concat[String]
	// Join[String]
	container_p.Split_Element[String, Character]
	container_p.Split_Offset[String, Character]

	logic_p.Equivalence[String]
	// If source is a `Split` result, no copy action need and just increase buffer write index.
}

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join
// Join is an operation
type Join[STR String] interface {
	// The join() method of Array instances creates and returns a new string by concatenating all of the elements in this array,
	// separated by commas or a specified separator string.
	Join(sep STR, con ...STR) (s STR, err error_p.Error)
}
