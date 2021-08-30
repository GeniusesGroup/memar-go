/* For license and copyright information please see LEGAL file in repository */

package protocol

// https://en.wikipedia.org/wiki/List_of_URI_schemes
type URI interface {
	Init(uri string)
	URI() string
	Scheme() string

	// Codec
}
