/* For license and copyright information please see the LEGAL file in the code repository */

package http_v1

import (
	"strings"

	container_p "memar/computer/adt/container/protocol"
	"memar/computer/buffer/byteslice/convert"
	buffer_p "memar/computer/buffer/protocol"
	"memar/libgo/net/http"
	error_p "memar/process/error/protocol"
)

// Header is represent HTTP header structure.
// Exported to let consumers use other methods that memar/libgo/net/http/protocol.Header
type Header struct {
	*http.Header
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (h *Header) Init(header *http.Header) (err error_p.Error) {
	h.Header = header
	return
}
func (h *Header) Reinit(header *http.Header) (err error_p.Error) {
	h.Header = header
	return
}
func (h *Header) Deinit() (err error_p.Error) {
	return
}

//memar:impl memar/codec/protocol.Codec
func (h *Header) Decode(source buffer_p.Buffer) (err error_p.Error) {
	// TODO:::
	return
}
func (h *Header) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	var encodedHeader = h.Marshal()
	err = destination.Unmarshal(encodedHeader)
	return
}
func (h *Header) Marshal() (httpHeader []byte) {
	httpHeader = make([]byte, 0, h.SerializationLength())
	httpHeader = h.MarshalTo(httpHeader)
	return
}
func (h *Header) MarshalTo(httpPacket []byte) []byte {
	var headers = h.Header_All()
	for _, kv := range headers {
		httpPacket = append(httpPacket, kv.Key()...)
		httpPacket = append(httpPacket, http.ColonSpace...)
		httpPacket = append(httpPacket, kv.Value()...)
		httpPacket = append(httpPacket, http.CRLF...)
	}
	return httpPacket
}
func (h *Header) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	var s = convert.UnsafeByteSliceToString(data)
	n, err = h.unmarshal(s)
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (h *Header) SerializationLength() (ln container_p.NumberOfElement) {
	for _, kv := range h.Header_All() {
		ln += container_p.NumberOfElement(len(kv.Key()))
		ln += 4 // 4=len(ColonSpace)+len(CRLF)
		ln += container_p.NumberOfElement(len(kv.Value()))
	}
	return
}

// unmarshal parses and decodes data of given httpPacket(without first line) to (h *Header).
// This method not respect to some RFCs like field-name in RFC7230, ... due to be more liberal in what it accept.
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit.
func (h *Header) unmarshal(s string) (headerEnd int, err error_p.Error) {
	var colonIndex, newLineIndex int
	var key, value string
	for {
		newLineIndex = strings.IndexByte(s, '\r')
		if newLineIndex < 3 {
			// newLineIndex == -1 >> broken or malformed packet, panic may occur!
			// newLineIndex == 0 >> End of headers part of packet, no panic
			// 1 < newLineIndex > 3 >> bad header || broken || malformed packet, panic may occur!
			return
		}

		colonIndex = strings.IndexByte(s[:newLineIndex], ':')
		switch colonIndex {
		case -1:
			// TODO::: Header key without value!?? Bad http packet!??
		default:
			key = s[:colonIndex]
			value = s[colonIndex+2 : newLineIndex] // +2 due to have a space after colon force by RFC &&
			h.Header_Add(key, value)               // TODO::: is legal to have multiple key in request header or use h.Set()??
		}
		newLineIndex += 2 // +2 due to have "\r\n" at end of each header line
		s = s[newLineIndex:]
		headerEnd += newLineIndex
	}
}
