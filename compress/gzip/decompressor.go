/* For license and copyright information please see LEGAL file in repository */

package gzip

import (
	"bytes"
	egzip "compress/gzip"
	"io"

	"../../protocol"
)

type Decompressor struct {
	sourceMT         protocol.MediaType
	comBuf           bytes.Buffer
	comLen           int
	zr               egzip.Reader
	decompressedData []byte
}

func (c *Decompressor) initByCodec(compressed protocol.Codec) {
	c.sourceMT = compressed.MediaType()
	c.comLen = compressed.Len()
	c.comBuf = *bytes.NewBuffer(compressed.Marshal())
	if err := c.zr.Reset(&c.comBuf); err != nil {
		return
	}
}

func (c *Decompressor) decompressAll() {
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
}

/*
********** protocol.Codec interface **********
 */

func (c *Decompressor) MediaType() protocol.MediaType       { return c.sourceMT }
func (c *Decompressor) CompressType() protocol.CompressType { return nil }

func (c *Decompressor) Decode(reader protocol.Reader) (err protocol.Error) {
	var goErr = c.zr.Reset(reader)
	if goErr != nil {
		// err =
	}
	return
}
func (c *Decompressor) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = c.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (c *Decompressor) Marshal() (data []byte) {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	return c.decompressedData
}
func (c *Decompressor) MarshalTo(data []byte) []byte {
	if c.decompressedData == nil {
		c.decompressAll()
	}
	return append(data, c.decompressedData...)
}
func (c *Decompressor) Unmarshal(data []byte) (err protocol.Error) {
	c.comLen = len(data)
	c.comBuf = *bytes.NewBuffer(data)
	var goErr = c.zr.Reset(&c.comBuf)
	if goErr != nil {
		// err =
	}
	return
}
func (c *Decompressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = c.Unmarshal(data)
	// TODO::: can return any remaining data?
	return
}

// Len return length of decompressed data
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
********** io package interfaces **********
 */

func (c *Decompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = c.zr.Reset(reader)
	return
}
func (c *Decompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	return io.Copy(w, &c.zr)
}
