/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// MediaType or MimeType protocol is the shape of any coding media-type.
// It is a special way to naming a DataType. So all MediaType implementors MUST be a DataType too, But not reverse.
// Means not all DataType need to implements MediaType
// It also implement our RFC details on https://github.com/GeniusesGroup/memar/blob/main/media-type.md
// https://en.wikipedia.org/wiki/Media_type
type MediaType /*[STR String]*/ interface {
	// Never change MediaType due to it adds unnecessary complicated troubleshooting on SDK.
	MediaType() /*STR*/ string // must "maintype "/" [tree "."] subtype ["+" suffix]* [";" parameters]"
}
