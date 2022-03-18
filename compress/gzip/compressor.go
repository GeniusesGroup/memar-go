/* For license and copyright information please see LEGAL file in repository */

package gzip

import (
	"bytes"
	egzip "compress/gzip"
	"io"

	compress ".."
	"../../protocol"
)

type Compressor struct {
	source         protocol.Codec
	options        protocol.CompressOptions
	compressedData []byte
}

func (c *Compressor) init() {
	var b bytes.Buffer
	b.Grow(c.source.Len())
	var gz, _ = egzip.NewWriterLevel(&b, int(c.options.CompressLevel))
	gz.Write(c.source.Marshal())
	gz.Close()
	c.compressedData = b.Bytes()
}

/*
********** protocol.Codec interface **********
 */

func (c *Compressor) MediaType() protocol.MediaType       { return c.source.MediaType() }
func (c *Compressor) CompressType() protocol.CompressType { return &GZIP }

func (c *Compressor) Decode(reader protocol.Reader) (err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (c *Compressor) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = c.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (c *Compressor) Marshal() (data []byte) {
	if c.compressedData == nil {
		c.init()
	}
	return c.compressedData
}
func (c *Compressor) MarshalTo(data []byte) []byte {
	if c.compressedData == nil {
		c.init()
	}
	return append(data, c.compressedData...)
}
func (c *Compressor) Unmarshal(data []byte) (err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (c *Compressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}

// Len return length of compressed data
func (c *Compressor) Len() (ln int) {
	if c.compressedData == nil {
		c.init()
	}
	return len(c.compressedData)
}

/*
********** io package interfaces **********
 */

func (c *Compressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (c *Compressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var gz, _ = egzip.NewWriterLevel(w, int(c.options.CompressLevel))
	var writeLen int
	writeLen, err = gz.Write(c.source.Marshal())
	gz.Close()
	totalWrite = int64(writeLen)
	return
}
