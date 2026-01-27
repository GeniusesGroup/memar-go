/* For license and copyright information please see the LEGAL file in the code repository */

package adt_p

// Abstract data types are theoretical entities, used (among other things) to simplify the description of abstract algorithms,
// to classify and evaluate data structures, and to formally describe the type systems of programming languages.
// Primitive data types are a form of Abstract data type. It is just that they are provided by the language makers.
// https://en.wikipedia.org/wiki/Abstract_data_type
// It is a special DataType. So all ADT implementors MUST be a DataType too, But not reverse.
// Means not all DataType need to implements ADT
type ADT interface {
	Empty
	Nil
	Null
}
