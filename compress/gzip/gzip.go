/* For license and copyright information please see the LEGAL file in the code repository */

package gzip

import (
	"memar/datatype"
	"memar/mediatype"
	"memar/protocol"
)

const (
	ContentEncoding = "gzip"
	Extension       = "gzip"
)

var GZIP gzip

type gzip struct {
	datatype.DataType
	mediatype.MT
}

//memar:impl memar/protocol.ObjectLifeCycle
func (g *gzip) Init() (err protocol.Error) {
	err = g.MT.Init("application/gzip")
	return
}

//memar:impl memar/protocol.DataType_Details
func (g *gzip) Status() protocol.SoftwareStatus    { return protocol.Software_StableRelease }
func (g *gzip) ReferenceURI() string               { return "" }
func (g *gzip) IssueDate() protocol.Time           { return nil }
func (g *gzip) ExpiryDate() protocol.Time          { return nil }
func (g *gzip) ExpireInFavorOf() protocol.DataType { return nil }

//memar:impl memar/protocol.MediaType
func (g *gzip) FileExtension() string { return Extension }

//memar:impl memar/protocol.CompressType
func (g *gzip) ContentEncoding() string { return ContentEncoding }
func (g *gzip) Compress(raw protocol.Codec, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	var com Compressor
	err = com.Init(raw, options)
	compressed = &com
	return
}
func (g *gzip) Decompress(compressed protocol.Codec) (raw protocol.Codec, err protocol.Error) {
	var dec Decompressor
	err = dec.Init(compressed)
	raw = &dec
	return
}
