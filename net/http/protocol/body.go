/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	codec_p "memar/codec/protocol"
)

// HTTP Body Semantic that USUALLY use in responses.
// In requests ALMOST ALWAYS each service HTTPHandler use [SK Socket] Buffer to decode to desire data type.
type Body interface {
	Body() codec_p.Codec
	SetBody(codec codec_p.Codec)
}
