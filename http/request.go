/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"
	"strings"

	"../convert"
	"../giti"
)

// Request is represent HTTP request protocol structure!
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	method  string
	uri     URI
	version string

	header header
	body
}

// NewRequest make new request with some default data
func NewRequest() *Request {
	var r Request
	r.header.init()
	return &r
}

func (r *Request) Method() string            { return r.method }
func (r *Request) URI() giti.HTTPURI         { return &r.uri }
func (r *Request) Version() string           { return r.version }
func (r *Request) SetMethod(method string)   { r.method = method }
func (r *Request) SetVersion(version string) { r.version = version }
func (r *Request) Header() giti.HTTPHeader   { return &r.header }

/*
********** giti.Codec interface **********
 */

func (r *Request) Decode(buf giti.Buffer) (err giti.Error) {
	var httpPacket = buf.GetUnread()
	err = r.UnMarshal(httpPacket)
	return
}

// Encode write request to given buf.
// Pass buf with enough cap. Make buf by r.Len() or grow it by buf.Grow(r.Len())
func (r *Request) Encode(buf giti.Buffer) {
	buf.WriteString(r.method)
	buf.WriteByte(SP)
	r.uri.Encode(buf)
	buf.WriteByte(SP)
	buf.WriteString(r.version)
	buf.WriteString(CRLF)

	r.header.Encode(buf)
	buf.WriteString(CRLF)

	r.body.Encode(buf)
}

// Marshal enecodes whole r *Request data and return httpPacket!
func (r *Request) Marshal() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.Len())
	httpPacket = r.MarshalTo(httpPacket)
	return
}

// MarshalTo enecodes whole r *Request data to given httpPacket and return it with new len!
func (r *Request) MarshalTo(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, r.method...)
	httpPacket = append(httpPacket, SP)
	httpPacket = r.uri.Marshal(httpPacket)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.header.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.body.MarshalTo(httpPacket)
	return httpPacket
}

// UnMarshal parses and decodes data of given httpPacket to r *Request.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Request) UnMarshal(httpPacket []byte) (err giti.Error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy request line and headers keys & values!
	var s = convert.UnsafeByteSliceToString(httpPacket)

	// First line: GET /index.html HTTP/1.0
	var index int
	index = strings.IndexByte(s, SP)
	if index == -1 {
		return ErrParseMethod
	}
	r.method = s[:index]
	s = s[index+1:]

	index = r.uri.UnMarshal(s)
	s = s[index+1:]

	index = strings.IndexByte(s, '\r')
	if index == -1 {
		return ErrParseVersion
	}
	r.version = s[:index]
	s = s[index+2:] // +2 due to have "\r\n"

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.method) + len(r.uri.raw) + len(r.version) + 4

	index += r.header.UnMarshal(s)

	r.uri.checkHost(&r.header)

	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
	// So it can be occur panic if very simple request end without any CRLF
	index += 2 // +2 due to have "\r\n" after header end

	r.body.checkEncodingAndSetBody(httpPacket[index:], &r.header)
	return
}

// MarshalWithoutBody enecodes r *Request data and return httpPacket without body part!
func (r *Request) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())

	httpPacket = append(httpPacket, r.method...)
	httpPacket = append(httpPacket, SP)
	httpPacket = r.uri.Marshal(httpPacket)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.header.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)
	return
}

// ReadFrom decodes r *Request data by read from given io.Reader!
// Declare to respect io.ReaderFrom interface!
func (r *Request) ReadFrom(reader io.Reader) (n int64, goErr error) {
	// Make a buffer to hold incoming data.
	var buf = make([]byte, MaxHTTPHeaderSize)
	var headerReadLength, bodyReadLength int
	var err giti.Error

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

	var contentLength = r.header.GetContentLength()
	// TODO::: is below logic check include all situations??
	if contentLength > 0 && r.body.Len() == 0 {
		var bodyRaw = make([]byte, contentLength)
		bodyReadLength, goErr = reader.Read(bodyRaw)
		if bodyReadLength != int(contentLength) {
			// goErr =
		}
		r.body.checkEncodingAndSetBody(bodyRaw, &r.header)
	}

	return int64(headerReadLength + bodyReadLength), goErr
}

// WriteTo enecodes r *Request data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (r *Request) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var reqMarshaled = r.MarshalWithoutBody()
	var headerWriteLength int

	headerWriteLength, err = w.Write(reqMarshaled)
	if err == nil {
		totalWrite, err = r.body.WriteTo(w)
	}

	totalWrite += int64(headerWriteLength)
	return
}

// LenWithoutBody return length of request without body length
func (r *Request) LenWithoutBody() (ln int) {
	ln = 6 // 6=1+1+2+2=len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.method)
	ln += r.uri.Len()
	ln += len(r.version)
	ln += r.header.Len()
	return
}

// Len return length of request
func (r *Request) Len() (ln int) {
	ln = r.LenWithoutBody()
	if r.body.Codec != nil {
		ln += r.body.Len()
	}
	return
}
