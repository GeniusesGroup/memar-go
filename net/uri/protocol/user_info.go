/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

// URI Userinfo >> "username[:password]"
type Field_UserInfo interface {
	Field_Username
	Field_Password
}
