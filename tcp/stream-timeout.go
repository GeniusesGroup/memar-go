/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/protocol"
)

//libgo:impl libgo/protocol.Timeout
func (s *Stream) SetTimeout(d protocol.Duration) (err protocol.Error) {
	err = s.SetReadTimeout(d)
	if err != nil {
		return
	}
	err = s.SetWriteTimeout(d)
	return
}
func (s *Stream) SetReadTimeout(d protocol.Duration) (err protocol.Error) {
	err = s.checkStream()
	if err != nil {
		return
	}

	if d < 0 {
		// no timeout
		// TODO::: is it ok??
		s.recv.readTimer.Stop()
		return
	}
	s.recv.readTimer.Reset(d)
	return
}
func (s *Stream) SetWriteTimeout(d protocol.Duration) (err protocol.Error) {
	err = s.checkStream()
	if err != nil {
		return
	}

	if d < 0 {
		// no timeout
		s.send.writeTimer.Stop()
		return
	}
	s.send.writeTimer.Reset(d)
	return
}
