/* For license and copyright information please see the LEGAL file in the code repository */

package compress_p

import (
	buffer_p "memar/computer/buffer/protocol"
	datatype_p "memar/computer/datatype/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	error_p "memar/process/error/protocol"
	file_p "memar/storage/file/protocol"
)

// CompressType is standard shape of any compress coding type
type CompressType interface {
	Compress(raw buffer_p.Buffer, options CompressOptions) (compressed buffer_p.Buffer, err error_p.Error)
	Decompress(compressed buffer_p.Buffer) (raw buffer_p.Buffer, err error_p.Error)

	Field_ContentEncoding
	datatype_p.DataType
	mediatype_p.Field_MediaType
	file_p.Field_Extension
}

type CompressOptions interface {
	Field_CompressLevel

	// Use source buffer and overwrite or make new buffer
	Overwrite() bool
}
