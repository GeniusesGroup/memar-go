/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var ErrPacketTooShort errPacketTooShort

type errPacketTooShort struct{ er.Err }

func (dt *errPacketTooShort) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/udp.wg.ietf.org; type=error; name=packet-too-short")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
