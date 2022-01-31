/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"io"

	"../protocol"
)

const (
	RawContentEncoding = "raw"
	RawExtension       = ""
)

var (
	RAW = NewCompressType("urn:giti:compress.protocol:data-structure:raw", RawContentEncoding, RawExtension, RAWCompressor, RAWDecompressor)
)

func RAWCompressor(raw protocol.Codec, compressLevel protocol.CompressLevel) (compress protocol.Codec) {
	return raw
}

func RAWDecompressor(compress protocol.Codec) (rawCodec protocol.Codec) {
	return compress
}

func RAWDecompressorFromSlice(data []byte) (rawCodec protocol.Codec) {
	var rawDecoder raw
	rawDecoder.data = data
	return &rawDecoder
}

func RAWDecompressorFromReader(reader protocol.Reader, readLen uint64) (rawCodec protocol.Codec) {
	var rawDecoder raw
	rawDecoder.reader = reader
	rawDecoder.readLen = readLen
	rawDecoder.Decode(reader)
	return &rawDecoder
}

type raw struct {
	data    []byte
	reader  protocol.Reader
	readLen uint64
}

/*
********** protocol.Codec interface **********
 */

func (r *raw) MediaType() protocol.MediaType       { return nil }
func (r *raw) CompressType() protocol.CompressType { return RAW }
func (r *raw) Len() (ln int)                       { return len(r.data) }

func (r *raw) Decode(reader protocol.Reader) (err protocol.Error) {
	if r.data == nil && r.readLen > 0 {
		r.data = make([]byte, r.readLen)
		io.ReadFull(reader, r.data)
	}
	return
}
func (r *raw) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = r.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (r *raw) Marshal() (data []byte)                     { return r.data }
func (r *raw) MarshalTo(data []byte) []byte               { return append(data, r.data...) }
func (r *raw) Unmarshal(data []byte) (err protocol.Error) { err = ErrSourceNotChangeable; return }
func (r *raw) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}

/*
********** io package interfaces **********
 */

func (r *raw) ReadFrom(reader io.Reader) (n int64, err error) { err = ErrSourceNotChangeable; return }
func (r *raw) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	writeLen, err = w.Write(r.data)
	totalWrite = int64(writeLen)
	return
}
