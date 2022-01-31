/* For license and copyright information please see LEGAL file in repository */

package http

import (
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

/*
********** protocol.Codec interface **********
 */

func (b *body) Len() int {
	if b.Codec != nil {
		return b.Codec.Len()
	}
	return 0
}
func (b *body) MediaType() protocol.MediaType {
	if b.Codec != nil {
		return b.Codec.MediaType()
	}
	return nil
}
func (b *body) CompressType() protocol.CompressType {
	if b.Codec != nil {
		return b.Codec.CompressType()
	}
	return nil
}
func (b *body) Decode(reader protocol.Reader) (err protocol.Error) {
	if b.Codec != nil {
		err = b.Codec.Decode(reader)
	}
	return
}
func (b *body) Encode(writer protocol.Writer) (err protocol.Error) {
	if b.Codec != nil {
		err = b.Codec.Encode(writer)
	}
	return
}
func (b *body) Marshal() (data []byte) {
	if b.Codec != nil {
		data = b.Codec.Marshal()
	}
	return
}
func (b *body) MarshalTo(data []byte) []byte {
	if b.Codec != nil {
		return b.Codec.MarshalTo(data)
	}
	return data
}
func (b *body) Unmarshal(data []byte) (err protocol.Error) {
	if b.Codec != nil {
		err = b.Codec.Unmarshal(data)
	}
	return
}
func (b *body) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	if b.Codec != nil {
		return b.Codec.UnmarshalFrom(data)
	}
	return
}

/*
********** local methods **********
 */

func (b *body) checkAndSetCodecAsIncomeBody(maybeBody []byte, c protocol.Codec, h *header) {
	// TODO::: check h.TransferEncoding() and h.ContentLength()??
	if len(maybeBody) > 0 {
		b.setReadedIncomeBody(maybeBody, h)
	} else {
		b.setCodecAsIncomeBody(c, h)
	}
}

func (b *body) checkAndSetReaderAsIncomeBody(maybeBody []byte, reader protocol.Reader, h *header) {
	// TODO::: check h.TransferEncoding() and h.ContentLength()??
	if len(maybeBody) > 0 {
		// check if body sent with header in one buffer
		b.setReadedIncomeBody(maybeBody, h)
	} else {
		b.setReaderAsIncomeBody(reader, h)
	}
}

func (b *body) checkAndSetIncomeBody(maybeBody []byte, h *header) (err protocol.Error) {
	// TODO::: check h.TransferEncoding() and h.ContentLength()??
	if len(maybeBody) > 0 {
		// Just if body marshaled with first line and headers we need to do any action here
		b.setReadedIncomeBody(maybeBody, h)
	} else {
		// err =
	}
	return
}

func (b *body) setCodecAsIncomeBody(c protocol.Codec, h *header) {
	b.Codec = c
	// TODO::: check transferEncoding & contentEncoding
	// TODO::: What about header length maybe other than stream income data length e.g. send body in multiple TCP.PSH flag set.
}

func (b *body) setReaderAsIncomeBody(reader protocol.Reader, h *header) {
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

func (b *body) readFull(reader protocol.Reader, h *header) (n int64, goErr error) {
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
