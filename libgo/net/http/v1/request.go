/* For license and copyright information please see the LEGAL file in the code repository */

package http_v1

import (
	"io"
	"strings"

	"memar/codec"
	"memar/codec/string/ascii"
	container_p "memar/computer/adt/container/protocol"
	"memar/computer/buffer/byteslice/convert"
	buffer_p "memar/computer/buffer/protocol"
	"memar/libgo/net/http"
	error_p "memar/process/error/protocol"
)

// Request is represent HTTP request protocol structure.
// https://tools.ietf.org/html/rfc2616#section-5
type Request struct {
	http.Request[ascii.STR[buffer_p.Buffer]]
	v1_header Header
	v1_body   incomeBody
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *Request) Init() (err error_p.Error) {
	err = r.Request.Init()
	if err != nil {
		return
	}
	err = r.v1_header.Init(&r.Request.Header)
	return
}
func (r *Request) Reinit() (err error_p.Error) {
	err = r.Request.Reinit()
	if err != nil {
		return
	}
	err = r.v1_header.Reinit(&r.Request.Header)
	return
}
func (r *Request) Deinit() (err error_p.Error) {
	err = r.Request.Deinit()
	if err != nil {
		return
	}
	err = r.v1_header.Deinit()
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (r *Request) SerializationLength() (ln container_p.NumberOfElement) {
	ln = r.LenWithoutBody()
	ln += r.Request.Body().SerializationLength()
	return
}

//memar:impl memar/protocol.Decoder
func (r *Request) Decode(source buffer_p.Buffer) (err error_p.Error) {
	if source.UnreadLength() > http.Length_Max_Header {
		// err =
		return
	}

	// Make a buffer to hold incoming data.
	// TODO::: change to get from buffer pool?? force to be on the thread(goroutine) stack??
	var data = make([]byte, 0, http.Length_Max_Header)
	data, err = source.MarshalTo(data)
	if !err.IsEmpty() {
		return
	}

	data, err = r.UnmarshalFrom(data)
	if err != nil {
		return
	}
	err = r.v1_body.checkAndSetCodecAsIncomeBody(data, source, &r.v1_header)
	return
}

//memar:impl memar/protocol.Encoder
func (r *Request) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	var lenWithoutBody = r.LenWithoutBody()
	var bodyLen = r.Body().SerializationLength()
	var wholeLen = lenWithoutBody + bodyLen
	// Check if whole request has fewer length than http.Length_Max_Header and Decide to send header and body separately
	if wholeLen > http.Length_Max_Header {
		var withoutBody = make([]byte, 0, lenWithoutBody)
		withoutBody = r.MarshalToWithoutBody(withoutBody)

		n, err = destination.Unmarshal(withoutBody)
		if err == nil && r.Body() != nil {
			var bodyWrote int
			bodyWrote, err = destination.Encode(r.Body())
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
func (r *Request) Marshal() (httpPacket []byte, err error_p.Error) {
	httpPacket = make([]byte, 0, r.SerializationLength())
	httpPacket, err = r.MarshalTo(httpPacket)
	return
}

// MarshalTo encodes whole r *Request data to given httpPacket and return it with new len.
func (r *Request) MarshalTo(httpPacket []byte) (added []byte, err error_p.Error) {
	httpPacket = append(httpPacket, r.Method()...)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = r.U.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = append(httpPacket, r.Version()...)
	httpPacket = append(httpPacket, http.CRLF...)

	httpPacket = r.v1_header.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, http.CRLF...)

	httpPacket, err = r.Body().MarshalTo(httpPacket)
	added = httpPacket
	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Request.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit.
func (r *Request) Unmarshal(httpPacket []byte) (n int, err error_p.Error) {
	var maybeBody []byte
	maybeBody, err = r.UnmarshalFrom(httpPacket)
	if err != nil {
		return
	}
	err = r.v1_body.checkAndSetIncomeBody(maybeBody, &r.v1_header)
	n = len(httpPacket)
	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Request until body start.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit.
func (r *Request) UnmarshalFrom(httpPacket []byte) (maybeBody []byte, err error_p.Error) {
	// By use unsafe pointer here all strings assign in Request will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy request line and headers keys & values.
	var s = convert.UnsafeByteSliceToString(httpPacket)
	var n int
	n, err = r.unmarshalFrom(s)
	maybeBody = httpPacket[n:]
	return
}

// ReadFrom decodes r *Request data by read from given io.Reader
//
//memar:impl std/io.ReaderFrom
func (r *Request) ReadFrom(reader io.Reader) (n int64, goErr error) {
	// Make a buffer to hold incoming data.
	var buf = make([]byte, http.Length_Max_Header)
	var headerReadLength int
	var err error_p.Error

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
	err = r.v1_body.checkAndSetReaderAsIncomeBody(buf, codec.ReaderAdaptor{reader}, &r.v1_header)
	n = int64(headerReadLength)
	return
}

// WriteTo encodes r(*Request) data and write it to given io.Writer
//
//memar:impl std/io.WriterTo
func (r *Request) WriteTo(writer io.Writer) (n int64, err error) {
	var lenWithoutBody = r.LenWithoutBody()
	var bodyLen = r.Body().SerializationLength()
	var wholeLen = lenWithoutBody + bodyLen
	// Check if whole request has fewer length than http.Length_Max_Header and Decide to send header and body separately
	if wholeLen > http.Length_Max_Header {
		var httpPacket = make([]byte, 0, lenWithoutBody)
		httpPacket = r.MarshalToWithoutBody(httpPacket)

		var headerWriteLength int
		headerWriteLength, err = writer.Write(httpPacket)
		if err == nil && r.Body() != nil {
			n, err = r.Body().WriteTo(writer)
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

//memar:impl memar/codec/string/protocol.Stringer
func (r *Request) ToString() (s string, err error_p.Error) {
	var req, _ = r.Marshal()
	s = convert.UnsafeByteSliceToString(req)
	return
}
func (r *Request) FromString(str string) (err error_p.Error) {
	_, err = r.unmarshalFrom(str)
	return
}

/*
********** Other methods **********
 */

// Unmarshal parses and decodes data of given httpPacket to r *Request until body start.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit.
func (r *Request) unmarshalFrom(httpPacket string) (n int, err error_p.Error) {
	// n hold httpPacket index and i hold s index in new sliced state.
	var i int

	n, err = r.parseFirstLine(httpPacket)
	if err != nil {
		return
	}
	n += 2 // +2 due to have "\r\n"
	httpPacket = httpPacket[n:]

	i, err = r.v1_header.unmarshal(httpPacket)
	n += i
	if err != nil {
		return
	}
	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body.
	// So it can be occur panic if very simple request end without any CRLF
	n += 2 // +2 due to have "\r\n" after header end

	r.CheckHost()

	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Request until body start.
// First line: GET /index.html HTTP/1.0
func (r *Request) parseFirstLine(s string) (si int, err error_p.Error) {
	// si hold s index and i hold s index in new sliced state.
	var i int

	i, err = r.parseMethod(s)
	if err != nil {
		return
	}
	i++ // +1 due to have ' '
	si = i
	s = s[i:]

	i, err = r.U.FromString(s)
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

func (r *Request) parseMethod(s string) (i int, err error_p.Error) {
	// First line: GET /index.html HTTP/1.0
	i = strings.IndexByte(s[:http.Length_Max_MethodMaxLength], http.SP)
	if i == -1 {
		// err = &errs.ErrParseMethod
		return
	}
	r.SetMethod(s[:i])
	return
}

func (r *Request) parseVersion(s string) (i int, err error_p.Error) {
	// First line: GET /index.html HTTP/1.0
	i = strings.IndexByte(s[:http.Length_Max_VersionMaxLength], '\r')
	if i == -1 {
		// err = &errs.ErrParseVersion
		return
	}
	r.SetVersion(s[:i])
	return
}

// MarshalWithoutBody encodes r *Request data and return httpPacket without body part.
func (r *Request) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())
	httpPacket = r.MarshalToWithoutBody(httpPacket)
	return
}

// MarshalWithoutBody encodes r *Request data and return httpPacket without body part.
func (r *Request) MarshalToWithoutBody(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, r.Method()...)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = r.U.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = append(httpPacket, r.Version()...)
	httpPacket = append(httpPacket, http.CRLF...)

	httpPacket = r.v1_header.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, http.CRLF...)
	return httpPacket
}

// LenWithoutBody return length of request without body length
func (r *Request) LenWithoutBody() (ln container_p.NumberOfElement) {
	ln = 6 // 6=1+1+2+2=len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += container_p.NumberOfElement(len(r.Method()))
	ln += r.U.SerializationLength()
	ln += container_p.NumberOfElement(len(r.Version()))
	ln += r.v1_header.SerializationLength()
	return
}
