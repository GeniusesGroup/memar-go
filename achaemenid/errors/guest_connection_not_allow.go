/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
	"memar/protocol"
)

var ErrGuestConnectionNotAllow errGuestConnectionNotAllow

type errGuestConnectionNotAllow struct{ er.Err }

func (dt *errGuestConnectionNotAllow) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "name=guest-connection-not-allow")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
