/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"strings"
)

// TransferEncoding return transfer encoding and notify if multiple exist
// To read multiple just call this method in a loop to get multiple became false
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding
// https://datatracker.ietf.org/doc/html/rfc2616#section-3.6
func (h *Header) TransferEncoding() TransferEncodings {
	var transferEncodings = h.Header_Get(HeaderKey_TransferEncoding)
	return TransferEncodings(transferEncodings)
}

// AddTransferEncoding add transfer encoding.
func (h *Header) AddTransferEncoding(te string) {
	h.Header_Add(HeaderKey_TransferEncoding, te)
}

type TransferEncodings string

// Last return last TransferEncoding and remove it from TransferEncodings.
//
// https://www.rfc-editor.org/rfc/rfc9112#section-6.1
// If one or more encodings have been applied to a representation,
// the sender that applied the encodings MUST generate a Transfer-Encoding header field
// that lists the content codings in the order in which they were applied.
// In other words, decode in the reverse order to the order in the header.
func (c *TransferEncodings) Last() (contentEncoding string, exist bool) {
	var te = string(*c)

	if len(te) == 0 {
		exist = false
		return
	}

	var commaIndex int = strings.LastIndexByte(te, Comma)
	if commaIndex == -1 {
		commaIndex = 0
		*c = ""
	} else {
		*c = TransferEncodings(te[:commaIndex-1])
	}
	contentEncoding = te[commaIndex+1:]
	exist = true
	return
}
