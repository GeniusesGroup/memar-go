/* For license and copyright information please see the LEGAL file in the code repository */

package gzip

import (
	"bytes"
	egzip "compress/gzip"
	"io"

	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	"memar/codec/compress"
	compress_errs "memar/codec/compress/errors"
	compress_p "memar/codec/compress/protocol"
	codec_p "memar/codec/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	error_p "memar/process/error/protocol"
)

type Compressor struct {
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
	b.Grow(int(source.SerializationLength()))
	var gz, _ = egzip.NewWriterLevel(&b, int(c.options.CompressLevel))
	gz.Write(comData)
	gz.Close()
	c.compressedData = b.Bytes()
	return
}

//memar:impl memar/codec/protocol.Codec
func (c *Compressor) MediaType() mediatype_p.MediaType      { return c.source.MediaType() }
func (c *Compressor) CompressType() compress_p.CompressType { return &GZIP }

func (c *Compressor) Decode(source buffer_p.Buffer) (err error_p.Error) {
	err = &compress_errs.SourceNotChangeable
	return
}
func (c *Compressor) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	n, err = destination.Decode(c)
	return
}
func (c *Compressor) Marshal() (data []byte, err error_p.Error) {
	if c.compressedData == nil {
		c.EncodeAll()
	}
	data = c.compressedData
	return
}
func (c *Compressor) MarshalTo(data []byte) (added []byte, err error_p.Error) {
	if c.compressedData == nil {
		c.EncodeAll()
	}
	added = append(data, c.compressedData...)
	return
}
func (c *Compressor) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	err = &compress_errs.SourceNotChangeable
	return
}

//memar:impl memar/codec/protocol.Field_Length
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
	err = &compress_errs.SourceNotChangeable
	return
}
func (c *Compressor) WriteTo(w io.Writer) (totalWrite int64, goErr error) {
	var comData, err = c.source.Marshal()
	if err != nil {
		goErr = err
		return
	}

	var gz, _ = egzip.NewWriterLevel(w, int(c.options.CompressLevel))
	var writeLen int
	writeLen, goErr = gz.Write(comData)
	gz.Close()
	totalWrite = int64(writeLen)
	return
}
