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

func RAWDecompressor(compress protocol.Codec) (raw protocol.Codec) {
	return compress
}

type raw []byte

/*
********** protocol.Codec interface **********
 */

func (r raw) MediaType() protocol.MediaType       { return nil }
func (r raw) CompressType() protocol.CompressType { return RAW }

func (r raw) Decode(reader io.Reader) (err protocol.Error) { err = ErrSourceNotChangeable; return }
func (r raw) Encode(writer io.Writer) (err error)          { _, err = r.WriteTo(writer); return }
func (r raw) Marshal() (data []byte)                       { return r }
func (r raw) MarshalTo(data []byte) []byte                 { return append(data, r...) }
func (r raw) Unmarshal(data []byte) (err protocol.Error)   { err = ErrSourceNotChangeable; return }
func (r raw) Len() (ln int)                                { return len(r) }

/*
********** io package interfaces **********
 */

func (r raw) ReadFrom(reader io.Reader) (n int64, err error) { err = ErrSourceNotChangeable; return }
func (r raw) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	writeLen, err = w.Write(r)
	totalWrite = int64(writeLen)
	return
}
