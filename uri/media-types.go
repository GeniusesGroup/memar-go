/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"github.com/GeniusesGroup/libgo/detail"
	"github.com/GeniusesGroup/libgo/mediatype"
	"github.com/GeniusesGroup/libgo/protocol"
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

//libgo:impl protocol.MediaType
func (m *mediaType) FileExtension() string           { return "uri" }
func (m *mediaType) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaType) IssueDate() protocol.Time            { return nil }
func (m *mediaType) ExpiryDate() protocol.Time           { return nil }
func (m *mediaType) ExpireInFavorOf() protocol.MediaType { return nil }
func (m *mediaType) Fields() []protocol.Field            { return nil }
