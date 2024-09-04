/* For license and copyright information please see the LEGAL file in the code repository */

package cl_p

type DataType_Primitive interface {
	Primitive() DataType_PrimitiveKind // DataType
}

// Primitive data types or built-in data types
// https://en.wikipedia.org/wiki/Primitive_data_type
type DataType_PrimitiveKind uint8

const (
	DataType_PrimitiveKind_Unset DataType_PrimitiveKind = iota

	// https://en.wikipedia.org/wiki/Function_type
	DataType_PrimitiveKind_Function
	// https://en.wikipedia.org/wiki/Method_(computer_programming)
	DataType_PrimitiveKind_Method

	DataType_PrimitiveKind_Object // or Structure that can be indicate by Type

	DataType_PrimitiveKind_Type // other type

	DataType_PrimitiveKind_Boolean
	DataType_PrimitiveKind_Array
	// DataType_PrimitiveKind_String UTF8, ...??

	// https://en.wikipedia.org/wiki/Natural_number
	DataType_PrimitiveKind_Natural // Any number > 0	-

	DataType_PrimitiveKind_Whole // Any number >= 0	-

	// https://en.wikipedia.org/wiki/Integer
	DataType_PrimitiveKind_Integer // also know as signed number is any number <>= 0

	// https://en.wikipedia.org/wiki/Rational_number
	DataType_PrimitiveKind_Rational // also knows as decimal, float, ...

	// https://en.wikipedia.org/wiki/Real_number
	DataType_PrimitiveKind_Real // also know as Irrational, Fractional

	// https://en.wikipedia.org/wiki/Complex_number
	DataType_PrimitiveKind_Complex //
)
