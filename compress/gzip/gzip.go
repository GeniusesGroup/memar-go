/* For license and copyright information please see LEGAL file in repository */

package gzip

import (
	compress ".."
	"../../mediatype"
	"../../protocol"
)

const (
	GZIPContentEncoding = "gzip"
	GZIPExtension       = "gzip"
)

var GZIP = gzip{
	CompressType: compress.New(GZIPContentEncoding, mediatype.New("domain/gzip.protocol.data-structure").SetFileExtension(GZIPExtension)),
}

type gzip struct {
	*compress.CompressType
}

func (g *gzip) Compress(raw protocol.Codec, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	compressed = &Compressor{
		source:  raw,
		options: options,
	}
	return
}
func (g *gzip) CompressBySlice(raw []byte, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	// TODO:::
	return
}
func (g *gzip) CompressByReader(raw protocol.Reader, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	// TODO:::
	return
}

func (g *gzip) Decompress(compressed protocol.Codec) (raw protocol.Codec, err protocol.Error) {
	var gzipDecoder Decompressor
	gzipDecoder.initByCodec(compressed)
	raw = &gzipDecoder
	return
}
func (g *gzip) DecompressFromSlice(compressed []byte) (raw protocol.Codec, err protocol.Error) {
	var gzipDecoder Decompressor
	err = gzipDecoder.Unmarshal(compressed)
	raw = &gzipDecoder
	return
}
func (g *gzip) DecompressFromReader(compressed protocol.Reader, compressedLen int) (raw protocol.Codec, err protocol.Error) {
	var gzipDecoder Decompressor
	gzipDecoder.comLen = compressedLen
	err = gzipDecoder.Decode(compressed)
	raw = &gzipDecoder
	return
}
