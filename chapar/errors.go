/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	er "libgo/error"
)

// Package errors
var (
	ErrShortFrameLength er.Error
	ErrLongFrameLength  er.Error
	ErrMTU              er.Error
	ErrPortNotExist     er.Error
	ErrPathAlreadyUse   er.Error
	ErrPathAlreadyExist er.Error
	ErrNotAcceptLastHop er.Error
)

func init() {
	ErrShortFrameLength.Init("domain/chapar.protocol; type=error; name=short-frame-length")
	ErrLongFrameLength.Init("domain/chapar.protocol; type=error; name=long-frame-length")
	ErrMTU.Init("domain/chapar.protocol; type=error; name=maximum-transmission-unit")
	ErrPortNotExist.Init("domain/chapar.protocol; type=error; name=port-not-exist")
	ErrPathAlreadyUse.Init("domain/chapar.protocol; type=error; name=path-already-use")
	ErrPathAlreadyExist.Init("domain/chapar.protocol; type=error; name=path-already-exist")
	ErrNotAcceptLastHop.Init("domain/chapar.protocol; type=error; name=not-accept-last-hop")
}
