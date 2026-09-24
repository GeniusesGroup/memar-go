/* For license and copyright information please see the LEGAL file in the code repository */

package http_headers

import (
	http_p "memar/libgo/net/http/protocol"
	"strings"
)

func AddContentEncoding(h http_p.Header, contentEncodings string) {
	h.Header_Add(Key_ContentEncoding, contentEncodings)
}

// ContentEncodings return content encodings
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
type ContentEncodings string

func (self *ContentEncodings) Init(h http_p.Header) {
	var contentEncodings = h.Header_Get(Key_ContentEncoding)
	self = ContentEncodings(contentEncodings)
}

// Last return last ContentEncoding and remove it from ContentEncodings.
// To read multiple just call this method in a loop to get multiple became false
//
// https://www.rfc-editor.org/rfc/rfc9110.html#section-8.4
// If one or more encodings have been applied to a representation,
// the sender that applied the encodings MUST generate a Content-Encoding header field
// that lists the content codings in the order in which they were applied.
// In other words, decode in the reverse order to the order in the header.
func (self *ContentEncodings) Last() (contentEncoding string, exist bool) {
	var ce = string(*self)

	if len(ce) == 0 {
		exist = false
		return
	}

	var commaIndex int = strings.LastIndexByte(ce, Comma)
	if commaIndex == -1 {
		commaIndex = 0
		*self = ""
	} else {
		*self = ContentEncodings(ce[:commaIndex-1])
	}
	contentEncoding = ce[commaIndex+1:]
	exist = true
	return
}
