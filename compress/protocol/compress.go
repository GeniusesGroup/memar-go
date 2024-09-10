/* For license and copyright information please see the LEGAL file in the code repository */

package compress_p

import (
	buffer_p "memar/buffer/protocol"
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
	mediatype_p "memar/mediatype/protocol"
	storage_p "memar/storage/protocol"
)

// CompressType is standard shape of any compress coding type
type CompressType[BUF buffer_p.Buffer, OPTs any] interface {
	Compress(raw BUF, options OPTs) (compressed BUF, err error_p.Error)
	Decompress(compressed BUF) (raw BUF, err error_p.Error)

	ContentEncoding
	datatype_p.DataType
	mediatype_p.MediaType
	storage_p.FileExtension
}

// ContentEncoding is standard shape of http compress coding type string
type ContentEncoding interface {
	// https://en.wikipedia.org/wiki/HTTP_compression
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
	// https://www.iana.org/assignments/http-parameters/http-parameters.xml#content-coding
	ContentEncoding() string
}
