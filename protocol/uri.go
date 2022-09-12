/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// URI indicate "Uniform Resource Identifier".
// Although many URI schemes are named after protocols, this does not
// imply that use of these URIs will result in access to the resource
// via the named protocol.
// https://en.wikipedia.org/wiki/List_of_URI_schemes
//
// https://datatracker.ietf.org/doc/html/rfc3986#section-3:
// foo://example.com:8042/over/there?name=ferret#nose
// \_/   \______________/\_________/ \_________/ \__/
//  |           |            |            |        |
// scheme     authority     path        query   fragment
//  |   _____________________|__
// / \ /                        \
// urn:example:animal:ferret:nose
type URI interface {
	Init(uri string)
	Set(scheme, authority, path, query, fragment string)

	URI() string // always return full URI e.g. HTTP-URL

	Scheme() string
	URI_Authority
	Path() string
	Query() string
	Fragment() string

	// Codec
}

type URI_Authority interface {
	Authority() string // [ userinfo "@" ] host [ ":" port ]

	URI_Userinfo
	Host() string
	Port() string
}

type URI_Userinfo interface {
	Userinfo() string // "username[:password]"

	Username() string
	Password() string
}

// URL indicate "Uniform Resource Locators".
// https://datatracker.ietf.org/doc/html/rfc1738
type URL = URI

// https://en.wikipedia.org/wiki/Uniform_Resource_Name
type URN interface {
	URI
	// URI() string    // e.g. "urn:isbn:0451450523"
	// Scheme() string // always return "urn"

	NID() string // NID is the namespace identifier e.g. "isbn"
	NSS() string // NSS is the namespace-specific   e.g. "0451450523"
}
