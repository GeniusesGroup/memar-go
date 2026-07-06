/* For license and copyright information please see the LEGAL file in the code repository */

package headers

import (
	"strconv"

	container_p "memar/computer/adt/container/protocol"
)

// ContentLength read all value about content in header
func (h *Header[STR]) ContentLength() (ne container_p.NumberOfElement) {
	var contentLength = h.Header_Get(Key_ContentLength)
	var l, _ = strconv.ParseUint(contentLength, 10, 64)
	ne = container_p.NumberOfElement(l)
	return
}

// AddContentLength add body length to header
func (h *Header[STR]) AddContentLength(bodyLength container_p.NumberOfElement) {
	h.Header_Add(Key_ContentLength, strconv.FormatInt(int64(bodyLength), 10))
}

// AddZeroContentLength set body length to header
func (h *Header[STR]) AddZeroContentLength() {
	h.Header_Add(Key_ContentLength, "0")
}

// AddContentLength add body length to header.
func (r *Request[STR]) AddContentLength() {
	if r.Body() != nil {
		var bodyLength = r.Body().SerializationLength()
		if bodyLength == 0 {
			r.Header.AddZeroContentLength()
		} else if bodyLength > 0 {
			r.Header.AddContentLength(r.body.SerializationLength())
		} else if bodyLength < 0 {
			r.AddTransferEncoding(HeaderValue_Chunked)
		}
	} else {
		r.AddZeroContentLength()
	}
}

// AddContentLength add body length to header.
func (r *Response[STR]) AddContentLength() {
	r.Header.AddContentLength(r.body.SerializationLength())
}
