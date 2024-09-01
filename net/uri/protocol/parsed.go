/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/string/protocol"
)

type Parsed[STR string_p.String] interface {
	Scheme[STR]

	// URI Authority >> [ userinfo "@" ] host [ ":" port ]
	// URI Userinfo >> "username[:password]"
	Username[STR]
	Password[STR]
	Host[STR]
	Port[STR]

	Path[STR]
	Query[STR]
	Fragment[STR]
}
