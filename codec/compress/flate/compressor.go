/* For license and copyright information please see the LEGAL file in the code repository */

package flate

import (
	"bytes"
	"compress/flate"
	"io"

	errs "memar/compress/errors"
	"memar/protocol"
)

type Compressor struct {
	source         protocol.Codec
	options        protocol.CompressOptions
	compressedData []byte
}

//memar:impl memar/protocol.ObjectLifeCycle
func (c *Compressor) Init(raw protocol.Codec, options protocol.CompressOptions) (err protocol.Error) {
	c.source = raw
	c.options = options
	return
}

func (c *Compressor) EncodeAll() (err protocol.Error) {
	var source = c.source

	var comData []byte
	comData, err = source.Marshal()
	if err != nil {
		return
	}

	var b bytes.Buffer
	b.Grow(c.source.Len())
	var def, _ = flate.NewWriter(&b, int(c.options.CompressLevel))
	def.Write(comData)
	def.Close()
	c.compressedData = b.Bytes()
	return
}

//memar:impl memar/protocol.Codec
func (c *Compressor) MediaType() protocol.MediaType       { return c.source.MediaType() }
func (c *Compressor) CompressType() protocol.CompressType { return &Deflate }

func (c *Compressor) Decode(source protocol.Codec) (n int, err protocol.Error) {
	err = &errs.ErrSourceNotChangeable
	return
}
func (c *Compressor) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	n, err = destination.Decode(c)
	return
}
func (c *Compressor) Marshal() (data []byte, err protocol.Error) {
	if c.compressedData == nil {
		err = c.EncodeAll()
	}
	data = c.compressedData
	return
}
func (c *Compressor) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	if c.compressedData == nil {
		err = c.EncodeAll()
	}
	added = append(data, c.compressedData...)
	return
}
func (c *Compressor) Unmarshal(data []byte) (n int, err protocol.Error) {
	err = &errs.ErrSourceNotChangeable
	return
}
func (c *Compressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = &errs.ErrSourceNotChangeable
	return
}

// Len return length of compressed data
func (c *Compressor) Len() (ln int) {
	if c.compressedData == nil {
		c.EncodeAll()
	}
	return len(c.compressedData)
}

/*
********** protocol.Buffer interface **********
 */

func (c *Compressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = &errs.ErrSourceNotChangeable
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
