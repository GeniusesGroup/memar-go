/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
	"memar/protocol"
)

var ErrBadRequest errBadRequest

type errBadRequest struct{ er.Err }

func (dt *errBadRequest) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "name=bad-request")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
