/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_Password interface {
	Password() Password
}

type Password interface {
	string_p.String
}
