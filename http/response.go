/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"
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

// Marshal enecodes whole r *Response data and return httpPacket!
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

// MarshalWithoutBody enecodes r *Response data and return httpPacket without body part!
func (r *Response) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())

	httpPacket = append(httpPacket, r.Version...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.StatusCode...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.ReasonPhrase...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.Header.Marshal(httpPacket)
	httpPacket = append(httpPacket, CRLF...)
	return
}

// ReadFrom decodes r *Response data by read from given io.Reader!
// Declare to respect io.ReaderFrom interface!
func (r *Response) ReadFrom(reader io.Reader) (n int64, goErr error) {
	// Make a buffer to hold incoming data.
	var buf = make([]byte, MaxHTTPHeaderSize)
	var readLength, bodyReadLength int

	// Read the incoming connection into the buffer.
	readLength, goErr = reader.Read(buf)
	if goErr != nil || readLength == 0 {
		return
	}

	buf = buf[:readLength]
	goErr = r.UnMarshal(buf)
	if goErr != nil {
		return int64(readLength), goErr
	}

	var contentLength = r.Header.GetContentLength()
	// TODO::: is below logic check include all situations??
	if contentLength > 0 && len(r.Body) == 0 {
		r.Body = make([]byte, contentLength)
		bodyReadLength, goErr = reader.Read(r.Body)
		if goErr != nil {
			return int64(readLength + bodyReadLength), goErr
		}
		if bodyReadLength != int(contentLength) {
			// goErr =
			return int64(readLength + bodyReadLength), goErr
		}
	}

	return int64(readLength + bodyReadLength), goErr
}

// WriteTo enecodes r *Response data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (r *Response) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var resMarshaled = r.MarshalWithoutBody()
	var headerWriteLength, bodyWrittenLength int

	headerWriteLength, err = w.Write(resMarshaled)
	if err == nil {
		bodyWrittenLength, err = w.Write(r.Body)
	}

	totalWrite = int64(headerWriteLength + bodyWrittenLength)
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

// LenWithoutBody return length of response without body length!
func (r *Response) LenWithoutBody() (ln int) {
	ln = 6 // 6=1+1+2+2=len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.Version)
	ln += len(r.StatusCode)
	ln += len(r.ReasonPhrase)
	ln += r.Header.Len()
	return
}

// Len return length of response
func (r *Response) Len() (ln int) {
	ln = r.LenWithoutBody()
	ln += len(r.Body)
	return
}
