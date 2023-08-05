/* For license and copyright information please see the LEGAL file in the code repository */

package http

import "strings"

type ContentEncodings string

func (c ContentEncodings) Next() (contentEncoding string, remain ContentEncodings) {
	var commaIndex int = strings.IndexByte(string(c), Comma)
	if commaIndex == -1 {
		commaIndex = len(c)
	} else {
		remain = c[commaIndex+1:]
	}
	contentEncoding = string(c[:commaIndex])
	return
}

// ContentEncoding return content encoding and notify if multiple exist
// To read multiple just call this method in a loop to get multiple became false
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
func (h *header) ContentEncoding() (ce string, multiple bool) {
	var contentEncoding = h.Get(HeaderKeyContentEncoding)
	var commaIndex int = strings.IndexByte(contentEncoding, Comma)
	if commaIndex == -1 {
		commaIndex = len(contentEncoding)
	} else {
		h.Replace(HeaderKeyContentEncoding, contentEncoding[commaIndex+1:])
		multiple = true
	}
	ce = contentEncoding[:commaIndex]
	return
}

func (h *header) SetContentEncoding(contentEncodings ...string) {
	h.Sets(HeaderKeyContentEncoding, contentEncodings)
}
