/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"
	"strconv"
	"strings"

	"../buffer"
	"../convert"
	er "../error"
	"../giti"
)

// Response is represent response protocol structure!
// https://tools.ietf.org/html/rfc2616#section-6
type Response struct {
	version      string
	statusCode   string
	reasonPhrase string

	header header

	body      []byte     // only for read from peer!
	bodyCodec giti.Codec // only for send to peer!
}

// NewResponse make new response with some default data and return it!
func NewResponse() *Response {
	var r Response
	r.header.init()
	return &r
}

func (r *Response) Version() string               { return r.version }
func (r *Response) SetVersion(version string)     { r.version = version }
func (r *Response) StatusCode() string            { return r.statusCode }
func (r *Response) ReasonPhrase() string          { return r.reasonPhrase }
func (r *Response) SetStatus(code, phrase string) { r.statusCode = code; r.reasonPhrase = phrase }
func (r *Response) Header() giti.HTTPHeader       { return &r.header }
func (r *Response) Body() []byte                  { return r.body }
func (r *Response) BodyCodec() giti.Codec         { return r.bodyCodec }
func (r *Response) SetBodyCodec(codec giti.Codec) { r.bodyCodec = codec }

// GetStatusCode get status code as uit16
func (r *Response) GetStatusCode() (code uint16, err giti.Error) {
	// TODO::: don't use strconv for such simple task
	var c, goErr = strconv.ParseUint(r.statusCode, 10, 16)
	if goErr != nil {
		return 0, ErrParseStatusCode
	}
	return uint16(c), nil
}

// GetError return realted er.Error in header of the Response
func (r *Response) GetError() (err giti.Error) {
	var errIDString = r.header.Get(HeaderKeyErrorID)
	var errID, _ = strconv.ParseUint(errIDString, 10, 64)
	err = er.Errors.GetErrorByID(errID)
	return
}

// SetError set given er.Error to header of the response
func (r *Response) SetError(err giti.Error) {
	r.header.Set(HeaderKeyErrorID, err.IDasString())
}

func (r *Response) Decode(buf giti.Buffer) (err giti.Error) {
	var httpPacket = buf.GetUnread()
	err = r.UnMarshal(httpPacket)
	return
}

// Encode write response to given buf.
// Pass buf with enough cap. Make buf by r.Len() or grow it by buf.Grow(r.Len())
func (r *Response) Encode(buf giti.Buffer) {
	buf.WriteString(r.version)
	buf.WriteByte(SP)
	buf.WriteString(r.statusCode)
	buf.WriteByte(SP)
	buf.WriteString(r.reasonPhrase)
	buf.WriteString(CRLF)

	r.header.Encode(buf)
	buf.WriteString(CRLF)

	r.bodyCodec.Encode(buf)
}

// Marshal enecodes whole r *Response data and return httpPacket!
func (r *Response) Marshal() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.Len())

	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.statusCode...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.reasonPhrase...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.header.Marshal(httpPacket)
	httpPacket = append(httpPacket, CRLF...)

	var buf = buffer.New(httpPacket)
	r.bodyCodec.Encode(buf)
	return
}

// UnMarshal parses and decodes data of given httpPacket to r *Response.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Response) UnMarshal(httpPacket []byte) (err giti.Error) {
	// By use unsafe pointer here all strings assign in Response will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy response line and headers keys & values!
	var s = convert.UnsafeByteSliceToString(httpPacket)

	// First line: HTTP/1.0 200 OK
	var index int
	index = strings.IndexByte(s, SP)
	if index == -1 {
		return ErrParseVersion
	}
	r.version = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, SP)
	if index == -1 {
		return ErrParseStatusCode
	}
	r.statusCode = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, '\r')
	if index == -1 {
		return ErrParseReasonPhrase
	}
	r.reasonPhrase = s[:index]
	s = s[index+2:] // +2 due to "\r\n"

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.version) + len(r.statusCode) + len(r.reasonPhrase) + 4

	index += r.header.UnMarshal(s)

	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
	// So it can be occur panic if very simple request end without any CRLF
	index += 2 // +2 due to have "\r\n" after header end

	r.body = httpPacket[index:]
	return
}

// MarshalWithoutBody enecodes r *Response data and return httpPacket without body part!
func (r *Response) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())

	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.statusCode...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.reasonPhrase...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.header.Marshal(httpPacket)
	httpPacket = append(httpPacket, CRLF...)
	return
}

// ReadFrom decodes r *Response data by read from given io.Reader!
// Declare to respect io.ReaderFrom interface!
func (r *Response) ReadFrom(reader io.Reader) (n int64, goErr error) {
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
	if contentLength > 0 && len(r.body) == 0 {
		r.body = make([]byte, contentLength)
		bodyReadLength, goErr = reader.Read(r.body)
		if bodyReadLength != int(contentLength) {
			// goErr =
		}
	}

	return int64(headerReadLength + bodyReadLength), goErr
}

// WriteTo enecodes r *Response data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (r *Response) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var resMarshaled = r.MarshalWithoutBody()
	var headerWriteLength, bodyWrittenLength int

	headerWriteLength, err = w.Write(resMarshaled)
	if err == nil {
		bodyWrittenLength, err = w.Write(r.body)
	}

	totalWrite = int64(headerWriteLength + bodyWrittenLength)
	return
}

// LenWithoutBody return length of response without body length!
func (r *Response) LenWithoutBody() (ln int) {
	ln = 6 // 6==1+1+2+2==len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.version)
	ln += len(r.statusCode)
	ln += len(r.reasonPhrase)
	ln += r.header.Len()
	return
}

// Len return length of response
func (r *Response) Len() (ln int) {
	ln = r.LenWithoutBody()
	ln += len(r.body)
	if r.bodyCodec != nil {
		ln += r.bodyCodec.Len()
	}
	return
}
