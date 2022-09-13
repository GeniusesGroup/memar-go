/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"github.com/GeniusesGroup/libgo/detail"
	"github.com/GeniusesGroup/libgo/mediatype"
	"github.com/GeniusesGroup/libgo/protocol"
)

var (
	MediaType         mediaType
	MediaTypeRequest  mediaTypeRequest
	MediaTypeResponse mediaTypeResponse
)

func init() {
	MediaType.Init("application/http")
	MediaType.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Hypertext Transfer Protocol",
		"An application layer protocol in the Internet protocol suite model for distributed, collaborative, hypermedia information",
		"",
		"",
		[]string{})

	MediaTypeRequest.Init("application/http; request")
	MediaTypeRequest.SetDetail(protocol.LanguageEnglish, domainEnglish, "Hypertext Transfer Protocol Request", "", "", "", []string{})

	MediaTypeResponse.Init("application/http; response")
	MediaTypeResponse.SetDetail(protocol.LanguageEnglish, domainEnglish, "Hypertext Transfer Protocol Response", "", "", "", []string{})
}

type mediaType struct {
	detail.DS
	mediatype.MT
}

//libgo:impl protocol.MediaType
func (m *mediaType) FileExtension() string           { return "http" }
func (m *mediaType) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaType) IssueDate() protocol.Time            { return nil }
func (m *mediaType) ExpiryDate() protocol.Time           { return nil }
func (m *mediaType) ExpireInFavorOf() protocol.MediaType { return nil }
func (m *mediaType) Fields() []protocol.Field            { return nil }

type mediaTypeRequest struct {
	detail.DS
	mediatype.MT
}

//libgo:impl protocol.MediaType
func (m *mediaTypeRequest) FileExtension() string           { return "req.http" }
func (m *mediaTypeRequest) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaTypeRequest) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaTypeRequest) IssueDate() protocol.Time            { return nil }
func (m *mediaTypeRequest) ExpiryDate() protocol.Time           { return nil }
func (m *mediaTypeRequest) ExpireInFavorOf() protocol.MediaType { return nil }
func (m *mediaTypeRequest) Fields() []protocol.Field            { return nil }

type mediaTypeResponse struct {
	detail.DS
	mediatype.MT
}

//libgo:impl protocol.MediaType
func (m *mediaTypeResponse) FileExtension() string           { return "res.http" }
func (m *mediaTypeResponse) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaTypeResponse) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaTypeResponse) IssueDate() protocol.Time            { return nil }
func (m *mediaTypeResponse) ExpiryDate() protocol.Time           { return nil }
func (m *mediaTypeResponse) ExpireInFavorOf() protocol.MediaType { return nil }
func (m *mediaTypeResponse) Fields() []protocol.Field            { return nil }
