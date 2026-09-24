/* For license and copyright information please see the LEGAL file in the code repository */

package validation_p

import (
	error_p "memar/process/error/protocol"
)

// **ATTENTION**::: strongly suggest use capsule_p.Assignment to prevent invalid state at first place.
type Validation interface {
	Validate() (err error_p.Error)
}
