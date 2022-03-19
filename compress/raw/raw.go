/* For license and copyright information please see LEGAL file in repository */

package raw

import (
	compress ".."
	"../../mediatype"
	"../../protocol"
)

const (
	RawContentEncoding = "raw"
	RawExtension       = ""
)

var RAW = raw{
	CompressType: compress.New(RawContentEncoding, mediatype.New("domain/raw.protocol.data-structure").SetFileExtension(RawExtension)),
}

type raw struct {
	*compress.CompressType
}

func (r *raw) Compress(raw protocol.Codec, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	return raw, nil
}
func (r *raw) CompressBySlice(raw []byte, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	compressed = &comDecom{
		data: raw,
	}
	return
}
func (r *raw) CompressByReader(raw protocol.Reader, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	// TODO:::
	return
}

func (r *raw) Decompress(compressed protocol.Codec) (raw protocol.Codec, err protocol.Error) {
	return compressed, nil
}
func (r *raw) DecompressFromSlice(compressed []byte) (raw protocol.Codec, err protocol.Error) {
	raw = &comDecom{
		data: compressed,
	}
	return
}
func (r *raw) DecompressFromReader(compressed protocol.Reader, compressedLen int) (raw protocol.Codec, err protocol.Error) {
	var rawDecoder = comDecom{
		reader:  compressed,
		readLen: compressedLen,
	}
	rawDecoder.Decode(compressed)
	raw = &rawDecoder
	return
}
