/* For license and copyright information please see the LEGAL file in the code repository */

package flate

import (
	"memar/datatype"
	"memar/mediatype"
	"memar/protocol"
)

const (
	ContentEncoding = "deflate"
	Extension       = "zz"
)

var Deflate deflate

type deflate struct {
	datatype.DataType
	mediatype.MT
}

//memar:impl memar/protocol.ObjectLifeCycle
func (d *deflate) Init() (err protocol.Error) {
	err = d.MT.Init("application/deflate")
	return
}

//memar:impl memar/protocol.DataType_Details
func (d *deflate) Status() protocol.SoftwareStatus    { return protocol.Software_StableRelease }
func (d *deflate) ReferenceURI() string               { return "" }
func (d *deflate) IssueDate() protocol.Time           { return nil }
func (d *deflate) ExpiryDate() protocol.Time          { return nil }
func (d *deflate) ExpireInFavorOf() protocol.DataType { return nil }

//memar:impl memar/protocol.MediaType
func (d *deflate) FileExtension() string { return Extension }

//memar:impl memar/protocol.CompressType
func (d *deflate) ContentEncoding() string { return ContentEncoding }
func (d *deflate) Compress(raw protocol.Codec, options protocol.CompressOptions) (compressed protocol.Codec, err protocol.Error) {
	var com Compressor
	err = com.Init(raw, options)
	compressed = &com
	return
}
func (d *deflate) Decompress(compressed protocol.Codec) (raw protocol.Codec, err protocol.Error) {
	var dec Decompressor
	err = dec.Init(compressed)
	raw = &dec
	return
}
