/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"
	"strconv"
	"strings"

	"../convert"
	"../mediatype"
	"../protocol"
)

// Response is represent response protocol structure!
// https://tools.ietf.org/html/rfc2616#section-6
type Response struct {
	version      string
	statusCode   string
	reasonPhrase string

	H header // Exported field to let consumers use other methods that protocol.HTTPHeader
	body
}

// NewResponse make new response with some default data and return it!
func NewResponse() *Response {
	var r Response
	r.H.init()
	return &r
}

func (r *Response) Version() string               { return r.version }
func (r *Response) StatusCode() string            { return r.statusCode }
func (r *Response) ReasonPhrase() string          { return r.reasonPhrase }
func (r *Response) SetVersion(version string)     { r.version = version }
func (r *Response) SetStatus(code, phrase string) { r.statusCode = code; r.reasonPhrase = phrase }
func (r *Response) Header() protocol.HTTPHeader   { return &r.H }

// GetStatusCode get status code as uit16
func (r *Response) GetStatusCode() (code uint16, err protocol.Error) {
	// TODO::: don't use strconv for such simple task
	var c, goErr = strconv.ParseUint(r.statusCode, 10, 16)
	if goErr != nil {
		return 0, ErrParseStatusCode
	}
	return uint16(c), nil
}

// GetError return realted er.Error in header of the Response
func (r *Response) GetError() (err protocol.Error) {
	var errIDString = r.H.Get(HeaderKeyErrorID)
	var errID, _ = strconv.ParseUint(errIDString, 10, 64)
	err = protocol.App.GetErrorByID(errID)
	return
}

// SetError set given er.Error to header of the response
func (r *Response) SetError(err protocol.Error) {
	r.H.Set(HeaderKeyErrorID, err.URN().IDasString())
}

// Redirect set given status and target location to the response
// httpRes.Redirect(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase, "http://www.google.com/")
func (r *Response) Redirect(code, phrase string, target string) {
	r.SetStatus(code, phrase)
	r.H.Set(HeaderKeyLocation, target)
}

/*
********** protocol.Codec interface **********
 */

func (r *Response) MediaType() protocol.MediaType       { return mediatype.HTTPResponse }
func (r *Response) CompressType() protocol.CompressType { return nil }
func (r *Response) Len() (ln int) {
	ln = r.LenWithoutBody()
	ln += r.body.Len()
	return
}

func (r *Response) Decode(reader protocol.Reader) (err protocol.Error) {
	// Make a buffer to hold incoming data.
	var buf = make([]byte, MaxHTTPHeaderSize)
	// Read the incoming connection into the buffer.
	var headerReadLength, goErr = reader.Read(buf)
	if goErr != nil || headerReadLength == 0 {
		// err = connection.ErrNoConnection
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

// Encode write response to given buf.
// Pass buf with enough cap. Make buf by r.Len() or grow it by buf.Grow(r.Len())
func (r *Response) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = r.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}

// Marshal enecodes whole r *Response data and return httpPacket!
func (r *Response) Marshal() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.Len())
	httpPacket = r.MarshalTo(httpPacket)
	return
}

// MarshalTo enecodes whole r *Response data to given httpPacket and return it by new len!
func (r *Response) MarshalTo(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.statusCode...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.reasonPhrase...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.H.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.body.MarshalTo(httpPacket)
	return httpPacket
}

// Unmarshal parses and decodes data of given httpPacket to r *Response.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Response) Unmarshal(httpPacket []byte) (err protocol.Error) {
	var maybeBody []byte
	maybeBody, err = r.UnmarshalFrom(httpPacket)
	if err != nil {
		return
	}
	err = r.body.checkAndSetIncomeBody(maybeBody, &r.H)
	return
}

// UnmarshalFrom parses and decodes data of given httpPacket to r *Response.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Response) UnmarshalFrom(httpPacket []byte) (maybeBody []byte, err protocol.Error) {
	// By use unsafe pointer here all strings assign in Response will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy response line and headers keys & values!
	var s = convert.UnsafeByteSliceToString(httpPacket)

	// First line: HTTP/1.0 200 OK
	var index int
	index = strings.IndexByte(s[:versionMaxLength], SP)
	if index == -1 {
		return httpPacket[:], ErrParseVersion
	}
	r.version = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s[:statusCodeMaxLength], SP)
	if index == -1 {
		return httpPacket[index:], ErrParseStatusCode
	}
	r.statusCode = s[:index]
	s = s[index+1:]

	index = strings.IndexByte(s, '\r')
	if index == -1 {
		return httpPacket[index:], ErrParseReasonPhrase
	}
	r.reasonPhrase = s[:index]
	s = s[index+2:] // +2 due to "\r\n"

	// TODO::: check performance below vs make new Int var for bodyStart and add to it in each IndexByte()
	// vs have 4 Int for each index
	index = len(r.version) + len(r.statusCode) + len(r.reasonPhrase) + 4

	index += r.H.Unmarshal(s)

	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
	// So it can be occur panic if very simple request end without any CRLF
	index += 2 // +2 due to have "\r\n" after header end
	return httpPacket[index:], nil
}

/*
********** io package interfaces **********
 */

// ReadFrom decodes r *Response data by read from given io.Reader!
// Declare to respect io.ReaderFrom interface!
func (r *Response) ReadFrom(reader io.Reader) (n int64, goErr error) {
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

// WriteTo enecodes r *Response data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (r *Response) WriteTo(writer io.Writer) (totalWrite int64, err error) {
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

		totalWrite = int64(bodyLen + headerWriteLength)
	} else {
		var httpPacket = make([]byte, 0, wholeLen)
		httpPacket = r.MarshalTo(httpPacket)
		var packetWriteLength int
		packetWriteLength, err = writer.Write(httpPacket)
		totalWrite = int64(packetWriteLength)
	}
	return
}

/*
********** local methods **********
 */

// MarshalWithoutBody enecodes r *Response data and return httpPacket without body part!
func (r *Response) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())
	httpPacket = r.MarshalToWithoutBody(httpPacket)
	return
}

// MarshalToWithoutBody enecodes r *Response data and return httpPacket without body part!
func (r *Response) MarshalToWithoutBody(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.statusCode...)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.reasonPhrase...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.H.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)
	return httpPacket
}

// LenWithoutBody return length of response without body length!
func (r *Response) LenWithoutBody() (ln int) {
	ln = 6 // 6==1+1+2+2==len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.version)
	ln += len(r.statusCode)
	ln += len(r.reasonPhrase)
	ln += r.H.Len()
	return
}
