/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"

	"../compress"
	"../protocol"
)

// body is represent HTTP body.
// Due to many performance impact, MediaType() method of body not return any true data. use header ContentType() method instead. This can be change if ...
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

func (b *body) setCodecAsIncomeBody(c protocol.Codec, h *header) {
	b.Codec = c
	// TODO::: check transferEncoding & contentEncoding
	// TODO::: What about header length maybe other than stream income data length e.g. send body in multiple TCP.PSH flag set.
}

func (b *body) setReaderAsIncomeBody(reader io.Reader, h *header) {
	var transferEncoding, _ = h.TransferEncoding()
	switch transferEncoding {
	case "":
		var contentLength = h.ContentLength()
		if contentLength > 0 {
			var contentEncoding, _ = h.ContentEncoding()
			// TODO::: add more encoding
			switch contentEncoding {
			case "":
				b.Codec = compress.RAWDecompressorFromReader(reader, contentLength)
			case HeaderValueChunked:
				// TODO:::
			case HeaderValueCompress:
				// TODO:::
			case HeaderValueDeflate:
				// TODO:::
			case HeaderValueGZIP:
				b.Codec = compress.GzipDecompressorFromReader(reader, int(contentLength))
			default:
				// TODO:::
			}
		}
	case HeaderValueChunked:
		// TODO:::
	default:
		// Like nginx, due to security, we only support a single Transfer-Encoding header field, and
		// only if set to "chunked".
		// err =
	}
}

func (b *body) setReadedIncomeBody(body []byte, h *header) {
	var contentEncoding, _ = h.ContentEncoding()
	// TODO::: add more encoding
	switch contentEncoding {
	case "":
		b.Codec = compress.RAWDecompressorFromSlice(body)
	case HeaderValueChunked:
		// TODO:::
	case HeaderValueCompress:
		// TODO:::
	case HeaderValueDeflate:
		// TODO:::
	case HeaderValueGZIP:
		b.Codec = compress.GzipDecompressorFromSlice(body)
	default:
		// TODO:::
	}
}

func (b *body) readFull(reader io.Reader, h *header) (n int64, goErr error) {
	var transferEncoding, _ = h.TransferEncoding()
	switch transferEncoding {
	case "":
		var contentLength = h.ContentLength()
		if contentLength > 0 {
			var bodyReadLength int
			var bodyRaw = make([]byte, contentLength)
			// TODO::: is below logic check include all situations? If body send by multi part e.g. multi TCP-PSH flag?
			bodyReadLength, goErr = reader.Read(bodyRaw)
			if bodyReadLength != int(contentLength) {
				// goErr =
			}

			b.setReadedIncomeBody(bodyRaw, h)
		}
	case HeaderValueChunked:
		// TODO:::
	default:
		// Like nginx, due to security, we only support a single Transfer-Encoding header field, and
		// only if set to "chunked".
		// err =
	}
	return
}
