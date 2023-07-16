/* For license and copyright information please see LEGAL file in repository */

package http

import "strconv"

// ContentLength read all value about content in header
func (h *header) ContentLength() (l uint64) {
	var contentLength = h.Get(HeaderKeyContentLength)
	l, _ = strconv.ParseUint(contentLength, 10, 64)
	return
}

// SetContentLength set body length to header
func (h *header) SetContentLength(bodyLength uint64) {
	h.Set(HeaderKeyContentLength, strconv.FormatUint(bodyLength, 10))
}

// SetZeroContentLength set body length to header
func (h *header) SetZeroContentLength() {
	h.Set(HeaderKeyContentLength, "0")
}

// SetContentLength set body length to header.
func (r *Request) SetContentLength() {
	r.H.Set(HeaderKeyContentLength, strconv.FormatUint(uint64(r.body.Len()), 10))
}

// SetContentLength set body length to header.
func (r *Response) SetContentLength() {
	r.H.Set(HeaderKeyContentLength, strconv.FormatUint(uint64(r.body.Len()), 10))
}
