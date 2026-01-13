/* For license and copyright information please see the LEGAL file in the code repository */

package mediatype_p

type Identifier /*[STR string_p.String]*/ interface {
	MediaTypeID_Base64() string // STR       // Base64 of ID
	// identifier_p.UUID_Hash // Hash of MediaType()
}
