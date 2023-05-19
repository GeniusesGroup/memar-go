/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type MediaTypes interface {
	RegisterMediaType(mt MediaType)
	GetMediaType(mt string) MediaType
	GetMediaTypeByID(id MediaTypeID) MediaType
	GetMediaTypeByFileExtension(ex string) MediaType
}

type MediaTypeID = ID

// MediaType or MimeType protocol is the shape of any coding media-type.
// It also implement our RFC details on https://github.com/GeniusesGroup/RFCs/blob/master/media-type.md
type MediaType interface {
	// Below names are case-insensitive.
	MediaType() string    // must "maintype "/" [tree "."] subtype ["+" suffix]* [";" parameters]"
	MainType() string     // must
	Tree() string         // if any
	SubType() string      // must
	Suffix() string       // if any
	Parameters() []string // if any

	FileExtension() string // if any

	Status() SoftwareStatus
	ReferenceURI() string
	IssueDate() Time
	ExpiryDate() Time
	ExpireInFavorOf() MediaType

	Object    // In explicit mediatype like domain maintype not like "application/json"
	UUID_Hash // Hash of MediaType()
	Details
	Stringer
}
