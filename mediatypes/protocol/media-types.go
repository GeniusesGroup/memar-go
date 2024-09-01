/* For license and copyright information please see the LEGAL file in the code repository */

package mts_p

import (
	"memar/protocol"
	string_p "memar/string/protocol"
)

type MediaTypes interface {
	Register(mt protocol.MediaType) (err protocol.Error)
	GetByMediaType(mediaType string_p.String) (mt protocol.MediaType, err protocol.Error)
	GetByFileExtension(ex string_p.String) (mt protocol.MediaType, err protocol.Error)
}
