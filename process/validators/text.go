/* For license and copyright information please see the LEGAL file in the code repository */

package validators

import (
	error_p "memar/process/error/protocol"
)

func ValidateTextLength(t string, min, max int) (err error_p.Error) {
	if len(t) < min {
		return &TextLack
	}
	if max != 0 && len(t) > max {
		return &TextOverFlow
	}
	return
}
