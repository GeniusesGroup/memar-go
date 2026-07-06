/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	uri_p "memar/net/uri/protocol"
)

// PseudoHeader_Request indicate request pseudo header.
// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
// STR MUST just include ASCII characters.
type PseudoHeader_Request /*[STR string_p.String]*/ interface {
	Field_Method

	// https://datatracker.ietf.org/doc/html/rfc2616#section-3.2
	// http_URL = "http:" "//" host [ ":" port ] [ abs_path [ "?" query ]]
	// URI() URI[String]
	uri_p.Field_Scheme // always return "http" or "https"
	uri_p.Field_Host
	uri_p.Field_Port
	uri_p.Field_Path
	uri_p.Field_Query

	Field_Version
}

// PseudoHeader_Response indicate response pseudo header.
// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
// STR MUST just include ASCII characters.
type PseudoHeader_Response /*[STR string_p.String]*/ interface {
	Field_Version
	Field_Status
	Method_Status
}
