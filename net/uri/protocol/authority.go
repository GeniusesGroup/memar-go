/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

// URI Authority >> [ userinfo "@" ] host [ ":" port ]
type Field_Authority interface {
	Field_UserInfo
	Field_Host
	Field_Port
}
