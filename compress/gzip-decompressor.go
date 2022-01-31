/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"bytes"
	"compress/gzip"
	"io"

	"../protocol"
)

type gzipDecompressor struct {
	sourceMT         protocol.MediaType
	comBuf           bytes.Buffer
	comLen           int
	zr               gzip.Reader
	decompressedData []byte
}

func (g *gzipDecompressor) initByCodec(compressed protocol.Codec) {
	g.sourceMT = compressed.MediaType()
	g.comLen = compressed.Len()
	g.comBuf = *bytes.NewBuffer(compressed.Marshal())
	if err := g.zr.Reset(&g.comBuf); err != nil {
		return
	}
}

func (g *gzipDecompressor) decompressAll() {
	// TODO::: which solution?
	// g.decompressedData, _ = io.ReadAll(gz)
	var decomBuf bytes.Buffer
	if g.comLen < 1 {
		decomBuf.Grow(512) // TODO::: 512 came from io.ReadAll, any suggestion??
	} else {
		decomBuf.Grow(g.comLen)
	}
	decomBuf.ReadFrom(&g.zr)
	g.decompressedData = decomBuf.Bytes()
}

/*
********** protocol.Codec interface **********
 */

func (g *gzipDecompressor) MediaType() protocol.MediaType       { return g.sourceMT }
func (g *gzipDecompressor) CompressType() protocol.CompressType { return nil }

func (g *gzipDecompressor) Decode(reader protocol.Reader) (err protocol.Error) {
	var goErr = g.zr.Reset(reader)
	if goErr != nil {
		// err =
	}
	return
}
func (g *gzipDecompressor) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = g.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (g *gzipDecompressor) Marshal() (data []byte) {
	if g.decompressedData == nil {
		g.decompressAll()
	}
	return g.decompressedData
}
func (g *gzipDecompressor) MarshalTo(data []byte) []byte {
	if g.decompressedData == nil {
		g.decompressAll()
	}
	return append(data, g.decompressedData...)
}
func (g *gzipDecompressor) Unmarshal(data []byte) (err protocol.Error) {
	g.comLen = len(data)
	g.comBuf = *bytes.NewBuffer(data)
	var goErr = g.zr.Reset(&g.comBuf)
	if goErr != nil {
		// err =
	}
	return
}
func (g *gzipDecompressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = g.Unmarshal(data)
	// TODO::: can return any remaining data?
	return
}

// Len return length of decompressed data
func (g *gzipDecompressor) Len() (ln int) {
	if g.decompressedData == nil {
		g.decompressAll()
	}
	return len(g.decompressedData)
}

func (g *gzipDecompressor) Discard() (err protocol.Error) {
	// io.Copy(io.Discard, b)
	return
}

/*
********** io package interfaces **********
 */

func (g *gzipDecompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = ErrSourceNotChangeable
	return
}
func (g *gzipDecompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	return io.Copy(w, &g.zr)
}
