/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
	"memar/protocol"
)

var ErrServiceCallByAlias errServiceCallByAlias

type errServiceCallByAlias struct{ er.Err }

func (dt *errServiceCallByAlias) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=service-call-by-alias")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
