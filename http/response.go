/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"strconv"
	"strings"
	"unsafe"
)

// Response is represent response protocol structure!
// https://tools.ietf.org/html/rfc2616#section-6
type Response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
	Header       header
	Body         []byte
}

// MakeNewResponse make new response with some default data
func MakeNewResponse() *Response {
	return &Response{
		Header: make(map[string][]string, 16),
	}
}

// Marshal enecodes r *Response data and append to given httpPacket
func (r *Response) Marshal() (httpPacket []byte, err error) {
	// Make packet by twice size of body
	httpPacket = make([]byte, 0, len(r.Body)*2)

	httpPacket = append(httpPacket, r.Version...)
	httpPacket = append(httpPacket, Space)
	httpPacket = append(httpPacket, r.StatusCode...)
	httpPacket = append(httpPacket, Space)
	httpPacket = append(httpPacket, r.ReasonPhrase...)
	httpPacket = append(httpPacket, CRLF...)
	r.Header.Marshal(&httpPacket)
	httpPacket = append(httpPacket, CRLF...)
	httpPacket = append(httpPacket, r.Body...)
	return
}

// UnMarshal parses and decodes data of given httpPacket to r *Response.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Response) UnMarshal(httpPacket []byte) (err error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy response line and headers keys & values!
	var s = *(*string)(unsafe.Pointer(&httpPacket))

	// First line: HTTP/1.0 200 OK
	var index int
	index = strings.IndexByte(s, ' ')
	if index == -1 {
		return ErrParsedErrorOnVersion
	}
	r.Version = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, ' ')
	if index == -1 {
		return ErrParsedErrorOnStatusCode
	}
	r.StatusCode = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, '\n')
	if index == -1 {
		return ErrParsedErrorOnReasonPhrase
	}
	r.ReasonPhrase = s[:index-1] // -1 due to "\r\n"
	s = s[index+1:]

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.Version) + len(r.StatusCode) + len(r.ReasonPhrase) + 4

	index += r.Header.UnMarshal(s)
	r.Body = httpPacket[index:]
	return
}

// GetStatusCode get status code as uit16
func (r *Response) GetStatusCode() (uint16, error) {
	// TODO::: don't use strconv for such simple task
	var c, err = strconv.ParseUint(r.StatusCode, 10, 16)
	return uint16(c), err
}

// SetStatusCode set given status uit16 code
func (r *Response) SetStatusCode(code uint16) {
	// TODO::: don't use strconv for such simple task
	r.StatusCode = strconv.FormatUint(uint64(code), 10)
}
