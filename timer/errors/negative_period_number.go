/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/protocol"
)

var ErrNegativePeriodNumber errNegativePeriodNumber

type errNegativePeriodNumber struct{ er.Err }

func (dt *errNegativePeriodNumber) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=timer; type=error; name=negative-period-number")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
