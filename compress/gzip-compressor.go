/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"bytes"
	"compress/gzip"
	"io"

	"../protocol"
)

type gzipCompressor struct {
	source         protocol.Codec
	compressLevel  protocol.CompressLevel
	compressedData []byte
}

func (g *gzipCompressor) init() {
	var b bytes.Buffer
	b.Grow(g.source.Len())
	var gz, _ = gzip.NewWriterLevel(&b, int(g.compressLevel))
	gz.Write(g.source.Marshal())
	gz.Close()
	g.compressedData = b.Bytes()
}

/*
********** protocol.Codec interface **********
 */

func (g *gzipCompressor) MediaType() protocol.MediaType       { return g.source.MediaType() }
func (g *gzipCompressor) CompressType() protocol.CompressType { return GZIP }

func (g *gzipCompressor) Decode(reader protocol.Reader) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}
func (g *gzipCompressor) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = g.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (g *gzipCompressor) Marshal() (data []byte) {
	if g.compressedData == nil {
		g.init()
	}
	return g.compressedData
}
func (g *gzipCompressor) MarshalTo(data []byte) []byte {
	if g.compressedData == nil {
		g.init()
	}
	return append(data, g.compressedData...)
}
func (g *gzipCompressor) Unmarshal(data []byte) (err protocol.Error) {
	err = ErrSourceNotChangeable
	return
}
func (g *gzipCompressor) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) { 
	err = ErrSourceNotChangeable
	return 
}

// Len return length of compressed data
func (g *gzipCompressor) Len() (ln int) {
	if g.compressedData == nil {
		g.init()
	}
	return len(g.compressedData)
}

/*
********** io package interfaces **********
 */

func (g *gzipCompressor) ReadFrom(reader io.Reader) (n int64, err error) {
	err = ErrSourceNotChangeable
	return
}
func (g *gzipCompressor) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var gz, _ = gzip.NewWriterLevel(w, int(g.compressLevel))
	var writeLen int
	writeLen, err = gz.Write(g.source.Marshal())
	gz.Close()
	totalWrite = int64(writeLen)
	return
}
