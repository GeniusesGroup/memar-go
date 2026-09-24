/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	http_headers "memar/libgo/net/http/headers"
	error_p "memar/process/error/protocol"
)

// Request is represent HTTP request protocol structure.
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	PseudoHeader_Request
	Header
	body

	datatype.DataType
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *Request) Init() (err error_p.Error) {
	err = r.PseudoHeader_Request.Init()
	if err != nil {
		return
	}
	err = r.Header.Init()
	if err != nil {
		return
	}
	err = r.body.Init()
	return
}
func (r *Request) Reinit() (err error_p.Error) {
	err = r.PseudoHeader_Request.Reinit()
	if err != nil {
		return
	}
	err = r.Header.Reinit()
	if err != nil {
		return
	}
	err = r.body.Reinit()
	return
}
func (r *Request) Deinit() (err error_p.Error) {
	err = r.PseudoHeader_Request.Deinit()
	if err != nil {
		return
	}
	err = r.Header.Deinit()
	if err != nil {
		return
	}
	err = r.body.Deinit()
	return
}

// CheckHost check host of request by RFC 7230, section 5.3 rules: Must treat
//
//	GET / HTTP/1.1
//	Host: geniuses.group
//
// and
//
//	GET https://geniuses.group/ HTTP/1.1
//	Host: apis.geniuses.group
//
// the same. In the second case, any Host line is ignored.
func (r *Request) CheckHost() {
	if r.URI.Authority() == "" {
		r.URI.SetAuthority(r.Header_Get(http_headers.Key_Host))
	}
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (r *Request) MediaType() string { return "application/http; request" }

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (r *Request) FileExtension() string { return "req.http" }

//memar:impl memar/datatype/protocol.DataType_Details
func (r *Request) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }
func (r *Request) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (r *Request) IssueDate() string                    { return "" }
func (r *Request) ExpiryDate() string                   { return "" }
func (r *Request) ExpireInFavorOf() datatype_p.DataType { return nil }
