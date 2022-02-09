/* For license and copyright information please see LEGAL file in repository */

package protocol

// https://en.wikipedia.org/wiki/List_of_URI_schemes
type URI interface {
	Init(uri string)
	URI() string
	Scheme() string

	// Codec
}

// https://en.wikipedia.org/wiki/Uniform_Resource_Name
type URN interface {
	URI
	// URI() string // e.g. "urn:isbn:0451450523"
	// Scheme() string // always return "urn"
	NID() string // NID is the namespace identifier
	NSS() string // NSS is the namespace-specific
}
