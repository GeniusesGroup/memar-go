/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Object_Member can use for any data like CLA flags or json fields or any other data structures
// even in compilers or runtime packages
type Object_Member_Field /*[T any]*/ interface {
	// Get() T
	// Set(value T)
	// Default() T

	Type() Object_Type
	Optional() bool

	// TODO::: add more

	Validate() Error

	SetDefault() // default value

	Object_Member
	Object_Member_Len
	Details
	Stringer
}

type Object_Member_Len interface {
	// Expected length
	MinLen() int // in byte or 8bit
	MaxLen() int // in byte or 8bit

	// existence length
	Len
}

type Object_Type uint8

const (
	Object_Type_Unset Object_Type = iota
	// Object_Type_Object             // or Structure that can be indicate by Type
	Object_Type_Type // other type
	Object_Type_Pointer
	Object_Type_Boolean
	Object_Type_Array
	// Object_Type_String UTF8, ...??
	Object_Type_Natural  // Any number > 0	- https://en.wikipedia.org/wiki/Natural_number
	Object_Type_Whole    // Any number >= 0	-
	Object_Type_Integer  // also know as signed number is any number <>= 0 - https://en.wikipedia.org/wiki/Integer
	Object_Type_Rational // also knows as decimal, float, ... - https://en.wikipedia.org/wiki/Rational_number
	Object_Type_Real     // also know as Irrational, Fractional - https://en.wikipedia.org/wiki/Real_number
	Object_Type_Complex  // - https://en.wikipedia.org/wiki/Complex_number
)
