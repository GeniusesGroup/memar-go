/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
)

var ErrServiceNotAcceptCLI errServiceNotAcceptCLI

type errServiceNotAcceptCLI struct{ er.Err }

func (dt *errServiceNotAcceptCLI) Init() (err error_p.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=service-not-accept-cli")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
