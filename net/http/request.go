/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/net/uri"
	"memar/protocol"
)

// Request is represent HTTP request protocol structure.
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	method  string
	uri     uri.URI
	version string

	H header // Exported field to let consumers use other methods that protocol.HTTPHeader
	body
}

//memar:impl memar/protocol.ObjectLifeCycle
func (r *Request) Init() (err protocol.Error) {
	err = r.H.Init()
	if err != nil {
		return
	}
	err = r.body.Init()
	return
}
func (r *Request) Reinit() (err protocol.Error) {
	r.method = ""
	err = r.uri.Reinit()
	if err != nil {
		return
	}
	r.version = ""
	err = r.H.Reinit()
	if err != nil {
		return
	}
	err = r.body.Reinit()
	return
}
func (r *Request) Deinit() (err protocol.Error) {
	err = r.H.Deinit()
	if err != nil {
		return
	}
	err = r.body.Deinit()
	return
}

//memar:impl memar/protocol.HTTPRequest
func (r *Request) Method() string              { return r.method }
func (r *Request) URI() protocol.URI           { return &r.uri }
func (r *Request) Version() string             { return r.version }
func (r *Request) SetMethod(method string)     { r.method = method }
func (r *Request) SetVersion(version string)   { r.version = version }
func (r *Request) Header() protocol.HTTPHeader { return &r.H }

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
	if r.uri.Authority() == "" {
		r.uri.SetAuthority(r.H.Get(HeaderKeyHost))
	}
}
