/* For license and copyright information please see the LEGAL file in the code repository */

package validation_p

import (
	error_p "memar/error/protocol"
)

// **ATTENTION**::: strongly suggest use primitive_p.Accessor to prevent invalid state at first place.
type Validation interface {
	Validate() (err error_p.Error)
}
