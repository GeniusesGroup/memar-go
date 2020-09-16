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
	var r Response
	r.Header.init()
	return &r
}

// Marshal enecodes r *Response data and append to given httpPacket
func (r *Response) Marshal() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.Len())

	httpPacket = append(httpPacket, r.Version...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.StatusCode...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.ReasonPhrase...)
	httpPacket = append(httpPacket, CRLF...)
	httpPacket = r.Header.Marshal(httpPacket)
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

	index = strings.IndexByte(s, '\r')
	if index == -1 {
		return ErrParsedErrorOnReasonPhrase
	}
	r.ReasonPhrase = s[:index]
	s = s[index+2:] // +2 due to "\r\n"

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.Version) + len(r.StatusCode) + len(r.ReasonPhrase) + 4

	index += r.Header.UnMarshal(s)

	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
	// So it can be occur panic if very simple request end without any CRLF
	index += 2 // +2 due to have "\r\n" after header end

	r.Body = httpPacket[index:]
	return
}

// GetStatusCode get status code as uit16
func (r *Response) GetStatusCode() (uint16, error) {
	// TODO::: don't use strconv for such simple task
	var c, err = strconv.ParseUint(r.StatusCode, 10, 16)
	return uint16(c), err
}

// SetStatus set given status code and phrase
func (r *Response) SetStatus(code, phrase string) {
	r.StatusCode = code
	r.ReasonPhrase = phrase
}

// SetError set given error to body of response
func (r *Response) SetError(err error) {
	r.Body = []byte(err.Error())
}

// Len return length of response
func (r *Response) Len() (ln int) {
	ln += len(r.Version)
	ln += len(r.StatusCode)
	ln += len(r.ReasonPhrase)
	ln += r.Header.Len()
	ln += 6 // 6=1+1+2+2=len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.Body)

	return
}
