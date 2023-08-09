/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var (
	ErrNoConnection errNoConnection
)

type (
	errNoConnection struct{ er.Err }
)

func (dt *errNoConnection) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=net; type=error; name=no-connection")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
