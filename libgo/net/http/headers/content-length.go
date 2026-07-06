/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"strconv"
)

// ContentLength read all value about content in header
func (h *Header) ContentLength() (l uint64) {
	var contentLength = h.Header_Get(HeaderKey_ContentLength)
	l, _ = strconv.ParseUint(contentLength, 10, 64)
	return
}

// AddContentLength add body length to header
func (h *Header) AddContentLength(bodyLength uint64) {
	h.Header_Add(HeaderKey_ContentLength, strconv.FormatUint(bodyLength, 10))
}

// AddZeroContentLength set body length to header
func (h *Header) AddZeroContentLength() {
	h.Header_Add(HeaderKey_ContentLength, "0")
}

// AddContentLength add body length to header.
func (r *Request) AddContentLength() {
	r.Header_Add(HeaderKey_ContentLength, strconv.FormatUint(uint64(r.body.Len()), 10))
}

// AddContentLength add body length to header.
func (r *Response) AddContentLength() {
	r.Header_Add(HeaderKey_ContentLength, strconv.FormatUint(uint64(r.body.Len()), 10))
}
