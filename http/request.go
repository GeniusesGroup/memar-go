/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"strings"
	"unsafe"
)

// Request is represent HTTP request protocol structure!
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	Method  string
	URI     URI
	Version string
	Header  header
	Body    []byte
}

// MakeNewRequest make new request with some default data
func MakeNewRequest() *Request {
	var r Request
	r.Header.init()
	return &r
}

// Marshal enecodes r *Request data and append to given httpPacket
func (r *Request) Marshal() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.Len())

	httpPacket = append(httpPacket, r.Method...)
	httpPacket = append(httpPacket, SP)

	httpPacket = r.URI.Marshal(httpPacket)
	httpPacket = append(httpPacket, SP)

	httpPacket = append(httpPacket, r.Version...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.Header.Marshal(httpPacket)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = append(httpPacket, r.Body...)
	return
}

// UnMarshal parses and decodes data of given httpPacket to r *Request.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Request) UnMarshal(httpPacket []byte) (err error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy request line and headers keys & values!
	var s = *(*string)(unsafe.Pointer(&httpPacket))

	// First line: GET /index.html HTTP/1.0
	var index int
	index = strings.IndexByte(s, ' ')
	if index == -1 {
		return ErrParsedErrorOnMethod
	}
	r.Method = s[:index]
	s = s[index+1:]

	index = r.URI.UnMarshal(s)
	s = s[index+1:]

	index = strings.IndexByte(s, '\r')
	if index == -1 {
		return ErrParsedErrorOnVersion
	}
	r.Version = s[:index]
	s = s[index+2:] // +2 due to have "\r\n"

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.Method) + len(r.URI.Raw) + len(r.Version) + 4

	index += r.Header.UnMarshal(s)

	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
	// So it can be occur panic if very simple request end without any CRLF
	index += 2 // +2 due to have "\r\n" after header end

	r.Body = httpPacket[index:]
	return
}

// GetHost returns host of request by RFC 7230, section 5.3 rules: Must treat
//		GET / HTTP/1.1
//		Host: www.sabz.city
// and
//		GET https://www.sabz.city/ HTTP/1.1
//		Host: apis.sabz.city
// the same. In the second case, any Host line is ignored.
func (r *Request) GetHost() (host string) {
	if len(r.URI.Authority) == 0 {
		return r.Header.Get(HeaderKeyHost)
	}
	return r.URI.Authority
}

// Len return length of request
func (r *Request) Len() (ln int) {
	ln += len(r.Method)
	ln += r.URI.Len()
	ln += len(r.Version)
	ln += r.Header.Len()
	ln += 6 // 6=1+1+2+2=len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.Body)

	return
}
