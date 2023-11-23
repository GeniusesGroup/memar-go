/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
	"memar/protocol"
)

var ErrInternalError errInternalError

type errInternalError struct{ er.Err }

func (dt *errInternalError) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "name=internal-error")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
