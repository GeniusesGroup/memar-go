/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrNotStandardID errNotStandardID

type errNotStandardID struct{ er.Err }

func (dt *errNotStandardID) Init() (err protocol.Error) {
	err = dt.Err.Init(domainBaseMediatype + "name=not-standard-id")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
