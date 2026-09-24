/* For license and copyright information please see the LEGAL file in the code repository */

package command_errs

import (
	error_p "memar/process/error/protocol"
)

var ErrServiceNotAcceptCLI errServiceNotAcceptCLI

type errServiceNotAcceptCLI struct{}

func (dt *errServiceNotAcceptCLI) Init() (err error_p.Error) {
	// err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=service-not-accept-cli")
	// if err != nil {
	// 	return
	// }
	// err = errors.Register(dt)
	return
}
