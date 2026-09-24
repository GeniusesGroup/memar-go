/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/process/errors"
)

var ErrPacketWrongLength errPacketWrongLength

type errPacketWrongLength struct{}

func (dt *errPacketWrongLength) Init() (err error_p.Error) {
	// err = dt.Err.Init("domain/udp.wg.ietf.org; type=error; name=packet-wrong-length")
	// if err != nil {
	// 	return
	// }
	// err = errors.Register(dt)
	return
}
