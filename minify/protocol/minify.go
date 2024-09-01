/* For license and copyright information please see the LEGAL file in the code repository */

package minify_p

import (
	codec_p "memar/codec/protocol"
	"memar/protocol"
)

// Minify replace given data with minify of them if possible.
type Minifier interface {
	Minify(data codec_p.Codec) (err protocol.Error)
	MinifyBytes(data []byte) (minified []byte, err protocol.Error)
}
