/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// In computer programming, a string is traditionally a sequence of characters,
// either as a literal constant or as some kind of variable.
// The latter may allow its elements to be mutated and the length changed, or it may be fixed (after creation).
// A string is generally considered as a data type and is often implemented as an array data structure of bytes
// (or words) that stores a sequence of elements, typically characters, using some character encoding.
// String may also denote more general arrays or other sequence (or list) data types and structures.
// https://en.wikipedia.org/wiki/String_(computer_science)
type String interface {
	CharacterEncoding() CharacterEncoding

	Array_Dynamic[Character]

	ADT_Compare[String]
	ADT_Concat[String]
	ADT_Split_Element[String, Character]
	ADT_Split_Offset[String, Character]

	// If source is a `Split` result, no copy action need and just increase buffer write index.
	DataType_Clone[String]
	DataType_Copy[String]
}

// Stringer code the data to/from human readable format. It can be any other format like JSON(not recommended).
type Stringer[STRING String] interface {
	Stringer_To[STRING]
	Stringer_From[STRING]
}

type Stringer_To[STRING String] interface {
	ToString() (str STRING, err Error)
}
type Stringer_From[STRING String] interface {
	FromString(str STRING) (err Error)
}

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join
// ADT_Concat is an operation
type ADT_Join[STRING String] interface {
	// The join() method of Array instances creates and returns a new string by concatenating all of the elements in this array,
	// separated by commas or a specified separator string.
	Join(sep STRING) (s STRING, err Error)
}
