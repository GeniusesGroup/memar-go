/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var ErrPacketWrongLength errPacketWrongLength

type errPacketWrongLength struct{ er.Err }

func (dt *errPacketWrongLength) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/udp.wg.ietf.org; type=error; name=packet-wrong-length")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}
