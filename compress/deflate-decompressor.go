/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"bytes"
	"compress/flate"
	"io"

	"../protocol"
)

type deflateDecompressor struct {
	source           protocol.Codec
	decompressedData []byte
}

func (d *deflateDecompressor) init() {
	var comBuf = bytes.NewBuffer(d.source.Marshal())
	var def = flate.NewReader(comBuf)

	// TODO::: which solution?
	// d.decompressedData, _ = io.ReadAll(def)
	var decomBuf bytes.Buffer
	decomBuf.Grow(d.source.Len())
	decomBuf.ReadFrom(def)
	d.decompressedData = decomBuf.Bytes()
}

/*
********** protocol.Codec interface **********
 */

func (d *deflateDecompressor) MediaType() protocol.MediaType       { return d.source.MediaType() }
func (d *deflateDecompressor) CompressType() protocol.CompressType { return Deflate }

func (d *deflateDecompressor) Decode(reader io.Reader) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}
func (d *deflateDecompressor) Encode(writer io.Writer) (err error) {
	_, err = d.WriteTo(writer)
	return
}
func (d *deflateDecompressor) Marshal() (data []byte) {
	if d.decompressedData == nil {
		d.init()
	}
	return d.decompressedData
}
func (d *deflateDecompressor) MarshalTo(data []byte) []byte {
	if d.decompressedData == nil {
		d.init()
	}
	return append(data, d.decompressedData...)
}
func (d *deflateDecompressor) Unmarshal(data []byte) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}

// Len return length of decompressed data
func (d *deflateDecompressor) Len() (ln int) {
	if d.decompressedData == nil {
		d.init()
	}
	return len(d.decompressedData)
}

/*
********** io package interfaces **********
 */

func (d *deflateDecompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = ErrSourceNotChangeable
	return
}
func (d *deflateDecompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var buf = bytes.NewBuffer(d.source.Marshal())
	var def = flate.NewReader(buf)
	return io.Copy(w, def)
}
