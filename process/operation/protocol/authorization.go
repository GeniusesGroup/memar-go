/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

import (
	error_p "memar/process/error/protocol"
)

// Authorize is part of operation process logic
type Authorize interface {
	authorize() error_p.Error
}
