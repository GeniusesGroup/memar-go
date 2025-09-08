/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	error_p "memar/error/protocol"
	"memar/errors"
)

var ErrFlagNotFound errFlagNotFound

type errFlagNotFound struct{ er.Err }

func (dt *errFlagNotFound) Init() (err error_p.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=flag-not_found")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
