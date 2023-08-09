/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrNegativeDuration errNegativeDuration

type errNegativeDuration struct{ er.Err }

func (dt *errNegativeDuration) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=timer; type=error; name=negative-duration")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
