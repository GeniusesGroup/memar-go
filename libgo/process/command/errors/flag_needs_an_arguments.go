/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
)

var ErrFlagNeedsAnArgument errFlagNeedsAnArgument

type errFlagNeedsAnArgument struct{ er.Err }

func (dt *errFlagNeedsAnArgument) Init() (err error_p.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=flag_needs_an_arguments")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
