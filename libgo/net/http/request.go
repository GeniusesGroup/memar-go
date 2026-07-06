/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	string_p "memar/codec/string/protocol"
)

// Request is represent HTTP request protocol structure.
// https://tools.ietf.org/html/rfc2616#section-5
type Request[STR string_p.String] struct {
	PseudoHeader_Request[STR]
	Header[STR]
	body

	datatype.DataType
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *Request[STR]) Init() (err error_p.Error) {
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
func (r *Request[STR]) Reinit() (err error_p.Error) {
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
func (r *Request[STR]) Deinit() (err error_p.Error) {
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
func (r *Request[STR]) CheckHost() {
	if r.URI.Authority() == "" {
		r.URI.SetAuthority(r.Header_Get(Key_Host))
	}
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (r *Request[STR]) MediaType() string { return "application/http; request" }

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (r *Request[STR]) FileExtension() string { return "req.http" }

//memar:impl memar/datatype/protocol.DataType_Details
func (r *Request[STR]) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }
func (r *Request[STR]) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (r *Request[STR]) IssueDate() string                    { return "" }
func (r *Request[STR]) ExpiryDate() string                   { return "" }
func (r *Request[STR]) ExpireInFavorOf() datatype_p.DataType { return nil }
