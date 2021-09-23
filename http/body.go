/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../codec"
	"../protocol"
)

// Request is represent HTTP body!
// https://datatracker.ietf.org/doc/html/rfc2616#section-4.3
type body struct {
	protocol.Codec
}

func (b body) Body() protocol.Codec         { return b.Codec }
func (b body) SetBody(codec protocol.Codec) { b.Codec = codec }

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
func (b body) setIncomeBody(rawBody []byte, h *header) {
	// var bodyContentEncodingName = h.Get(HeaderKeyContentEncoding)
	// TODO:::
	var rawBodyCodec = codec.Raw(rawBody)
	b.Codec = &rawBodyCodec
}
