/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/string/protocol"
)

// https://en.wikipedia.org/wiki/Uniform_Resource_Name
type URN[STR string_p.String] interface {
	URI[STR]    // e.g. "urn:isbn:0451450523"
	Scheme[STR] // always return "urn"

	URN_NID[STR]
	URN_NSS[STR]
}

type URN_NID[STR string_p.String] interface {
	NID() STR // NID is the namespace identifier e.g. "isbn"
}
type URN_NSS[STR string_p.String] interface {
	NSS() STR // NSS is the namespace-specific   e.g. "0451450523"
}
