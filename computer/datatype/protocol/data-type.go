/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

type Field_DataType interface {
	DataType() DataType
}

type DataType interface {
	Identifiers
	Field_LifeCycle
	Details
}

type Details interface {
	ReferenceURI() string // uri
	IssueDate() string    // Time
	ExpiryDate() string   // Time
	ExpireInFavorOf() DataType

	Detail
}
