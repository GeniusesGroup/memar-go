/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// DataType in computer science and computer programming, is a collection or grouping of data values,
// usually specified by a set of possible values, a set of allowed operations on these values,
// and/or a representation of these values as machine types even in compilers or runtime packages.
// It can use for any data like CLA flags or json fields or any other data structures
// https://en.wikipedia.org/wiki/Data_type
type DataType interface {
	DataType_ID
	DataType_Details
}

// DataTypeID use as a way to distinguish data-types as domains.
type DataTypeID uint64

type DataType_ID interface {
	DataTypeID() DataTypeID    // first 64bit of Hash of MediaType()
	DataTypeID_string() string // Base64 of ID
}

type DataType_Details interface {
	Status() SoftwareStatus
	ReferenceURI() string
	IssueDate() string  // Time
	ExpiryDate() string // Time
	ExpireInFavorOf() DataType

	Detail
}
