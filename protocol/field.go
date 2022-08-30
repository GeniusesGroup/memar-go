/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Structure interface {
	Fields() []Field
}

// Field can use for any data like CLA flags or json fields or any other data structures
// even in compilers or runtime packages
type Field interface {
	Name() string
	Abbreviation() string
	Type() FieldType
	Size() int // len in byte
	Optional() bool
	Immutable() bool
	Atomic() bool
	// TODO::: add more
	Validate() Error

	SetDefault() // default value

	Details
	Stringer
}

type FieldType uint8

const (
	FieldType_Unset FieldType = iota
	FieldType_Type            // other type
	FieldType_Pointer
	FieldType_Boolean
	FieldType_Function
	FieldType_Structure
	FieldType_Array
	FieldType_Natural  // Any number > 0	- https://en.wikipedia.org/wiki/Natural_number
	FieldType_Whole    // Any number >= 0	-
	FieldType_Integer  // also know as signed number is any number <>= 0 - https://en.wikipedia.org/wiki/Integer
	FieldType_Rational // also knows as decimal, float, ... - https://en.wikipedia.org/wiki/Rational_number
	FieldType_Real     // also know as Irrational, Fractional - https://en.wikipedia.org/wiki/Real_number
	FieldType_Complex  // - https://en.wikipedia.org/wiki/Complex_number
)
