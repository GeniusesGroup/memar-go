/* For license and copyright information please see LEGAL file in repository */

package http

import "strconv"

// GetContentLength read all value about content in header
func (h *header) GetContentLength() (l uint64) {
	var contentLength = h.Get(HeaderKeyContentLength)
	l, _ = strconv.ParseUint(contentLength, 10, 64)
	return
}

// SetContentLength set body length to header.
func (r *Request) SetContentLength() {
	r.Header.Set(HeaderKeyContentLength, strconv.FormatUint(uint64(len(r.Body)), 10))
}
