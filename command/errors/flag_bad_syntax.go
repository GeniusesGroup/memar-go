/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
	"memar/errors"
	"memar/protocol"
)

var ErrFlagBadSyntax errFlagBadSyntax

type errFlagBadSyntax struct{ er.Err }

func (dt *errFlagBadSyntax) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=flag_bad_syntax")
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}
