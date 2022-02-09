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
// It must also implement our RFC details on https://github.com/GeniusesGroup/RFCs/blob/master/media-type.md
type MediaType interface {
	UUID() [32]byte     // Hash of MediaType()
	ID() uint64         // first 64bit of UUID
	IDasString() string // Base64 of ID

	MediaType() string    // must
	Type() string         // must
	Tree() string         // if any
	SubType() string      // must
	Suffix() string       // if any
	Parameters() []string // if any
	FileExtension() string

	Status() SoftwareStatus
	IssueDate() TimeUnixSec  // TODO::: Temporary use TimeUnixSec instead of Time
	ExpiryDate() TimeUnixSec // TODO::: Temporary use TimeUnixSec instead of Time
	ExpireInFavorOf() MediaType
	Details() []MediaTypeDetail
	Detail(lang LanguageID) MediaTypeDetail
}

type MediaTypeDetail interface {
	Language() LanguageID
	// Domain return locale domain name that MediaType belongs to it.
	// More userfriendly domian name to show to users on screens.
	Domain() string
	// Summary return locale general summary MediaType text that gives the main points in a concise form
	Summary() string
	// Overview return locale general MediaType text that gives the main ideas without explaining all the details
	Overview() string
	// UserNote return locale note that user do when face this MediaType
	// Description text that gives the main ideas with explaining all the details and purposes.
	UserNote() string
	// DevNote return locale technical advice for developers
	// Description text that gives the main ideas with explaining all the details and purposes.
	DevNote() string
	// TAGS return locale MediaType tags to sort MediaType in groups for any purpose e.g. in GUI to help org manager to give service delegate authorization to staffs.
	TAGS() []string
}
