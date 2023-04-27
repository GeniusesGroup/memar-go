/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Object_Type uint8

const (
	Object_Type_Unset Object_Type = iota
	Object_Type_Type              // other type
	Object_Type_Pointer
	Object_Type_Boolean
	Object_Type_Function
	Object_Type_Structure
	Object_Type_Array
	// Object_Type_String UTF8, ...??
	Object_Type_Natural  // Any number > 0	- https://en.wikipedia.org/wiki/Natural_number
	Object_Type_Whole    // Any number >= 0	-
	Object_Type_Integer  // also know as signed number is any number <>= 0 - https://en.wikipedia.org/wiki/Integer
	Object_Type_Rational // also knows as decimal, float, ... - https://en.wikipedia.org/wiki/Rational_number
	Object_Type_Real     // also know as Irrational, Fractional - https://en.wikipedia.org/wiki/Real_number
	Object_Type_Complex  // - https://en.wikipedia.org/wiki/Complex_number
)
