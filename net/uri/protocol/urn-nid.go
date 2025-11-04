/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_URN_NID interface {
	NID() URN_NID
}

// NID is the namespace identifier e.g. "isbn"
type URN_NID interface {
	string_p.String
}
