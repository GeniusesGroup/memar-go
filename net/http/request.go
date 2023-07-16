/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"io"
	"strings"

	"libgo/codec"
	"libgo/convert"
	"libgo/protocol"
	"libgo/uri"
)

// Request is represent HTTP request protocol structure.
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	method  string
	uri     uri.URI
	version string

	H header // Exported field to let consumers use other methods that protocol.HTTPHeader
	body
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (r *Request) Init() {
	r.H.Init()
	r.body.Init()
}
func (r *Request) Reinit() {
	r.method = ""
	r.uri.Reinit()
	r.version = ""
	r.H.Reinit()
	r.body.Reinit()
}
func (r *Request) Deinit() {
	r.H.Deinit()
	r.body.Deinit()
}

//libgo:impl libgo/protocol.HTTPRequest
func (r *Request) Method() string              { return r.method }
func (r *Request) URI() protocol.URI           { return &r.uri }
func (r *Request) Version() string             { return r.version }
func (r *Request) SetMethod(method string)     { r.method = method }
func (r *Request) SetVersion(version string)   { r.version = version }
func (r *Request) Header() protocol.HTTPHeader { return &r.H }

//libgo:impl libgo/protocol.Codec
func (r *Request) MediaType() protocol.MediaType       { return &MediaTypeRequest }
func (r *Request) CompressType() protocol.CompressType { return nil }
func (r *Request) Len() (ln int) {
	ln = r.LenWithoutBody()
	ln += r.body.Len()
	return
}
func (r *Request) Decode(source protocol.Codec) (n int, err protocol.Error) {
	if source.Len() > MaxHTTPHeaderSize {
		// err =
		return
	}

	// Make a buffer to hold incoming data.
	// TODO::: change to get from buffer pool?? force to be on the thread(goroutine) stack??
	var data = make([]byte, 0, MaxHTTPHeaderSize)
	data, err = source.MarshalTo(data)
	if err != nil {
		return
	}

	data, err = r.UnmarshalFrom(data)
	if err != nil {
		return
	}
	err = r.body.checkAndSetCodecAsIncomeBody(data, source, &r.H)
	return
}
func (r *Request) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	var lenWithoutBody = r.LenWithoutBody()
	var bodyLen = r.body.Len()
	var wholeLen = lenWithoutBody + bodyLen
	// Check if whole request has fewer length than MaxHTTPHeaderSize and Decide to send header and body separately
	if wholeLen > MaxHTTPHeaderSize {
		var withoutBody = make([]byte, 0, lenWithoutBody)
		withoutBody = r.MarshalToWithoutBody(withoutBody)

		n, err = destination.Unmarshal(withoutBody)
		if err == nil && r.body.Codec != nil {
			var bodyWrote int
			bodyWrote, err = destination.Encode(&r.body)
			n += bodyWrote
		}
	} else {
		var httpPacket = make([]byte, 0, wholeLen)
		httpPacket, err = r.MarshalTo(httpPacket)
		n, err = destination.Unmarshal(httpPacket)
	}
	return
}

// Marshal encodes whole r *Request data and return httpPacket.
func (r *Request) Marshal() (httpPacket []byte, err protocol.Error) {
	httpPacket = make([]byte, 0, r.Len())
	httpPacket, err = r.MarshalTo(httpPacket)
	return
}

// MarshalTo encodes whole r *Request data to given httpPacket and return it with new len.
func (r *Request) MarshalTo(httpPacket []byte) (added []byte, err protocol.Error) {
	httpPacket = append(httpPacket, r.method...)
	httpPacket = append(httpPacket, SP)
	httpPacket = r.uri.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, SP)
	httpPacket = append(httpPacket, r.version...)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket = r.H.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket, err = r.body.MarshalTo(httpPacket)
	added = httpPacket
	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Request.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit.
func (r *Request) Unmarshal(httpPacket []byte) (n int, err protocol.Error) {
	var maybeBody []byte
	maybeBody, err = r.UnmarshalFrom(httpPacket)
	if err != nil {
		return
	}
	err = r.body.checkAndSetIncomeBody(maybeBody, &r.H)
	n = len(httpPacket)
	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Request until body start.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit.
func (r *Request) UnmarshalFrom(httpPacket []byte) (maybeBody []byte, err protocol.Error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy request line and headers keys & values.
	var s = convert.UnsafeByteSliceToString(httpPacket)

	// si hold s index and i hold s index in new sliced state.
	var si, i int

	si, err = r.parseFirstLine(s)
	if err != nil {
		maybeBody = httpPacket[si:]
		return
	}
	si += 2 // +2 due to have "\r\n"
	s = s[si:]

	i, err = r.H.unmarshal(s)
	if err != nil {
		maybeBody = httpPacket[i:]
		return
	}
	si += i
	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body.
	// So it can be occur panic if very simple request end without any CRLF
	si += 2 // +2 due to have "\r\n" after header end

	r.checkHost()

	return httpPacket[si:], nil
}

// ReadFrom decodes r *Request data by read from given io.Reader
//
//libgo:impl go/io.ReaderFrom
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
	err = r.body.checkAndSetReaderAsIncomeBody(buf, codec.ReaderAdaptor{reader}, &r.H)
	n = int64(headerReadLength)
	return
}

// WriteTo encodes r(*Request) data and write it to given io.Writer
//
//libgo:impl go/io.WriterTo
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
			n, err = r.body.WriteTo(writer)
		}
		n += int64(headerWriteLength)
	} else {
		var httpPacket = make([]byte, 0, wholeLen)
		httpPacket, _ = r.MarshalTo(httpPacket)
		var packetWriteLength int
		packetWriteLength, err = writer.Write(httpPacket)
		n = int64(packetWriteLength)
	}
	return
}

/*
********** Other methods **********
 */

// Unmarshal parses and decodes data of given httpPacket to r *Request until body start.
// First line: GET /index.html HTTP/1.0
func (r *Request) parseFirstLine(s string) (si int, err protocol.Error) {
	// si hold s index and i hold s index in new sliced state.
	var i int

	i, err = r.parseMethod(s)
	if err != nil {
		return
	}
	i++ // +1 due to have ' '
	si = i
	s = s[i:]

	i, err = r.uri.UnmarshalFromString(s)
	if err != nil {
		return
	}
	i++ // +1 due to have ' '
	si += i
	s = s[i:]

	i, err = r.parseVersion(s)
	if err != nil {
		return
	}
	si += i
	return
}

func (r *Request) parseMethod(s string) (i int, err protocol.Error) {
	// First line: GET /index.html HTTP/1.0
	i = strings.IndexByte(s[:methodMaxLength], SP)
	if i == -1 {
		err = &ErrParseMethod
		return
	}
	r.method = s[:i]
	return
}

func (r *Request) parseVersion(s string) (i int, err protocol.Error) {
	// First line: GET /index.html HTTP/1.0
	i = strings.IndexByte(s[:versionMaxLength], '\r')
	if i == -1 {
		err = &ErrParseVersion
		return
	}
	r.version = s[:i]
	return
}

// MarshalWithoutBody encodes r *Request data and return httpPacket without body part!
func (r *Request) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())
	httpPacket = r.MarshalToWithoutBody(httpPacket)
	return
}

// MarshalWithoutBody encodes r *Request data and return httpPacket without body part!
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

/*
********** local Request methods **********
 */

// checkHost check host of request by RFC 7230, section 5.3 rules: Must treat
//
//	GET / HTTP/1.1
//	Host: geniuses.group
//
// and
//
//	GET https://geniuses.group/ HTTP/1.1
//	Host: apis.geniuses.group
//
// the same. In the second case, any Host line is ignored.
func (r *Request) checkHost() {
	if r.uri.Authority() == "" {
		r.uri.SetAuthority(r.H.Get(HeaderKeyHost))
	}
}
