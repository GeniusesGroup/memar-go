/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

type Field_ID interface {
	DataTypeID() ID
}

// ID use as a way to distinguish data-types as domains.
// It MUST fill by first 64bit of Hash of MediaType()
type ID uint64
