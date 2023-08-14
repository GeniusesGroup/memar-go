/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var ErrNotProvideIdentifier errNotProvideIdentifier

type errNotProvideIdentifier struct{ er.Err }

func (dt *errNotProvideIdentifier) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=errors; type=error; name=not_provide_identifier")
	if err != nil {
		return
	}
	err = Register(dt)
	return
}
