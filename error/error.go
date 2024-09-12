/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	error_p "memar/error/protocol"
)

// IsEqual compare two Error.
// Suggest to add more logic to Error.Equal() logic such as chain situation!
func IsEqual(base, with error_p.Error) bool {
	if base == nil && with == nil {
		return true
	}
	if base != nil &&
		with != nil &&
		base.MediaType() == with.MediaType() &&
		base.DataTypeID() == with.DataTypeID() {
		return true
	}
	return false
}
