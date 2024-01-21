/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/protocol"
)

// Request is represent HTTP request protocol structure.
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	PseudoHeader_Request
	Header
	body
}

//memar:impl memar/protocol.ObjectLifeCycle
func (r *Request) Init() (err protocol.Error) {
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
func (r *Request) Reinit() (err protocol.Error) {
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
func (r *Request) Deinit() (err protocol.Error) {
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
	if r.U.Authority() == "" {
		r.U.SetAuthority(r.Header_Get(HeaderKey_Host))
	}
}
