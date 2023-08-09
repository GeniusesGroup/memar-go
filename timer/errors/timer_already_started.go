/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrTimerAlreadyStarted errTimerAlreadyStarted

type errTimerAlreadyStarted struct{ er.Err }

func (dt *errTimerAlreadyStarted) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=timer; type=error; name=timer-already-started")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
