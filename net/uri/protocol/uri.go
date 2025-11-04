/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_URI interface {
	// URI Return full URI e.g. HTTP-URL
	URI() URI
}

// URI indicate "Uniform Resource Identifier".
// Although many URI schemes are named after protocols, this does not
// imply that use of these URIs will result in access to the resource
// via the named protocol.
// https://en.wikipedia.org/wiki/List_of_URI_schemes
//
// https://datatracker.ietf.org/doc/html/rfc3986#section-3:
// foo://example.com:8042/over/there?name=ferret#nose
// \_/   \______________/\_________/ \_________/ \__/
//
//	|           |            |            |        |
//
// scheme     authority     path        query   fragment
//
//	|   _____________________|__
//
// / \ /                        \
// urn:example:animal:ferret:nose
type URI interface {
	string_p.String
}
