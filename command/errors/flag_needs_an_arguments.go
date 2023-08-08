/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrFlagNeedsAnArgument errFlagNeedsAnArgument

type errFlagNeedsAnArgument struct{ er.Err }

func (dt *errFlagNeedsAnArgument) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=flag_needs_an_arguments")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
