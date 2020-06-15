/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"net/url"
	"strings"
	"unsafe"
)

// Request is represent HTTP request protocol structure!
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	Method  string
	URI     string
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
	httpPacket = append(httpPacket, r.URI...)
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
	r.URI = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, '\n')
	if index == -1 {
		return ErrParsedErrorOnVersion
	}
	r.Version = s[:index-1] // -1 due to have "\r\n"
	s = s[index+1:]

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.Method) + len(r.URI) + len(r.Method) + 4

	index += r.Header.UnMarshal(s)
	r.Body = httpPacket[index:]
	return
}

// GetURI returns decodes of r.URI
func (r *Request) GetURI() (u *url.URL, err error) {
	return url.Parse(r.URI)
}

// SetURI encodes given URI to request
func (r *Request) SetURI(u *url.URL) {
	r.Header.SetValue(HeaderKeyHost, u.Host)
	r.URI = u.Path + "?" + u.RawQuery
}
