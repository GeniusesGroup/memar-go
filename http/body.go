/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"

	"../compress"
	"../protocol"
)

// Request is represent HTTP body!
// https://datatracker.ietf.org/doc/html/rfc2616#section-4.3
type body struct {
	protocol.Codec
}

func (b *body) Body() protocol.Codec         { return b.Codec }
func (b *body) SetBody(codec protocol.Codec) { b.Codec = codec }
func (b *body) Len() int {
	if b.Codec != nil {
		return b.Codec.Len()
	}
	return 0
}

/*
********** local methods **********
 */

func (b *body) readFrom(reader io.Reader, h *header) (n int64, goErr error) {
	var transferEncoding, _ = h.TransferEncoding()
	switch transferEncoding {
	case "":
		break
	case HeaderValueChunked:
		// TODO:::
		return
	default:
		// Like nginx, due to security, we only support a single Transfer-Encoding header field, and
		// only if set to "chunked".
		// err =
		return
	}

	var contentLength = h.ContentLength()
	if contentLength > 0 {
		var bodyReadLength int
		var bodyRaw = make([]byte, contentLength)
		// TODO::: is below logic check include all situations? If body send by multi part e.g. multi TCP-PSH flag?
		bodyReadLength, goErr = reader.Read(bodyRaw)
		if bodyReadLength != int(contentLength) {
			// goErr =
		}

		b.setIncomeBody(bodyRaw, h)
	}

	return
}

func (b *body) setIncomeBody(body []byte, h *header) {
	var contentEncoding, _ = h.ContentEncoding()
	// TODO::: add more encoding
	switch contentEncoding {
	// TODO:::
	case "":
		b.Codec = compress.Raw(body)
	default:
		// TODO:::
	}
}
