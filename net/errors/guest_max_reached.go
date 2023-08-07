/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var (
	ErrGuestMaxReached errGuestMaxReached
)

type (
	errGuestMaxReached struct{ er.Err }
)

func (dt *errGuestMaxReached) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=connection; type=error; name=guest-max-reached")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
