/* For license and copyright information please see LEGAL file in repository */

package flate

import (
	"bytes"
	"compress/flate"
	"io"

	compress ".."
	"../../protocol"
)

type Decompressor struct {
	source           protocol.Codec
	decompressedData []byte
}

func (d *Decompressor) init() {
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

func (d *Decompressor) MediaType() protocol.MediaType       { return d.source.MediaType() }
func (d *Decompressor) CompressType() protocol.CompressType { return nil }

func (d *Decompressor) Decode(reader protocol.Reader) (err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (d *Decompressor) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = d.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (d *Decompressor) Marshal() (data []byte) {
	if d.decompressedData == nil {
		d.init()
	}
	return d.decompressedData
}
func (d *Decompressor) MarshalTo(data []byte) []byte {
	if d.decompressedData == nil {
		d.init()
	}
	return append(data, d.decompressedData...)
}
func (d *Decompressor) Unmarshal(data []byte) (err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (d *Decompressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}

// Len return length of decompressed data
func (d *Decompressor) Len() (ln int) {
	if d.decompressedData == nil {
		d.init()
	}
	return len(d.decompressedData)
}

/*
********** io package interfaces **********
 */

func (d *Decompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (d *Decompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var buf = bytes.NewBuffer(d.source.Marshal())
	var def = flate.NewReader(buf)
	return io.Copy(w, def)
}
