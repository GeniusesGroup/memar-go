/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_Method /*[STR string_p.String]*/ interface {
	Method() string_p.String
}
