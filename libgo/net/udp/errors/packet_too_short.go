/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/process/errors"
)

var ErrPacketTooShort errPacketTooShort

type errPacketTooShort struct{ }

func (dt *errPacketTooShort) Init() (err error_p.Error) {
	// err = dt.Err.Init("domain/udp.wg.ietf.org; type=error; name=packet-too-short")
	// if err != nil {
	// 	return
	// }
	// err = errors.Register(dt)
	return
}
