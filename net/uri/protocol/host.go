/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_Host interface {
	Host() Host
}

type Host interface {
	string_p.String
}
