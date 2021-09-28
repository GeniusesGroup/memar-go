/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"../protocol"
)

const (
	GZIPContentEncoding = "gzip"
	GZIPExtension       = "gzip"
)

var (
	GZIP = NewCompressType("urn:giti:compress.protocol:data-structure:gzip", GZIPContentEncoding, GZIPExtension, GzipCompressor, GzipDecompressor)
)

func GzipCompressor(raw protocol.Codec, compressLevel protocol.CompressLevel) (compress protocol.Codec) {
	var gzip = gzipCompressor{
		source:        raw,
		compressLevel: compressLevel,
	}
	return &gzip
}

func GzipDecompressor(compress protocol.Codec) (raw protocol.Codec) {
	var gzip = gzipDecompressor{
		source: compress,
	}
	return &gzip
}
