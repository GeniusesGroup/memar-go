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

// Response is represent response protocol structure.
// https://tools.ietf.org/html/rfc2616#section-6
type Response struct {
	http.Response[ascii.STR[buffer_p.Buffer]]
	v1_header Header
	v1_body   incomeBody
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *Response) Init() (err error_p.Error) {
	err = r.Response.Init()
	if err != nil {
		return
	}
	err = r.v1_header.Init(&r.Response.Header)
	return
}
func (r *Response) Reinit() (err error_p.Error) {
	err = r.Response.Reinit()
	if err != nil {
		return
	}
	err = r.v1_header.Reinit(&r.Response.Header)
	return
}
func (r *Response) Deinit() (err error_p.Error) {
	err = r.Response.Deinit()
	if err != nil {
		return
	}
	err = r.v1_header.Deinit()
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (r *Response) SerializationLength() (ln container_p.NumberOfElement) {
	ln = r.LenWithoutBody()
	ln += r.Body().SerializationLength()
	return
}

//memar:impl memar/protocol.Decoder
func (r *Response) Decode(source buffer_p.Buffer) (err error_p.Error) {
	// TODO::: logic below not working when sender send header with large body!
	if source.UnreadLength() > http.Length_Max_Header {
		// err =
		return
	}

	// Make a buffer to hold incoming data.
	// TODO::: change to get from buffer pool??
	var buf = make([]byte, 0, http.Length_Max_Header)
	// Read the incoming connection into the buffer.
	buf, err = source.MarshalTo(buf)
	if err != nil {
		// err = connection.ErrNoConnection
		return
	}

	buf, err = r.UnmarshalFrom(buf)
	if err != nil {
		return
	}
	err = r.body.checkAndSetCodecAsIncomeBody(buf, source, &r.H)
	return
}

//memar:impl memar/protocol.Encoder
func (r *Response) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	var lenWithoutBody = r.LenWithoutBody()
	var bodyLen = r.body.SerializationLength()
	var wholeLen = lenWithoutBody + bodyLen
	// Check if whole request has fewer length than http.Length_Max_Header and Decide to send header and body separately
	if wholeLen > http.Length_Max_Header {
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

// Marshal encodes whole r *Response data and return httpPacket!
func (r *Response) Marshal() (httpPacket []byte, err error_p.Error) {
	httpPacket = make([]byte, 0, r.SerializationLength())
	httpPacket, err = r.MarshalTo(httpPacket)
	return
}

// MarshalTo encodes whole r *Response data to given httpPacket and return it by new len!
func (r *Response) MarshalTo(httpPacket []byte) (added []byte, err error_p.Error) {
	httpPacket = append(httpPacket, r.Version()...)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = append(httpPacket, r.StatusCode()...)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = append(httpPacket, r.ReasonPhrase()...)
	httpPacket = append(httpPacket, http.CRLF...)

	httpPacket = r.H.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, CRLF...)

	httpPacket, err = r.body.MarshalTo(httpPacket)
	added = httpPacket
	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Response.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Response) Unmarshal(httpPacket []byte) (n int, err error_p.Error) {
	var maybeBody []byte
	maybeBody, err = r.UnmarshalFrom(httpPacket)
	if err != nil {
		return
	}
	err = r.body.checkAndSetIncomeBody(maybeBody, &r.H)
	n = len(httpPacket)
	return
}

// UnmarshalFrom parses and decodes data of given httpPacket to r *Response.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (r *Response) UnmarshalFrom(httpPacket []byte) (maybeBody []byte, err error_p.Error) {
	// By use unsafe pointer here all strings assign in Response will just point to httpPacket slice
	// and no need to alloc lot of new memory locations and copy response line and headers keys & values!
	var s = convert.UnsafeByteSliceToString(httpPacket)
	var n int
	n, err = r.unmarshalFrom(s)
	maybeBody = httpPacket[n:]
	return
}

/*
********** buffer_p.Buffer interface **********
 */

// ReadFrom decodes r *Response data by read from given io.Reader!
// Declare to respect io.ReaderFrom interface!
func (r *Response) ReadFrom(reader io.Reader) (n int64, goErr error) {
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
	err = r.body.checkAndSetReaderAsIncomeBody(buf, codec.ReaderAdaptor{reader}, &r.H)
	n = int64(headerReadLength)
	return
}

// WriteTo encodes r *Response data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (r *Response) WriteTo(writer io.Writer) (n int64, err error) {
	var lenWithoutBody = r.LenWithoutBody()
	var bodyLen = r.body.SerializationLength()
	var wholeLen = lenWithoutBody + bodyLen
	// Check if whole request has fewer length than http.Length_Max_Header and Decide to send header and body separately
	if wholeLen > http.Length_Max_Header {
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

//memar:impl memar/codec/string/protocol.Stringer
func (r *Response) ToString() (str string, err error_p.Error) {
	var req, _ = r.Marshal()
	str = convert.UnsafeByteSliceToString(req)
	return
}
func (r *Response) FromString(str string) (err error_p.Error) {
	_, err = r.unmarshalFrom(str)
	return
}

func (r *Response) SetStatusByError(err error_p.Error) {
	r.SetStatus(StatusBadRequestCode, StatusBadRequestPhrase)
	// if err != nil {
	// 	const errorHeaders = "\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: close\r\n\r\n"

	// 	switch {
	// 	case err == errTooLarge:
	// 		// Their HTTP client may or may not be
	// 		// able to read this if we're
	// 		// responding to them and hanging up
	// 		// while they're still writing their
	// 		// request. Undefined behavior.
	// 		const publicErr = "431 Request Header Fields Too Large"
	// 		fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
	// 		c.closeWriteAndWait()
	// 		return

	// 	case isUnsupportedTEError(err):
	// 		// Respond as per RFC 7230 Section 3.3.1 which says,
	// 		//      A server that receives a request message with a
	// 		//      transfer coding it does not understand SHOULD
	// 		//      respond with 501 (Unimplemented).
	// 		code := StatusNotImplemented

	// 		// We purposefully aren't echoing back the transfer-encoding's value,
	// 		// so as to mitigate the risk of cross side scripting by an attacker.
	// 		fmt.Fprintf(c.rwc, "HTTP/1.1 %d %s%sUnsupported transfer encoding", code, StatusText(code), errorHeaders)
	// 		return

	// 	case isCommonNetReadError(err):
	// 		return // don't reply

	// 	default:
	// 		if v, ok := err.(statusError); ok {
	// 			fmt.Fprintf(c.rwc, "HTTP/1.1 %d %s: %s%s%d %s: %s", v.code, StatusText(v.code), v.text, errorHeaders, v.code, StatusText(v.code), v.text)
	// 			return
	// 		}
	// 		publicErr := "400 Bad Request"
	// 		fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
	// 		return
	// 	}
	// }
}

/*
********** local methods **********
 */

// Unmarshal parses and decodes data of given httpPacket to r *Request until body start.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit.
func (r *Response) unmarshalFrom(httpPacket string) (n int, err error_p.Error) {
	// n hold httpPacket index and i hold s index in new sliced state.
	var i int

	n, err = r.parseFirstLine(httpPacket)
	if err != nil {
		return
	}
	n += 2 // +2 due to have "\r\n"
	httpPacket = httpPacket[n:]

	i, err = r.H.unmarshal(httpPacket)
	n += i
	if err != nil {
		return
	}
	// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
	// So it can be occur panic if very simple request end without any CRLF
	n += 2 // +2 due to have "\r\n" after header end

	return
}

// Unmarshal parses and decodes data of given httpPacket to r *Request until body start.
// First line: HTTP/1.0 200 OK
func (r *Response) parseFirstLine(s string) (si int, err error_p.Error) {
	// si hold s index and i hold s index in new sliced state.
	var i int

	i, err = r.parseVersion(s)
	if err != nil {
		return
	}
	i++ // +1 due to have ' '
	si = i
	s = s[i:]

	i, err = r.parseStatusCode(s)
	if err != nil {
		return
	}
	i++ // +1 due to have ' '
	si += i
	s = s[i:]

	i, err = r.parseReasonPhrase(s)
	if err != nil {
		return
	}
	si += i
	return
}

func (r *Response) parseVersion(s string) (i int, err error_p.Error) {
	i = strings.IndexByte(s[:versionMaxLength], SP)
	if i == -1 {
		// err = &errs.ErrParseVersion
		return
	}
	r.version = s[:i]
	return
}

func (r *Response) parseStatusCode(s string) (i int, err error_p.Error) {
	// First line: GET /index.html HTTP/1.0
	i = strings.IndexByte(s[:statusCodeMaxLength], SP)
	if i == -1 {
		// err = &errs.ErrParseStatusCode
		return
	}
	r.statusCode = s[:i]
	return
}

func (r *Response) parseReasonPhrase(s string) (i int, err error_p.Error) {
	i = strings.IndexByte(s, '\r')
	if i == -1 {
		// err = &errs.ErrParseReasonPhrase
		return
	}
	r.SetReasonPhrase(s[:i])
	return
}

// MarshalWithoutBody encodes r *Response data and return httpPacket without body part.
func (r *Response) MarshalWithoutBody() (httpPacket []byte) {
	httpPacket = make([]byte, 0, r.LenWithoutBody())
	httpPacket = r.MarshalToWithoutBody(httpPacket)
	return
}

// MarshalToWithoutBody encodes r *Response data and return httpPacket without body part.
func (r *Response) MarshalToWithoutBody(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, r.Version()...)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = append(httpPacket, r.StatusCode()...)
	httpPacket = append(httpPacket, http.SP)
	httpPacket = append(httpPacket, r.ReasonPhrase()...)
	httpPacket = append(httpPacket, http.CRLF...)

	httpPacket = r.v1_header.MarshalTo(httpPacket)
	httpPacket = append(httpPacket, http.CRLF...)
	return httpPacket
}

// LenWithoutBody return length of response without body length!
func (r *Response) LenWithoutBody() (ln container_p.NumberOfElement) {
	ln = 6 // 6==1+1+2+2==len(SP)+len(SP)+len(CRLF)+len(CRLF)
	ln += len(r.Version())
	ln += len(r.statusCode)
	ln += len(r.reasonPhrase)
	ln += r.H.SerializationLength()
	return
}
