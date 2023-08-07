/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var (
	ErrProtocolHandler errProtocolHandler
)

type (
	errProtocolHandler struct{ er.Err }
)

func (dt *errProtocolHandler) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=connection; type=error; name=protocol-handler")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
