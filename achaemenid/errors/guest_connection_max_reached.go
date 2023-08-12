/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrGuestConnectionMaxReached errGuestConnectionMaxReached

type errGuestConnectionMaxReached struct{ er.Err }

func (dt *errGuestConnectionMaxReached) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "name=guest-connection-max-reached")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
