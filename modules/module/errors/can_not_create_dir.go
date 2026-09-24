/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/process/errors"
)

var ErrCanNotCreateDir errCanNotCreateDir

type errCanNotCreateDir struct{ }

func (d *errCanNotCreateDir) Init() {
	// d.Err.Init("domain/memar.scm.geniuses.group; module=modules/module; type=error; name=can-not-create-dir")
	// errors.Register(d)
}
