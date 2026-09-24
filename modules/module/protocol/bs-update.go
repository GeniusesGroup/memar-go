/* For license and copyright information please see the LEGAL file in the code repository */

package mp

import (
	uri_p "memar/net/uri/protocol"
)

type (
	BS_Update_Request interface {
		Name() DomainName
		uri_p.Field_Path
		// Transport() []string
	}
)
