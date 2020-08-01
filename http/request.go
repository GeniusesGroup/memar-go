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
func MakeNewRequest() (r *Request) {
	return &Request{
		Header: make(map[string][]string, 16),
	}
}

// Marshal enecodes r *Request data and append to given httpPacket
func (r *Request) Marshal() (httpPacket []byte) {
	// Make packet by twice size of body
	httpPacket = make([]byte, 0, len(r.Body)*2)

	httpPacket = append(httpPacket, r.Method...)
	httpPacket = append(httpPacket, Space)

	r.Header.SetValue(HeaderKeyHost, r.URI.Authority)
	httpPacket = append(httpPacket, r.URI.MarshalRequestURI()...)
	httpPacket = append(httpPacket, Space)

	httpPacket = append(httpPacket, r.Version...)
	httpPacket = append(httpPacket, CRLF...)

	r.Header.Marshal(&httpPacket)
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

	index = strings.IndexByte(s, ' ')
	if index == -1 {
		return ErrParsedErrorOnURI
	}
	r.URI.UnMarshal(s[:index])
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
		return r.Header.GetValue(HeaderKeyHost)
	}
	return r.URI.Authority
}
