/* For license and copyright information please see the LEGAL file in the code repository */

package flate

import (
	codec_p "memar/codec/protocol"
	"memar/codec/compress"
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

const (
	ContentEncoding = "deflate"
	Extension       = "zz"
)

var Deflate deflate

type deflate struct {
	datatype.DataType
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (d *deflate) MediaType() string { return "application/deflate" }

//memar:impl memar/datatype/protocol.DataType_Details
func (d *deflate) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_StableRelease }
func (d *deflate) ReferenceURI() string                 { return "" }
func (d *deflate) IssueDate() string                    { return "" }
func (d *deflate) ExpiryDate() string                   { return "" }
func (d *deflate) ExpireInFavorOf() datatype_p.DataType { return nil }

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (d *deflate) FileExtension() string { return Extension }

//memar:impl memar/compress_p.CompressType
func (d *deflate) ContentEncoding() string { return ContentEncoding }
func (d *deflate) Compress(raw codec_p.Codec, options compress.Options) (compressed codec_p.Codec, err error_p.Error) {
	var com Compressor
	err = com.Init(raw, options)
	compressed = &com
	return
}
func (d *deflate) Decompress(compressed codec_p.Codec) (raw codec_p.Codec, err error_p.Error) {
	var dec Decompressor
	err = dec.Init(compressed)
	raw = &dec
	return
}
