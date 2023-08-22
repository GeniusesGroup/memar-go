/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"libgo/detail"
	"libgo/mediatype"
	"libgo/protocol"
)

var (
	MediaType mediaType
)

func init() {
	MediaType.Init("application/uri")
	MediaType.SetDetail(protocol.LanguageEnglish, "URI", "", "", "", "", []string{})
}

type mediaType struct {
	detail.DS
	mediatype.MT
}

//libgo:impl libgo/protocol.MediaType
func (m *mediaType) FileExtension() string           { return "uri" }
func (m *mediaType) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaType) IssueDate() protocol.Time            { return nil }
func (m *mediaType) ExpiryDate() protocol.Time           { return nil }
func (m *mediaType) ExpireInFavorOf() protocol.MediaType { return nil }

//libgo:impl libgo/protocol.Object
func (m *mediaType) Fields() []protocol.Object_Member_Field   { return nil }
func (m *mediaType) Methods() []protocol.Object_Member_Method { return nil }
