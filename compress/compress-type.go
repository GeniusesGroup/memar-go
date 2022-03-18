/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	"../mediatype"
	"../protocol"
)

type CompressType struct {
	mediatype       *mediatype.MediaType
	contentEncoding string
	extension       string
}

func (ct *CompressType) MediaType() protocol.MediaType { return ct.mediatype }
func (ct *CompressType) ContentEncoding() string       { return ct.contentEncoding }
func (ct *CompressType) FileExtension() string         { return ct.extension }

func New(contentEncoding string, mediatype *mediatype.MediaType) (ct *CompressType) {
	if mediatype == nil {
		panic("CompressType doesn't has a valid MediaType. Can't make it.")
	}
	if contentEncoding == "" {
		panic("CompressType doesn't has a valid ContentEncoding. Can't make it.")
	}
	ct = &CompressType{
		mediatype:       mediatype,
		contentEncoding: contentEncoding,
		extension:       mediatype.FileExtension(),
	}
	return
}
