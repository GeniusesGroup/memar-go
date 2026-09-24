/* For license and copyright information please see the LEGAL file in the code repository */

package mediatype_p

// Parsed MediaType or MimeType protocol is the shape of any coding media-type.
// It is a special way to naming a DataType. So all MediaType implementors MUST be a DataType too, But not reverse.
// Means not all DataType need to implements MediaType
// It also implement our RFC details on https://github.com/GeniusesGroup/memar/blob/main/media-type.md
// https://en.wikipedia.org/wiki/Media_type
type Parsed interface {
	Field_MediaType

	// Below names are case-insensitive.
	MainType() string     // must
	Tree() string         // if any
	SubType() string      // must
	Suffix() string       // if any
	Parameters() []string // if any
}
