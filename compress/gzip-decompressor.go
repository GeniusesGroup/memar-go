/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"bytes"
	"compress/gzip"
	"io"

	"../protocol"
)

type gzipDecompressor struct {
	source           protocol.Codec
	decompressedData []byte
}

func (g *gzipDecompressor) init() {
	var comBuf = bytes.NewBuffer(g.source.Marshal())
	var gz, _ = gzip.NewReader(comBuf)

	// TODO::: which solution?
	// g.decompressedData, _ = io.ReadAll(gz)
	var decomBuf bytes.Buffer
	decomBuf.Grow(g.source.Len())
	decomBuf.ReadFrom(gz)
	g.decompressedData = decomBuf.Bytes()
}

/*
********** protocol.Codec interface **********
 */

func (g *gzipDecompressor) MediaType() protocol.MediaType       { return g.source.MediaType() }
func (g *gzipDecompressor) CompressType() protocol.CompressType { return nil }

func (g *gzipDecompressor) Decode(reader io.Reader) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}
func (g *gzipDecompressor) Encode(writer io.Writer) (err error) { _, err = g.WriteTo(writer); return }
func (g *gzipDecompressor) Marshal() (data []byte) {
	if g.decompressedData == nil {
		g.init()
	}
	return g.decompressedData
}
func (g *gzipDecompressor) MarshalTo(data []byte) []byte {
	if g.decompressedData == nil {
		g.init()
	}
	return append(data, g.decompressedData...)
}
func (g *gzipDecompressor) Unmarshal(data []byte) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}

// Len return length of decompressed data
func (g *gzipDecompressor) Len() (ln int) {
	if g.decompressedData == nil {
		g.init()
	}
	return len(g.decompressedData)
}

/*
********** io package interfaces **********
 */

func (g *gzipDecompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = ErrSourceNotChangeable
	return
}
func (g *gzipDecompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var buf = bytes.NewBuffer(g.source.Marshal())
	var gz, _ = gzip.NewReader(buf)
	return io.Copy(w, gz)
}
