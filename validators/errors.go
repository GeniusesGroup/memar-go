/* For license and copyright information please see the LEGAL file in the code repository */

package validators

import (
	er "libgo/error"
)

// Declare package errors
var (
	ErrTextOverFlow         er.Error
	ErrTextLack             er.Error
	ErrTextIllegalCharacter er.Error
)

func init() {
	ErrTextOverFlow.Init("domain/libgo.scm.geniuses.group; package=validators; type=error; name=text-over-flow")
	ErrTextLack.Init("domain/libgo.scm.geniuses.group; package=validators; type=error; name=text-lack")
	ErrTextIllegalCharacter.Init("domain/libgo.scm.geniuses.group; package=validators; type=error; name=text-illegal-character")
}
