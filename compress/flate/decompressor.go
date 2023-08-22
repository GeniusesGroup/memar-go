/* For license and copyright information please see the LEGAL file in the code repository */

package flate

import (
	"bytes"
	"compress/flate"
	"io"

	errs "memar/compress/errors"
	"memar/protocol"
)

type Decompressor struct {
	source           protocol.Codec
	sourceMT         protocol.MediaType
	decompressedData []byte
}

//memar:impl memar/protocol.ObjectLifeCycle
func (d *Decompressor) Init(source protocol.Codec) (err protocol.Error) {
	d.source = source
	d.sourceMT = source.MediaType()
	return
}

func (d *Decompressor) decompressAll() (err protocol.Error) {
	var source = d.source

	var comData []byte
	comData, err = source.Marshal()
	if err != nil {
		return
	}

	var comBuf = bytes.NewBuffer(comData)
	var def = flate.NewReader(comBuf)

	// TODO::: which solution?
	// d.decompressedData, _ = io.ReadAll(def)
	var decomBuf bytes.Buffer
	decomBuf.Grow(source.Len())
	decomBuf.ReadFrom(def)
	d.decompressedData = decomBuf.Bytes()
	return
}

//memar:impl memar/protocol.Codec
func (d *Decompressor) MediaType() protocol.MediaType       { return d.sourceMT }
func (d *Decompressor) CompressType() protocol.CompressType { return nil }

func (d *Decompressor) Decode(source protocol.Codec) (n int, err protocol.Error) {
	err = &errs.ErrSourceNotChangeable
	return
}
func (d *Decompressor) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	n, err = destination.Decode(d)
	return
}
func (d *Decompressor) Marshal() (data []byte, err protocol.Error) {
	if d.decompressedData == nil {
		err = d.decompressAll()
	}
	data = d.decompressedData
	return
}
func (d *Decompressor) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	if d.decompressedData == nil {
		err = d.decompressAll()
	}
	added = append(data, d.decompressedData...)
	return
}
func (d *Decompressor) Unmarshal(data []byte) (n int, err protocol.Error) {
	err = &errs.ErrSourceNotChangeable
	return
}
func (d *Decompressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = &errs.ErrSourceNotChangeable
	return
}

// Len return length of decompressed data
func (d *Decompressor) Len() (ln int) {
	if d.decompressedData == nil {
		d.decompressAll()
	}
	return len(d.decompressedData)
}

//memar:impl memar/protocol.Buffer
func (d *Decompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = &errs.ErrSourceNotChangeable
	return
}
func (d *Decompressor) WriteTo(w io.Writer) (totalWrite int64, goErr error) {
	var comData, err = d.source.Marshal()
	if err != nil {
		goErr = err
		return
	}
	var buf = bytes.NewBuffer(comData)
	var def = flate.NewReader(buf)
	return io.Copy(w, def)
}
