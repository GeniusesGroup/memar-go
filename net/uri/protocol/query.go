/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/string/protocol"
)

type Query[STR string_p.String] interface {
	Query() STR
}
