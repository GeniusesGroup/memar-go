/* For license and copyright information please see the LEGAL file in the code repository */

package flate

import (
	"bytes"
	"compress/flate"
	"io"

	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	"memar/codec/compress"
	// compress_errs "memar/codec/compress/errors"
	codec_p "memar/codec/protocol"
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
)

type Compressor struct {
	datatype.DataType
	source         codec_p.Codec
	options        compress.Options
	compressedData []byte
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (c *Compressor) Init(raw codec_p.Codec, options compress.Options) (err error_p.Error) {
	c.source = raw
	c.options = options
	return
}

func (c *Compressor) EncodeAll() (err error_p.Error) {
	var source = c.source

	var comData []byte
	comData, err = source.Marshal()
	if err != nil {
		return
	}

	var b bytes.Buffer
	b.Grow(int(c.source.SerializationLength()))
	var def, _ = flate.NewWriter(&b, int(c.options.CompressLevel))
	def.Write(comData)
	def.Close()
	c.compressedData = b.Bytes()
	return
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (c *Compressor) MediaType() string { return "" }

//memar:impl memar/codec/protocol.Codec
func (c *Compressor) Decode(source buffer_p.Buffer) (err error_p.Error) {
	// err = &compress_errs.SourceNotChangeable
	return
}
func (c *Compressor) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	err = destination.Decode(c)
	return
}
func (c *Compressor) Marshal() (data []byte, err error_p.Error) {
	if c.compressedData == nil {
		err = c.EncodeAll()
	}
	data = c.compressedData
	return
}
func (c *Compressor) MarshalTo(data []byte) (added []byte, err error_p.Error) {
	if c.compressedData == nil {
		err = c.EncodeAll()
	}
	added = append(data, c.compressedData...)
	return
}
func (c *Compressor) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	// err = &compress_errs.SourceNotChangeable
	return
}

// Len return length of compressed data
func (c *Compressor) SerializationLength() (ln container_p.NumberOfElement) {
	if c.compressedData == nil {
		c.EncodeAll()
	}
	ln = container_p.NumberOfElement(len(c.compressedData))
	return
}

/*
********** buffer_p.Buffer interface **********
 */

func (c *Compressor) ReadFrom(reader io.Reader) (n int64, err error) {
	// err = &compress_errs.SourceNotChangeable
	return
}
func (c *Compressor) WriteTo(w io.Writer) (totalWrite int64, goErr error) {
	var comData, err = c.source.Marshal()
	if err != nil {
		goErr = err
		return
	}

	var def, _ = flate.NewWriter(w, int(c.options.CompressLevel))
	var writeLen int
	writeLen, goErr = def.Write(comData)
	def.Close()
	totalWrite = int64(writeLen)
	return
}
