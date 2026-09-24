/* For license and copyright information please see the LEGAL file in the code repository */

package flate

import (
	"bytes"
	"compress/flate"
	"io"

	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	errs "memar/codec/compress/errors"
	codec_p "memar/codec/protocol"
	error_p "memar/process/error/protocol"
)

type Decompressor struct {
	source           codec_p.Codec
	sourceMT         string
	decompressedData []byte
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (d *Decompressor) Init(source codec_p.Codec) (err error_p.Error) {
	d.source = source
	d.sourceMT = source.MediaType()
	return
}

func (d *Decompressor) decompressAll() (err error_p.Error) {
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
	decomBuf.Grow(int(source.SerializationLength()))
	decomBuf.ReadFrom(def)
	d.decompressedData = decomBuf.Bytes()
	return
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (d *Decompressor) MediaType() string { return d.sourceMT }

func (d *Decompressor) Decode(source buffer_p.Buffer) (err error_p.Error) {
	// err = &errs.SourceNotChangeable
	return
}
func (d *Decompressor) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	n, err = destination.Decode(d)
	return
}
func (d *Decompressor) Marshal() (data []byte, err error_p.Error) {
	if d.decompressedData == nil {
		err = d.decompressAll()
	}
	data = d.decompressedData
	return
}
func (d *Decompressor) MarshalTo(data []byte) (added []byte, err error_p.Error) {
	if d.decompressedData == nil {
		err = d.decompressAll()
	}
	added = append(data, d.decompressedData...)
	return
}
func (d *Decompressor) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	// err = &errs.SourceNotChangeable
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (d *Decompressor) SerializationLength() (ln container_p.NumberOfElement) {
	if d.decompressedData == nil {
		d.decompressAll()
	}
	ln = container_p.NumberOfElement(len(d.decompressedData))
	return
}

//memar:impl memar/buffer_p.Buffer
func (d *Decompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	// err = &errs.SourceNotChangeable
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
