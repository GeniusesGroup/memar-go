/* For license and copyright information please see LEGAL file in repository */

package protocol

// https://en.wikipedia.org/wiki/Uniform_Resource_Name
type URN interface {
	URI
	// URI() string // e.g. "urn:isbn:0451450523"
	// Scheme() string // always return "urn"
	NID() string // NID is the namespace identifier
	NSS() string // NSS is the namespace-specific
}

// URN is the giti standard urn : https://github.com/GeniusesGroup/RFCs/blob/master/Giti.md#URN
// URN() example: urn:giti:{{domain-name}}:page:{{page-name}}
// NID() always return "giti"
// Other standards:
// https://datatracker.ietf.org/doc/html/rfc2609
type GitiURN interface {
	URN
	// NID() string // always return "giti"
	UUID() [32]byte
	ID() uint64
	IDasString() string // return base64 of ID.
	Domain() string
	Scope() string // service, page, data-structure, data, app-node, ...
	Name() string  // It must be unique in the domain scope e.g. "product" in "page" scope of the "domain"
}
