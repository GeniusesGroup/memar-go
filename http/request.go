/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"
	"strings"

	"../convert"
	"../mediatype"
	"../protocol"
)

// Request is represent HTTP request protocol structure!
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	method  string
	uri     URI
	version string

	H header // Exported field to let consumers use other methods that protocol.HTTPHeader
	body
}

// NewRequest make new request with some default data
func NewRequest() *Request {
	var r Request
	r.H.init()
	return &r
}

func (r *Request) Method() string              { return r.method }
func (r *Request) URI() protocol.HTTPURI       { return &r.uri }
func (r *Request) Version() string             { return r.version }
func (r *Request) SetMethod(method string)     { r.method = method }
func (r *Request) SetVersion(version string)   { r.version = version }
func (r *Request) Header() protocol.HTTPHeader { return &r.H }

/*
********** protocol.Codec interface **********
 */

func (r *Request) MediaType() protocol.MediaType       { return mediatype.HTTPRequest }
func (r *Request) CompressType() protocol.CompressType { return nil }
func (r *Request) Len() (ln int) {
	ln = r.LenWithoutBody()
	ln += r.body.Len()
	return
}

func (r *Request) Decode(reader protocol.Reader) (err protocol.Error) {
	// Make a buffer to hold incoming data.
	var buf = make([]byte, MaxHTTPHeaderSize)
	// Read the incoming connection into the buffer.
	var headerReadLength, goErr = reader.Read(buf)
	if goErr != nil || headerReadLength == 0 {
		// err =
		return
	}

	buf = buf[:headerReadLength]
	buf, err = r.UnmarshalFrom(buf)
	if err != nil {
		return err
	}
	r.body.checkAndSetReaderAsIncomeBody(buf, reader, &r.H)
	return
}

// Encode write request to given buffer writer.
func (r *Request) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = r.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}

// Marshal enecodes whole r *Request data and return httpPacket.
func (r *Request) Marshal() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.Len())
	httpPacket = r.MarshalTo(httpPacket)
	return
}

// MarshalTo enecodes whole r *Request data to given httpPacket and return it with new len.
func (r *Request) MarshalTo(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, r.method...)
	httpPacket = append(httpPacket, SP)
	httpPacket = r.uri.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.H.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.body.MarshalTo(httpPacket)
	return httpPacket
}

// Unmarshal parses and decodes data of given httpPacket to r *Request.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Request) Unmarshal(httpPacket []byte) (err protocol.Error) {
	var maybeBody []byte
	maybeBody, err = r.UnmarshalFrom(httpPacket)
	if err != nil {
		return
	}
	err = r.body.checkAndSetIncomeBody(maybeBody, &r.H)
	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Request.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Request) UnmarshalFrom(httpPacket []byte) (maybeBody []byte, err protocol.Error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy request line and headers keys & values!
	var s = convert.UnsafeByteSliceToString(httpPacket)

	// First line: GET /index.html HTTP/1.0
	var index int
	index = strings.IndexByte(s[:methodMaxLength], SP)
	if index == -1 {
		return httpPacket[:], ErrParseMethod
	}
	r.method = s[:index]
	s = s[index+1:]

	index = r.uri.unmarshalFrom(s)
	s = s[index+1:]

	index = strings.IndexByte(s[:versionMaxLength], '\r')
	if index == -1 {
		return httpPacket[index:], ErrParseVersion
	}
	r.version = s[:index]
	s = s[index+2:] // +2 due to have "\r\n"

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.method) + len(r.uri.uri) + len(r.version) + 4

	index += r.H.Unmarshal(s)

	r.uri.checkHost(&r.H)

	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
	// So it can be occur panic if very simple request end without any CRLF
	index += 2 // +2 due to have "\r\n" after header end
	return httpPacket[index:], nil
}

/*
********** io package interfaces **********
 */

// ReadFrom decodes r *Request data by read from given io.Reader!
// Declare to respect io.ReaderFrom interface!
func (r *Request) ReadFrom(reader io.Reader) (n int64, goErr error) {
	// Make a buffer to hold incoming data.
	var buf = make([]byte, MaxHTTPHeaderSize)
	var headerReadLength int
	var err protocol.Error

	// Read the incoming connection into the buffer.
	headerReadLength, goErr = reader.Read(buf)
	if goErr != nil || headerReadLength == 0 {
		return
	}

	buf = buf[:headerReadLength]
	buf, err = r.UnmarshalFrom(buf)
	if err != nil {
		return int64(headerReadLength), err
	}
	r.body.checkAndSetReaderAsIncomeBody(buf, reader, &r.H)

	n = int64(headerReadLength)
	return
}

// WriteTo enecodes r *Request data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (r *Request) WriteTo(writer io.Writer) (n int64, err error) {
	var lenWithoutBody = r.LenWithoutBody()
	var bodyLen = r.body.Len()
	var wholeLen = lenWithoutBody + bodyLen
	// Check if whole request has fewer length than MaxHTTPHeaderSize and Decide to send header and body separately
	if wholeLen > MaxHTTPHeaderSize {
		var httpPacket = make([]byte, 0, lenWithoutBody)
		httpPacket = r.MarshalToWithoutBody(httpPacket)

		var headerWriteLength int
		headerWriteLength, err = writer.Write(httpPacket)
		if err == nil && r.body.Codec != nil {
			err = r.body.Encode(writer)
		}

		n = int64(bodyLen + headerWriteLength)
	} else {
		var httpPacket = make([]byte, 0, wholeLen)
		httpPacket = r.MarshalTo(httpPacket)
		var packetWriteLength int
		packetWriteLength, err = writer.Write(httpPacket)
		n = int64(packetWriteLength)
	}
	return
}

/*
********** Other methods **********
 */

// MarshalWithoutBody enecodes r *Request data and return httpPacket without body part!
func (r *Request) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())
	httpPacket = r.MarshalToWithoutBody(httpPacket)
	return
}

// MarshalWithoutBody enecodes r *Request data and return httpPacket without body part!
func (r *Request) MarshalToWithoutBody(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, r.method...)
	httpPacket = append(httpPacket, SP)
	httpPacket = r.uri.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.H.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)
	return httpPacket
}

// LenWithoutBody return length of request without body length
func (r *Request) LenWithoutBody() (ln int) {
	ln = 6 // 6=1+1+2+2=len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.method)
	ln += r.uri.Len()
	ln += len(r.version)
	ln += r.H.Len()
	return
}
