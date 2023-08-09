/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrTimerNotInit errTimerNotInit

type errTimerNotInit struct{ er.Err }

func (dt *errTimerNotInit) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=timer; type=error; name=timer-not-initialize")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
