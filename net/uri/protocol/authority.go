/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/string/protocol"
)

type Username[STR string_p.String] interface {
	Username() STR
}

type Password[STR string_p.String] interface {
	Password() STR
}

type Host[STR string_p.String] interface {
	Host() STR
}

type Port[STR string_p.String] interface {
	Port() STR
}