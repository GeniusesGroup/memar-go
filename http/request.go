/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"
	"strings"

	"../convert"
	er "../error"
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

// Marshal enecodes whole r *Request data and return httpPacket!
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
func (r *Request) UnMarshal(httpPacket []byte) (err *er.Error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy request line and headers keys & values!
	var s = convert.UnsafeByteSliceToString(httpPacket)

	// First line: GET /index.html HTTP/1.0
	var index int
	index = strings.IndexByte(s, ' ')
	if index == -1 {
		return ErrParseMethod
	}
	r.Method = s[:index]
	s = s[index+1:]

	index = r.URI.UnMarshal(s)
	s = s[index+1:]

	index = strings.IndexByte(s, '\r')
	if index == -1 {
		return ErrParseVersion
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

// MarshalWithoutBody enecodes r *Request data and return httpPacket without body part!
func (r *Request) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())

	httpPacket = append(httpPacket, r.Method...)
	httpPacket = append(httpPacket, SP)
	httpPacket = r.URI.Marshal(httpPacket)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.Version...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.Header.Marshal(httpPacket)
	httpPacket = append(httpPacket, CRLF...)
	return
}

// ReadFrom decodes r *Request data by read from given io.Reader!
// Declare to respect io.ReaderFrom interface!
func (r *Request) ReadFrom(reader io.Reader) (n int64, goErr error) {
	// Make a buffer to hold incoming data.
	var buf = make([]byte, MaxHTTPHeaderSize)
	var headerReadLength, bodyReadLength int
	var err *er.Error

	// Read the incoming connection into the buffer.
	headerReadLength, goErr = reader.Read(buf)
	if goErr != nil || headerReadLength == 0 {
		return
	}

	buf = buf[:headerReadLength]
	err = r.UnMarshal(buf)
	if err != nil {
		return int64(headerReadLength), err
	}

	var contentLength = r.Header.GetContentLength()
	// TODO::: is below logic check include all situations??
	if contentLength > 0 && len(r.Body) == 0 {
		r.Body = make([]byte, contentLength)
		bodyReadLength, goErr = reader.Read(r.Body)
		if bodyReadLength != int(contentLength) {
			// goErr =
		}
	}

	return int64(headerReadLength + bodyReadLength), goErr
}

// WriteTo enecodes r *Request data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (r *Request) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var reqMarshaled = r.MarshalWithoutBody()
	var headerWriteLength, bodyWrittenLength int

	headerWriteLength, err = w.Write(reqMarshaled)
	if err == nil {
		bodyWrittenLength, err = w.Write(r.Body)
	}

	totalWrite = int64(headerWriteLength + bodyWrittenLength)
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
	if r.URI.Authority == "" {
		return r.Header.Get(HeaderKeyHost)
	}
	return r.URI.Authority
}

// LenWithoutBody return length of request without body length
func (r *Request) LenWithoutBody() (ln int) {
	ln = 6 // 6=1+1+2+2=len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.Method)
	ln += r.URI.Len()
	ln += len(r.Version)
	ln += r.Header.Len()
	return
}

// Len return length of request
func (r *Request) Len() (ln int) {
	ln = r.LenWithoutBody()
	ln += len(r.Body)
	return
}
