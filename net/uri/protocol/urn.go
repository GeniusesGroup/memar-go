/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

// URN is a division of URI
// e.g. "urn:isbn:0451450523"
// https://en.wikipedia.org/wiki/Uniform_Resource_Name
type URN interface {
	Field_URI

	// always return "urn"
	Field_Scheme

	Field_URN_NID
	Field_URN_NSS
}
