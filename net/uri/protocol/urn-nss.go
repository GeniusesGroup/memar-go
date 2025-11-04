/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

import (
	string_p "memar/codec/string/protocol"
)

type Field_URN_NSS interface {
	NSS() URN_NSS
}

// NSS is the namespace-specific e.g. "0451450523"
type URN_NSS interface {
	string_p.String
}
