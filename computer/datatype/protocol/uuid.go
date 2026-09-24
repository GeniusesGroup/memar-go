/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

type Field_UUID interface {
	DataTypeUUID() UUID
}

// UUID use as a way to distinguish data-types as domains.
// It MUST fill by any UUID mechanism that guaranty it will be globally unique.
// TODO::: Hash of MediaType()??
// identifier_p.UUID_Hash
type UUID [32]byte

type Field_UUID_Base64 interface {
	// DataTypeUUID_Base64 return Base64 of UUID
	DataTypeUUID_Base64() string
}
