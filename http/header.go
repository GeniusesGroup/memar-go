/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"strings"

	"golang.org/x/exp/maps"

	"libgo/convert"
	"libgo/protocol"
)

// header is represent HTTP header structure!
type header struct {
	headers    map[string][]string
	valuesPool []string // shared backing array for headers' values
}

func (h *header) Init() {
	h.headers = make(map[string][]string, headerInitMapLen)
	h.valuesPool = make([]string, headerValuesPoolLen)
}
func (h *header) Reinit() {
	maps.Clear(h.headers)
	h.valuesPool = h.valuesPool[:0]
}
func (h *header) Deinit() {}

// Get returns the first value associated with the given key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Get(key string) string {
	var values = h.headers[key]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

// Gets returns all values associated with the given key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Gets(key string) []string {
	var values = h.headers[key]
	if len(values) > 0 {
		return values
	}
	return nil
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Add(key, value string) {
	var values []string = h.headers[key]
	if values == nil {
		h.Set(key, value)
	} else {
		h.headers[key] = append(values, value)
	}
}

// Adds append given values to end of given key exiting values.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Adds(key string, values []string) {
	h.headers[key] = append(h.headers[key], values...)
}

// Set replace given value in given key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Set(key string, value string) {
	if len(h.valuesPool) == 0 {
		h.valuesPool = make([]string, headerValuesPoolLen)
	}
	// More than likely this will be a single-element key. Most headers aren't multi-valued.
	// Set the capacity on valuesPool[0] to 1, so any future append won't extend the slice into the other strings.
	var values []string = h.valuesPool[:1:1]
	h.valuesPool = h.valuesPool[1:]
	values[0] = value
	h.headers[key] = values
}

// Sets sets the header entries associated with key to
// the single element value. It replaces any existing values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Sets(key string, values []string) {
	h.headers[key] = values
}

// Replace change the exiting header entry associated with key to the given single element value.
// It use for some logic like TransferEncoding(), ...
func (h *header) Replace(key string, value string) {
	h.headers[key][0] = value
}

// Del deletes the values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Del(key string) {
	delete(h.headers, key)
}

// Exclude eliminate headers by given keys!
func (h *header) Exclude(exclude map[string]bool) {
	for key := range exclude {
		delete(h.headers, key)
	}
}

//libgo:impl libgo/protocol.Codec
func (h *header) Decode(source protocol.Codec) (n int, err protocol.Error) {
	// TODO:::
	return
}
func (h *header) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	var encodedHeader = h.Marshal()
	n, err = destination.Unmarshal(encodedHeader)
	return
}
func (h *header) Marshal() (httpHeader []byte) {
	httpHeader = make([]byte, 0, h.Len())
	httpHeader = h.MarshalTo(httpHeader)
	return
}
func (h *header) MarshalTo(httpPacket []byte) []byte {
	for key, values := range h.headers {
		// TODO::: some header key must not inline by coma like set-cookie. check if other need this exception.
		switch key {
		case HeaderKeySetCookie:
			for _, value := range values {
				httpPacket = append(httpPacket, key...)
				httpPacket = append(httpPacket, ColonSpace...)
				httpPacket = append(httpPacket, value...)
				httpPacket = append(httpPacket, CRLF...)
			}
		default:
			httpPacket = append(httpPacket, key...)
			httpPacket = append(httpPacket, ColonSpace...)
			for _, value := range values {
				httpPacket = append(httpPacket, value...)
				httpPacket = append(httpPacket, Comma)
			}
			httpPacket = httpPacket[:len(httpPacket)-1] // Remove trailing comma
			httpPacket = append(httpPacket, CRLF...)
		}
	}
	return httpPacket
}
func (h *header) Unmarshal(data []byte) (n int, err protocol.Error) {
	var s = convert.UnsafeByteSliceToString(data)
	n, err = h.unmarshal(s)
	return
}
func (h *header) Len() (ln int) {
	for key, values := range h.headers {
		ln += len(key)
		ln += 4 // 4=len(ColonSpace)+len(CRLF)
		for _, value := range values {
			ln += len(value)
			ln++ // 1=len(Coma)
		}
	}
	return
}

/*
********** Other methods **********
 */

// unmarshal parses and decodes data of given httpPacket(without first line) to (h *header).
// This method not respect to some RFCs like field-name in RFC7230, ... due to be more liberal in what it accept!
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (h *header) unmarshal(s string) (headerEnd int, err protocol.Error) {
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
			h.Add(key, value)                      // TODO::: is legal to have multiple key in request header or use h.Set()??
		}
		newLineIndex += 2 // +2 due to have "\r\n" at end of each header line
		s = s[newLineIndex:]
		headerEnd += newLineIndex
	}
}
