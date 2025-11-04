/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_Username interface {
	Username() Username
}

type Username interface {
	string_p.String
}
