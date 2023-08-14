/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var ErrNotExist errNotExist

type errNotExist struct{ er.Err }

func (dt *errNotExist) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=errors; type=error; name=not_exist")
	if err != nil {
		return
	}
	err = Register(dt)
	return
}
