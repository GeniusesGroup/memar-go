/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type String interface {
	CharacterEncoding() CharacterEncoding

	Array_Dynamic[any]
}

// Stringer code the data to/from human readable format. It can be any other format like JSON(not recommended).
type Stringer interface {
	Stringer_To
	Stringer_From
}

type Stringer_To interface {
	ToString() (str string, err Error)
}
type Stringer_From interface {
	FromString(str string) (err Error)
}

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join
// ADT_Concat is an operation
type ADT_Join interface {
	// The join() method of Array instances creates and returns a new string by concatenating all of the elements in this array,
	// separated by commas or a specified separator string.
	Join(sep string) (s String, err Error)
}
