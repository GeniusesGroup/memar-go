/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var (
	ErrNotFound errNotFound
)

type errNotFound struct{ er.Err }

func (dt *errNotFound) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=errors; type=error; name=not_found")
	if err != nil {
		return
	}
	err = Register(dt)
	return
}
