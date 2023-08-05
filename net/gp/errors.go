/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	er "memar/error"
)

// Errors
var (
	ErrFrameLength           er.Error
	ErrBadFrameType          er.Error
	ErrFrameArrivedAnterior  er.Error
	ErrFrameArrivedPosterior er.Error
)

func init() {
	ErrFrameLength.Init("domain/gp.scm.geniuses.group; type=error; name=frame-length")
	ErrBadFrameType.Init("domain/gp.scm.geniuses.group; type=error; name=bad-frame-type")
	ErrFrameArrivedAnterior.Init("domain/gp.scm.geniuses.group; type=error; name=frame-arrived-anterior")
	ErrFrameArrivedPosterior.Init("domain/gp.scm.geniuses.group; type=error; name=frame-arrived-posterior")
}
