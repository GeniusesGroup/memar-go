/* For license and copyright information please see the LEGAL file in the code repository */

package gzip

import (
	buffer_p "memar/computer/buffer/protocol"
	"memar/codec/compress"
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

var GZIP gzip

type gzip struct {
	datatype.DataType
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (g *gzip) Init() (err error_p.Error) {
	return
}

//memar:impl memar/datatype/protocol.DataType_Details
func (g *gzip) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_StableRelease }
func (g *gzip) ReferenceURI() string                 { return "" }
func (g *gzip) IssueDate() string                    { return "" }
func (g *gzip) ExpiryDate() string                   { return "" }
func (g *gzip) ExpireInFavorOf() datatype_p.DataType { return nil }

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (g *gzip) MediaType() string { return "application/gzip" }

//memar:impl memar/storage/protocol.FileExtension
func (g *gzip) FileExtension() string { return Extension }

//memar:impl memar/compress_p.CompressType
func (g *gzip) ContentEncoding() string { return ContentEncoding }
func (g *gzip) Compress(raw buffer_p.Buffer, options compress.Options) (compressed buffer_p.Buffer, err error_p.Error) {
	var com Compressor
	err = com.Init(raw, options)
	compressed = &com
	return
}
func (g *gzip) Decompress(compressed buffer_p.Buffer) (raw buffer_p.Buffer, err error_p.Error) {
	var dec Decompressor
	err = dec.Init(compressed)
	raw = &dec
	return
}
