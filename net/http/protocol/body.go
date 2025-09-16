/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	codec_p "memar/codec/protocol"
)

// HTTP Body Semantic that USUALLY use in responses.
// In requests ALMOST ALWAYS each service HTTPHandler use [SK Socket] Buffer to decode to desire data type.
type Field_Body interface {
	Body() Body
	SetBody(codec Body)
}

type Body interface {
	codec_p.Codec
}
