/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_Version /*[STR string_p.String]*/ interface {
	Version() string_p.String
}
