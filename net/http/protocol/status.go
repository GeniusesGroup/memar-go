/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	string_p "memar/string/protocol"
)

type Status[STR string_p.String] interface {
	StatusCode() STR
	ReasonPhrase() STR

	SetStatus(statusCode, reasonPhrase STR)
}
