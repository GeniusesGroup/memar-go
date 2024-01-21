/* For license and copyright information please see the LEGAL file in the code repository */

package http

import "strings"

// ContentEncoding return content encoding and notify if multiple exist
// To read multiple just call this method in a loop to get multiple became false
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
func (h *Header) ContentEncoding() ContentEncodings {
	var contentEncodings = h.Header_Get(HeaderKey_ContentEncoding)
	return ContentEncodings(contentEncodings)
}

func (h *Header) AddContentEncoding(contentEncodings string) {
	h.Header_Add(HeaderKey_ContentEncoding, contentEncodings)
}

type ContentEncodings string

// Last return last ContentEncoding and remove it from ContentEncodings.
//
// https://www.rfc-editor.org/rfc/rfc9110.html#section-8.4
// If one or more encodings have been applied to a representation,
// the sender that applied the encodings MUST generate a Content-Encoding header field
// that lists the content codings in the order in which they were applied.
// In other words, decode in the reverse order to the order in the header.
func (c *ContentEncodings) Last() (contentEncoding string, exist bool) {
	var ce = string(*c)

	if len(ce) == 0 {
		exist = false
		return
	}

	var commaIndex int = strings.LastIndexByte(ce, Comma)
	if commaIndex == -1 {
		commaIndex = 0
		*c = ""
	} else {
		*c = ContentEncodings(ce[:commaIndex-1])
	}
	contentEncoding = ce[commaIndex+1:]
	exist = true
	return
}
