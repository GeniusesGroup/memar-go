/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrTimerRacyAccess errTimerRacyAccess

type errTimerRacyAccess struct{ er.Err }

func (dt *errTimerRacyAccess) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=timer; type=error; name=timer-racy-access")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
