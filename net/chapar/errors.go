/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	er "libgo/error"
)

// Package errors
var (
	ErrShortFrameLength er.Error
	ErrLongFrameLength  er.Error
	ErrBadFrameID       er.Error
	ErrMTU              er.Error
	ErrPortNotExist     er.Error
	ErrPathAlreadyUse   er.Error
	ErrPathAlreadyExist er.Error
	ErrNotAcceptLastHop er.Error
)

func init() {
	ErrShortFrameLength.Init("domain/chapar.scm.geniuses.group; type=error; name=short-frame-length")
	ErrLongFrameLength.Init("domain/chapar.scm.geniuses.group; type=error; name=long-frame-length")
	ErrBadFrameID.Init("domain/chapar.scm.geniuses.group; type=error; name=bad-frame-id")
	ErrMTU.Init("domain/chapar.scm.geniuses.group; type=error; name=maximum-transmission-unit")
	ErrPortNotExist.Init("domain/chapar.scm.geniuses.group; type=error; name=port-not-exist")
	ErrPathAlreadyUse.Init("domain/chapar.scm.geniuses.group; type=error; name=path-already-use")
	ErrPathAlreadyExist.Init("domain/chapar.scm.geniuses.group; type=error; name=path-already-exist")
	ErrNotAcceptLastHop.Init("domain/chapar.scm.geniuses.group; type=error; name=not-accept-last-hop")
}
