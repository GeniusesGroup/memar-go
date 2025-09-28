/* For license and copyright information please see the LEGAL file in the code repository */

package validators

import (
	"libgo/protocol"
)

func ValidateTextLength(t string, min, max int) (err protocol.Error) {
	if len(t) < min {
		return &ErrTextLack
	}
	if max != 0 && len(t) > max {
		return &ErrTextOverFlow
	}
	return
}
