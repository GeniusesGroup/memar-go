/* For license and copyright information please see LEGAL file in repository */

package protocol

// MediaType is standard shape of any coding media-type
// MediaType or MimeType standrad list can found here:
// http://www.iana.org/assignments/media-types/media-types.xhtml
// https://en.wikipedia.org/wiki/Media_type
// https://tools.ietf.org/html/rfc6838
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
type MediaType interface {
	URN() GitiURN
	MediaType() string
	MainType() string
	SubType() string
	FileExtension() string
}
