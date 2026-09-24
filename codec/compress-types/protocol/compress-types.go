/* For license and copyright information please see the LEGAL file in the code repository */

package cts_p

import (
	compress_p "memar/codec/compress/protocol"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

type CompressTypes interface {
	Register(ct compress_p.CompressType) (err error_p.Error)
	GetByID(id datatype_p.ID) (ct compress_p.CompressType, err error_p.Error)
	GetByMediaType(mt string) (ct compress_p.CompressType, err error_p.Error)
	GetByFileExtension(ex string) (ct compress_p.CompressType, err error_p.Error)
	GetByContentEncoding(ce string) (ct compress_p.CompressType, err error_p.Error)

	ContentEncodings() []string
}
