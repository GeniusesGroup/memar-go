/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../codec"
	"../giti"
)

// Request is represent HTTP body!
// https://datatracker.ietf.org/doc/html/rfc2616#section-4.3
type body struct {
	giti.Codec
}

func (b body) Body() giti.Codec         { return b.Codec }
func (b body) SetBody(codec giti.Codec) { b.Codec = codec }

func (b body) checkEncodingAndSetBody(rawBody []byte, h *header) {
	// TODO::: first check body encoding!

	var rawBodyCodec = codec.Raw(rawBody)
	b.Codec = &rawBodyCodec
}

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
func (b body) setBodyByEncoding(rawBody []byte, h *header) {
	// httpRes.header.Set(HeaderKeyContentEncoding, "gzip")
	// var b bytes.Buffer
	// var gz = gzip.NewWriter(&b)
	// gz.Write(httpRes.Body)
	// gz.Close()
	// httpRes.Body = b.Bytes()
}
