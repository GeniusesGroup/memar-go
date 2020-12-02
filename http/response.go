/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"strconv"
	"strings"

	"../convert"
	er "../error"
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
func (r *Response) UnMarshal(httpPacket []byte) (err *er.Error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy response line and headers keys & values!
	var s = convert.UnsafeByteSliceToString(httpPacket)

	// First line: HTTP/1.0 200 OK
	var index int
	index = strings.IndexByte(s, ' ')
	if index == -1 {
		return ErrParseVersion
	}
	r.Version = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, ' ')
	if index == -1 {
		return ErrParseStatusCode
	}
	r.StatusCode = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, '\r')
	if index == -1 {
		return ErrParseReasonPhrase
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
func (r *Response) GetStatusCode() (code uint16, err *er.Error) {
	// TODO::: don't use strconv for such simple task
	var c, goErr = strconv.ParseUint(r.StatusCode, 10, 16)
	if goErr != nil {
		return 0, ErrParseStatusCode
	}
	return uint16(c), nil
}

// SetStatus set given status code and phrase
func (r *Response) SetStatus(code, phrase string) {
	r.StatusCode = code
	r.ReasonPhrase = phrase
}

// SetError set given er.Error to header of the response
func (r *Response) SetError(err *er.Error) {
	r.Header.Set(HeaderKeyErrorID, err.IDasString())
}

// SetContentLength set body length to header.
func (r *Response) SetContentLength() {
	r.Header.Set(HeaderKeyContentLength, strconv.FormatUint(uint64(len(r.Body)), 10))
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
