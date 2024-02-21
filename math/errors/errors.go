/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
)

// Errors
var (
	ErrEmptyValue      er.Error
	ErrValueOutOfRange er.Error
	ErrBadValue        er.Error
)

func init() {
	ErrEmptyValue.Init("domain/memar.scm.geniuses.group; package=convert; type=error; name=empty-value")
	ErrValueOutOfRange.Init("domain/memar.scm.geniuses.group; package=convert; type=error; name=value-out-of-range")
	ErrBadValue.Init("domain/memar.scm.geniuses.group; package=convert; type=error; name=bad-value")
}
