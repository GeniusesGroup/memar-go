/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"libgo/detail"
	"libgo/mediatype"
	"libgo/protocol"
)

var (
	MediaType         mediaType
	MediaTypeRequest  mediaTypeRequest
	MediaTypeResponse mediaTypeResponse
)

func init() {
	MediaType.Init("application/http")
	MediaTypeRequest.Init("application/http; request")
	MediaTypeResponse.Init("application/http; response")
}

type mediaType struct {
	detail.DS
	mediatype.MT
}

//libgo:impl libgo/protocol.MediaType
func (m *mediaType) FileExtension() string           { return "http" }
func (m *mediaType) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaType) IssueDate() protocol.Time            { return nil }
func (m *mediaType) ExpiryDate() protocol.Time           { return nil }
func (m *mediaType) ExpireInFavorOf() protocol.MediaType { return nil }
func (m *mediaType) Members() []protocol.Object_Member   { return nil }

type mediaTypeRequest struct {
	detail.DS
	mediatype.MT
}

//libgo:impl /libgolibgo/protocol.MediaType
func (m *mediaTypeRequest) FileExtension() string           { return "req.http" }
func (m *mediaTypeRequest) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaTypeRequest) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaTypeRequest) IssueDate() protocol.Time            { return nil }
func (m *mediaTypeRequest) ExpiryDate() protocol.Time           { return nil }
func (m *mediaTypeRequest) ExpireInFavorOf() protocol.MediaType { return nil }
func (m *mediaTypeRequest) Members() []protocol.Object_Member   { return nil }

type mediaTypeResponse struct {
	detail.DS
	mediatype.MT
}

//libgo:impl libgo/protocol.MediaType
func (m *mediaTypeResponse) FileExtension() string           { return "res.http" }
func (m *mediaTypeResponse) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaTypeResponse) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaTypeResponse) IssueDate() protocol.Time            { return nil }
func (m *mediaTypeResponse) ExpiryDate() protocol.Time           { return nil }
func (m *mediaTypeResponse) ExpireInFavorOf() protocol.MediaType { return nil }
func (m *mediaTypeResponse) Members() []protocol.Object_Member   { return nil }
