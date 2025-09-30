/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
	"memar/protocol"
)

var ErrSourceNotChangeable errSourceNotChangeable

type errSourceNotChangeable struct{ er.Err }

func (dt *errSourceNotChangeable) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "source-not-changeable")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
