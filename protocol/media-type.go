/* For license and copyright information please see LEGAL file in repository */

package protocol

type MediaTypes interface {
	RegisterMediaType(mt MediaType)
	GetMediaType(mt string) MediaType
	GetMediaTypeByID(id uint64) MediaType
	GetMediaTypeByFileExtension(ex string) MediaType
}

// MediaType is standard shape of any coding media-type
// MediaType or MimeType standrad list can found here:
// http://www.iana.org/assignments/media-types/media-types.xhtml
// https://en.wikipedia.org/wiki/Media_type
// https://tools.ietf.org/html/rfc6838
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
type MediaType interface {
	ID() uint64        // first 64bit of Hash of MediaType()
	MediaType() string // must
	Type() string      // must
	Tree() string      // if any
	SubType() string   // must
	Suffix() string    // if any
	Parameter() string // if any
	FileExtension() string
	Status() SoftwareStatus
	IssueDate() Time
	ExpiryDate() Time
	ExpireInFavorOf() MediaType
}
