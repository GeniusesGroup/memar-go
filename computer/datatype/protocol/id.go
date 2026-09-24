/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

type Field_ID interface {
	DataTypeID() ID
}

// ID use as a way to distinguish data-types as domains.
// Use first 64bit of `datatype_p.UUID`
type ID uint64

type Field_ID_Base64 interface {
	// DataTypeID_Base64 return Base64 of ID
	DataTypeID_Base64() string
}
