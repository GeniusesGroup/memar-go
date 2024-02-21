/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// DataType in computer science and computer programming, is a collection or grouping of data values,
// usually specified by a set of possible values, a set of allowed operations on these values,
// and/or a representation of these values as machine types even in compilers or runtime packages.
// It can use for any data like CLA flags or json fields or any other data structures
// https://en.wikipedia.org/wiki/Data_type
type DataType /*[T any]*/ interface {
	// DataType_DefaultValue[T]
	// DataType_Accessor[T]
	// DataType_AtomicAccessor[T]
	// DataType_Clone[T]
	// DataType_Copy[T]

	// DataType_Validation
	// DataType_Locker

	// TODO::: add more

	// existence length
	// DataType_ExpectedLength
	// Len

	DataType_Details
	Stringer // value stringer
}

type DataType_Details interface {
	Status() SoftwareStatus
	ReferenceURI() string
	IssueDate() Time
	ExpiryDate() Time
	ExpireInFavorOf() DataType

	Detail
}
