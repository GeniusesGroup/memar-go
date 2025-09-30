/* For license and copyright information please see the LEGAL file in the code repository */

package gzip

import (
	"bytes"
	egzip "compress/gzip"
	"io"

	errs "memar/compress/errors"
	"memar/protocol"
)

type Decompressor struct {
	source           protocol.Codec
	sourceMT         protocol.MediaType
	comBuf           bytes.Buffer
	comLen           int
	zr               egzip.Reader
	decompressedData []byte
}

//memar:impl memar/protocol.ObjectLifeCycle
func (c *Decompressor) Init(compressed protocol.Codec) (err protocol.Error) {
	c.source = compressed
	c.sourceMT = compressed.MediaType()
	c.comLen = compressed.Len()
	return
}

func (c *Decompressor) decompressAll() (err protocol.Error) {
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

//memar:impl memar/protocol.Codec
func (c *Decompressor) MediaType() protocol.MediaType       { return c.sourceMT }
func (c *Decompressor) CompressType() protocol.CompressType { return nil }
func (c *Decompressor) Decode(source protocol.Codec) (n int, err protocol.Error) {
	err = &errs.ErrSourceNotChangeable
	return
}
func (c *Decompressor) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	n, err = destination.Decode(c)
	return
}
func (c *Decompressor) Marshal() (data []byte, err protocol.Error) {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	data = c.decompressedData
	return
}
func (c *Decompressor) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	added = append(data, c.decompressedData...)
	return
}
func (c *Decompressor) Unmarshal(data []byte) (n int, err protocol.Error) {
	c.comLen = len(data)
	c.comBuf = *bytes.NewBuffer(data)
	var goErr = c.zr.Reset(&c.comBuf)
	if goErr != nil {
		// err =
	}
	return
}
func (c *Decompressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	var n int
	n, err = c.Unmarshal(data)
	remaining = data[n:]
	return
}
func (c *Decompressor) Len() (ln int) {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	return len(c.decompressedData)
}

func (c *Decompressor) Discard() (err protocol.Error) {
	// io.Copy(io.Discard, b)
	return
}

/*
********** protocol.Buffer interface **********
 */

func (c *Decompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = c.zr.Reset(reader)
	return
}
func (c *Decompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	return io.Copy(w, &c.zr)
}
