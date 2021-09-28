/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"bytes"
	"compress/flate"
	"io"

	"../protocol"
)

type deflateCompressor struct {
	source         protocol.Codec
	compressLevel  protocol.CompressLevel
	compressedData []byte
}

func (d *deflateCompressor) init() {
	var b bytes.Buffer
	b.Grow(d.source.Len())
	var def, _ = flate.NewWriter(&b, int(d.compressLevel))
	def.Write(d.source.Marshal())
	def.Close()
	d.compressedData = b.Bytes()
}

/*
********** protocol.Codec interface **********
 */

func (d *deflateCompressor) MediaType() protocol.MediaType       { return d.source.MediaType() }
func (d *deflateCompressor) CompressType() protocol.CompressType { return Deflate }

func (d *deflateCompressor) Decode(reader io.Reader) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}
func (d *deflateCompressor) Encode(writer io.Writer) (err error) { _, err = d.WriteTo(writer); return }
func (d *deflateCompressor) Marshal() (data []byte) {
	if d.compressedData == nil {
		d.init()
	}
	return d.compressedData
}
func (d *deflateCompressor) MarshalTo(data []byte) []byte {
	if d.compressedData == nil {
		d.init()
	}
	return append(data, d.compressedData...)
}
func (d *deflateCompressor) Unmarshal(data []byte) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}

// Len return length of compressed data
func (d *deflateCompressor) Len() (ln int) {
	if d.compressedData == nil {
		d.init()
	}
	return len(d.compressedData)
}

/*
********** io package interfaces **********
 */

func (d *deflateCompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = ErrSourceNotChangeable
	return
}
func (d *deflateCompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var def, _ = flate.NewWriter(w, int(d.compressLevel))
	var writeLen int
	writeLen, err = def.Write(d.source.Marshal())
	def.Close()
	totalWrite = int64(writeLen)
	return
}
