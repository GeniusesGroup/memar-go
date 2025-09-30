/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	array_p "memar/adt/array/protocol"
	container_p "memar/adt/container/protocol"
	// buffer_p "memar/buffer/protocol"
	error_p "memar/error/protocol"
	logic_p "memar/math/logic/protocol"
	memory_p "memar/storage/memory/protocol"
)

// In computer programming, a string is traditionally a sequence of characters,
// either as a literal constant or as some kind of variable.
// The latter may allow its elements to be mutated and the length changed, or it may be fixed (after creation).
// A string is generally considered as a data type and is often implemented as an array data structure of bytes
// (or words) that stores a sequence of elements, typically characters, using some character encoding.
// String may also denote more general arrays or other sequence (or list) data types and structures.
// https://en.wikipedia.org/wiki/String_(computer_science)
type String interface {
	CharacterEncoding() CharacterEncoding
	// Buffer() buffer_p.Buffer

	array_p.Dynamic[Character]

	container_p.Compare[String]
	container_p.Concat[String]
	// Join[String]
	container_p.Split_Element[String, Character]
	container_p.Split_Offset[String, Character]

	logic_p.Equivalence[String]
	// If source is a `Split` result, no copy action need and just increase buffer write index.
	memory_p.Clone[String]
	memory_p.Copy[String]
}

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join
// Join is an operation
type Join[STR String] interface {
	// The join() method of Array instances creates and returns a new string by concatenating all of the elements in this array,
	// separated by commas or a specified separator string.
	Join(sep STR, con ...STR) (s STR, err error_p.Error)
}
