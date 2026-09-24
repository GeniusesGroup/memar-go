/* For license and copyright information please see the LEGAL file in the code repository */

package mp

import (
	uri_p "memar/net/uri/protocol"
)

// memar module new domain-name=node path="./modules/"
type (
	BS_New_Request interface {
		Name() DomainName
		uri_p.Field_Path
	}

	BS_New_Response interface {
		Dirs() []string
		FilesPath() []string // path with file name start from given path
		FilesData() [][]byte // files data
	}
)
