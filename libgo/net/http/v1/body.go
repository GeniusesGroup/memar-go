/* For license and copyright information please see the LEGAL file in the code repository */

package http_v1

import (
	"io"

	cts "memar/codec/compress-types"
	compress_p "memar/codec/compress/protocol"
	codec_p "memar/codec/protocol"
	container_p "memar/computer/adt/container/protocol"
	"memar/computer/buffer"
	buffer_p "memar/computer/buffer/protocol"
	"memar/libgo/net/http"
	http_headers "memar/libgo/net/http/headers"
	error_p "memar/process/error/protocol"
)

// incomeBody is represent HTTP body.
// Due to many performance impact, MediaType() method of body not return any true data.
// use header ContentType() method instead. This can be change if ...
// https://datatracker.ietf.org/doc/html/rfc2616#section-4.3
type incomeBody struct {
	buf              buffer_p.Buffer
	contentLength    container_p.NumberOfElement
	transferEncoding http_headers.TransferEncodings
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (b *incomeBody) Init(buf buffer_p.Buffer, h *http.Header) (err error_p.Error) {
	b.buf = buf
	b.contentLength = h.ContentLength()
	b.transferEncoding = h.TransferEncoding()
	return
}
func (b *incomeBody) Reinit(buf buffer_p.Buffer, h *http.Header) (err error_p.Error) {
	err = b.Init(buf, h)
	return
}
func (b *incomeBody) Deinit() (err error_p.Error) { return }

func (b *incomeBody) checkAndSetCodecAsIncomeBody(maybeBody []byte, c codec_p.Codec, h *Header) (err error_p.Error) {
	var transferEncoding = h.TransferEncoding()
	switch transferEncoding {
	case "":
		var contentLength = h.ContentLength()
		var maybeBodyLength = len(maybeBody)
		if maybeBodyLength == int(contentLength) {
			b.setReadedIncomeBody(maybeBody, h)
		} else {
			// TODO::: allow this situation that peer send some part of body with header??
			// Header length maybe other than stream income data length e.g. send body in multiple TCP.PSH flag set.
			if maybeBodyLength > 0 {
				var bodySlice = make([]byte, maybeBodyLength, contentLength)
				copy(bodySlice, maybeBody)
				for {
					bodySlice, err = c.MarshalTo(bodySlice)
					if err != nil || len(bodySlice) == int(contentLength) {
						break
					}
				}
				b.setReadedIncomeBody(bodySlice, h)
			} else {
				b.setCodecAsIncomeBody(c, h)
			}
		}
	case http.HeaderValue_Chunked:
		// TODO:::
	default:
		// Like nginx, due to security, we only support a single Transfer-Encoding header field, and
		// only if set to "chunked".
		// err =
	}
	return
}

func (b *incomeBody) checkAndSetReaderAsIncomeBody(maybeBody []byte, reader io.Reader, h *Header) (err error_p.Error) {
	var transferEncoding = h.TransferEncoding()
	switch transferEncoding {
	case "":
		var contentLength = h.ContentLength()
		var maybeBodyLength = len(maybeBody)
		if maybeBodyLength == int(contentLength) {
			b.setReadedIncomeBody(maybeBody, h)
		} else {
			// Header length maybe other than stream income data length e.g. send body in multiple TCP.PSH flag set.
			if maybeBodyLength > 0 {
				var bodyReadLength int
				var goErr error
				var bodySlice = make([]byte, contentLength)
				bodyReadLength, goErr = reader.Read(bodySlice[maybeBodyLength:])
				if goErr != nil {
					// err =
					return
				}
				if bodyReadLength+maybeBodyLength != int(contentLength) {
					// err =
					return
				}
				copy(bodySlice, maybeBody)
				b.setReadedIncomeBody(bodySlice, h)
			} else {
				var bodyAsCodec = buffer.Flat.DecompressFromReader(reader, int(contentLength))
				b.setCodecAsIncomeBody(bodyAsCodec, h)
			}
		}
	case http.HeaderValue_Chunked:
		// TODO:::
	default:
		// Like nginx, due to security, we only support a single Transfer-Encoding header field, and
		// only if set to "chunked".
		// err =
	}
	return
}

// Call this method just if body marshaled with first line and headers.
func (b *incomeBody) checkAndSetIncomeBody(body []byte, h *Header) (err error_p.Error) {
	var bodyLength = len(body)
	if bodyLength > 0 {
		var contentLength = h.ContentLength()
		if bodyLength == int(contentLength) {
			b.setReadedIncomeBody(body, h)
		} else {
			// err = &errs.ErrBodySizeMismatch
		}
	}
	return
}

func (b *incomeBody) setCodecAsIncomeBody(c codec_p.Codec, h *Header) (err error_p.Error) {
	var contentEncoding = h.ContentEncodings()
	if contentEncoding == "" {
		b.Codec = c
		return
	}

	var compressType compress_p.CompressType
	compressType, err = cts.GetByContentEncoding(contentEncoding)
	if err != nil {
		return
	}
	b.Codec, err = compressType.Decompress(c)
	return
}

func (b *incomeBody) setReadedIncomeBody(body []byte, h *Header) (err error_p.Error) {
	var contentEncoding = h.ContentEncodings()
	if contentEncoding == "" {
		b.Codec, err = buffer.RAW.DecompressFromSlice(body)
		return
	}

	var compressType compress_p.CompressType
	compressType, err = cts.GetByContentEncoding(contentEncoding)
	if err != nil {
		return
	}
	b.Codec, err = compressType.DecompressFromSlice(body)
	return
}
