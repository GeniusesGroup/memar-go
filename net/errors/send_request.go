/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var (
	ErrSendRequest errSendRequest
)

type (
	errSendRequest struct{ er.Err }
)

func (dt *errSendRequest) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=connection; type=error; name=send-request")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
