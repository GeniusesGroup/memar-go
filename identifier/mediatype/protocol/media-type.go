/* For license and copyright information please see the LEGAL file in the code repository */

package mediatype_p

type Field_MediaType interface {
	MediaType() MediaType
}

// Never change MediaType due to it adds unnecessary complicated troubleshooting on SDK.
// must return >> "maintype "/" [tree "."] subtype ["+" suffix]* [";" parameters]"
type MediaType = string
