/* For license and copyright information please see the LEGAL file in the code repository */

package cts_p

import (
	error_p "memar/error/protocol"
)

type CompressTypes interface {
	Register(ct CompressType) (err error_p.Error)
	GetByID(id ID) (ct CompressType, err error_p.Error)
	GetByMediaType(mt string) (ct CompressType, err error_p.Error)
	GetByFileExtension(ex string) (ct CompressType, err error_p.Error)
	GetByContentEncoding(ce string) (ct CompressType, err error_p.Error)

	ContentEncodings() []string
}
