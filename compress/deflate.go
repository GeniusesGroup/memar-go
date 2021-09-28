/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"../protocol"
)

const (
	DeflateContentEncoding = "deflate"
	DeflateExtension       = "zz"
)

var (
	Deflate = NewCompressType("urn:giti:compress.protocol:data-structure:deflate", DeflateContentEncoding, DeflateExtension, DeflateCompressor, DeflateDecompressor)
)

func DeflateCompressor(raw protocol.Codec, compressLevel protocol.CompressLevel) (compress protocol.Codec) {
	var deflate = deflateCompressor{
		source:        raw,
		compressLevel: compressLevel,
	}
	return &deflate
}

func DeflateDecompressor(compress protocol.Codec) (raw protocol.Codec) {
	var deflate = deflateDecompressor{
		source: compress,
	}
	return &deflate
}
