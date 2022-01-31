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

func GzipCompressor(gzip protocol.Codec, compressLevel protocol.CompressLevel) (compress protocol.Codec) {
	return &gzipCompressor{
		source:        gzip,
		compressLevel: compressLevel,
	}
}

func GzipDecompressor(compressed protocol.Codec) (gzip protocol.Codec) {
	var gzipDecoder gzipDecompressor
	gzipDecoder.initByCodec(compressed)
	return &gzipDecoder
}

func GzipDecompressorFromSlice(compressed []byte) (gzip protocol.Codec) {
	var gzipDecoder gzipDecompressor
	gzipDecoder.Unmarshal(compressed)
	return &gzipDecoder
}

func GzipDecompressorFromReader(reader protocol.Reader, compressedLen int) (gzip protocol.Codec) {
	var gzipDecoder gzipDecompressor
	gzipDecoder.comLen = compressedLen
	gzipDecoder.Decode(reader)
	return &gzipDecoder
}
