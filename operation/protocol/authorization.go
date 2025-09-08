/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

import (
	datatype_p "memar/datatype/protocol"
)

// Authorize ...
type Authorize interface {
	authorize() error_p.Error
}
