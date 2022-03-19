/* For license and copyright information please see LEGAL file in repository */

package flate

import (
	compress ".."
	"../../mediatype"
	"../../protocol"
)

const (
	ContentEncoding = "deflate"
	Extension       = "zz"
)

var Deflate = deflate{
	CompressType: compress.New(ContentEncoding, mediatype.New("domain/deflate.protocol.data-structure").SetFileExtension(Extension)),
}

type deflate struct {
	*compress.CompressType
}

func (d *deflate) Compress(raw protocol.Codec, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	compressed = &Compressor{
		source:  raw,
		options: options,
	}
	return
}
func (d *deflate) CompressBySlice(raw []byte, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	// TODO:::
	return
}
func (d *deflate) CompressByReader(raw protocol.Reader, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	// TODO:::
	return
}

func (d *deflate) Decompress(compressed protocol.Codec) (raw protocol.Codec, err protocol.Error) {
	raw = &Decompressor{
		source: compressed,
	}
	return
}
func (d *deflate) DecompressFromSlice(compressed []byte) (raw protocol.Codec, err protocol.Error) {
	// TODO:::
	return
}
func (d *deflate) DecompressFromReader(compressed protocol.Reader, compressedLen int) (raw protocol.Codec, err protocol.Error) {
	// TODO:::
	return
}
