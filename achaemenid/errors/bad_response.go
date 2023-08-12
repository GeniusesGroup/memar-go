/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrBadResponse errBadResponse

type errBadResponse struct{ er.Err }

func (dt *errBadResponse) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "name=bad-response")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
