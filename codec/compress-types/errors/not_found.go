/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
	"memar/protocol"
)

var ErrNotFound errNotFound

type errNotFound struct{ er.Err }

func (dt *errNotFound) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "not-found")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
