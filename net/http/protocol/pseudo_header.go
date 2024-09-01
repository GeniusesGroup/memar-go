/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	uri_p "memar/net/uri/protocol"
	string_p "memar/string/protocol"
)

// PseudoHeader_Request indicate request pseudo header.
// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
// STR MUST just include ASCII characters.
type PseudoHeader_Request[STR string_p.String] interface {
	Method[STR]

	// https://datatracker.ietf.org/doc/html/rfc2616#section-3.2
	// http_URL = "http:" "//" host [ ":" port ] [ abs_path [ "?" query ]]
	// URI() URI[String]
	uri_p.Scheme[STR] // always return "http" or "https"
	uri_p.Host[STR]
	uri_p.Port[STR]
	uri_p.Path[STR]
	uri_p.Query[STR]

	Version[STR]
}

// PseudoHeader_Response indicate response pseudo header.
// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
// STR MUST just include ASCII characters.
type PseudoHeader_Response[STR string_p.String] interface {
	Version[STR]
	Status[STR]
}
