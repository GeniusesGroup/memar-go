/* For license and copyright information please see the LEGAL file in the code repository */

package mts_p

import (
	error_p "memar/error/protocol"
	mediatype_p "memar/mediatype/protocol"
	string_p "memar/string/protocol"
)

type MediaTypes interface {
	Register(mt mediatype_p.MediaType) (err error_p.Error)
	GetByMediaType(mediaType string_p.String) (mt mediatype_p.MediaType, err error_p.Error)
	GetByFileExtension(ex string_p.String) (mt mediatype_p.MediaType, err error_p.Error)
}
