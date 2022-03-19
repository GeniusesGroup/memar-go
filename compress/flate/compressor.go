/* For license and copyright information please see LEGAL file in repository */

package flate

import (
	"bytes"
	"compress/flate"
	"io"

	compress ".."
	"../../protocol"
)

type Compressor struct {
	source         protocol.Codec
	options        protocol.CompressOptions
	compressedData []byte
}

func (d *Compressor) init() {
	var b bytes.Buffer
	b.Grow(d.source.Len())
	var def, _ = flate.NewWriter(&b, int(d.options.CompressLevel))
	def.Write(d.source.Marshal())
	def.Close()
	d.compressedData = b.Bytes()
}

/*
********** protocol.Codec interface **********
 */

func (d *Compressor) MediaType() protocol.MediaType       { return d.source.MediaType() }
func (d *Compressor) CompressType() protocol.CompressType { return &Deflate }

func (d *Compressor) Decode(reader protocol.Reader) (err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (d *Compressor) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = d.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (d *Compressor) Marshal() (data []byte) {
	if d.compressedData == nil {
		d.init()
	}
	return d.compressedData
}
func (d *Compressor) MarshalTo(data []byte) []byte {
	if d.compressedData == nil {
		d.init()
	}
	return append(data, d.compressedData...)
}
func (d *Compressor) Unmarshal(data []byte) (err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (d *Compressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}

// Len return length of compressed data
func (d *Compressor) Len() (ln int) {
	if d.compressedData == nil {
		d.init()
	}
	return len(d.compressedData)
}

/*
********** io package interfaces **********
 */

func (d *Compressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (d *Compressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var def, _ = flate.NewWriter(w, int(d.options.CompressLevel))
	var writeLen int
	writeLen, err = def.Write(d.source.Marshal())
	def.Close()
	totalWrite = int64(writeLen)
	return
}
