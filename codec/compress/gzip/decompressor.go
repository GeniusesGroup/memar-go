/* For license and copyright information please see the LEGAL file in the code repository */

package gzip

import (
	"bytes"
	egzip "compress/gzip"
	"io"

	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	codec_p "memar/codec/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	error_p "memar/process/error/protocol"
)

type Decompressor struct {
	source           codec_p.Codec
	sourceMT         mediatype_p.MediaType
	comBuf           bytes.Buffer
	comLen           container_p.NumberOfElement
	zr               egzip.Reader
	decompressedData []byte
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (c *Decompressor) Init(compressed codec_p.Codec) (err error_p.Error) {
	c.source = compressed
	c.sourceMT = compressed.MediaType()
	c.comLen = compressed.SerializationLength()
	return
}

func (c *Decompressor) decompressAll() (err error_p.Error) {
	var comData []byte
	comData, err = c.source.Marshal()
	if err != nil {
		return
	}

	c.comBuf = *bytes.NewBuffer(comData)
	var goErr = c.zr.Reset(&c.comBuf)
	if goErr != nil {
		// err =
		return
	}

	// TODO::: which solution?
	// c.decompressedData, _ = io.ReadAll(gz)
	var decomBuf bytes.Buffer
	if c.comLen < 1 {
		decomBuf.Grow(512) // TODO::: 512 came from io.ReadAll, any suggestion??
	} else {
		decomBuf.Grow(c.comLen)
	}
	decomBuf.ReadFrom(&c.zr)
	c.decompressedData = decomBuf.Bytes()
	return
}

// func (c *Decompressor) MediaType() mediatype_p.MediaType    { return c.sourceMT }
// func (c *Decompressor) CompressType() compress_p.CompressType { return nil }
//
//memar:impl memar/codec/protocol.Codec
func (c *Decompressor) Decode(source buffer_p.Buffer) (err error_p.Error) {
	// err = &errs.SourceNotChangeable
	return
}
func (c *Decompressor) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	err = destination.Decode(c)
	return
}
func (c *Decompressor) Marshal() (data []byte, err error_p.Error) {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	data = c.decompressedData
	return
}
func (c *Decompressor) MarshalTo(data []byte) (added []byte, err error_p.Error) {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	added = append(data, c.decompressedData...)
	return
}
func (c *Decompressor) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	c.comLen = len(source)
	c.comBuf = *bytes.NewBuffer(source)
	var goErr = c.zr.Reset(&c.comBuf)
	if goErr != nil {
		// err =
	}
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (c *Decompressor) SerializationLength() (ln container_p.NumberOfElement) {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	ln = container_p.NumberOfElement(len(c.decompressedData))
	return
}

func (c *Decompressor) Discard() (err error_p.Error) {
	// io.Copy(io.Discard, b)
	return
}

/*
********** buffer_p.Buffer interface **********
 */

func (c *Decompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = c.zr.Reset(reader)
	return
}
func (c *Decompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	return io.Copy(w, &c.zr)
}
