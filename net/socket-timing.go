/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
)

//memar:impl memar/protocol.ObjectLifeCycle
func (sk *Socket) initTimeout() (err protocol.Error) {
	err = sk.readTimer.Init(sk)
	if err != nil {
		return
	}
	err = sk.writeTimer.Init(sk)
	return
}
func (sk *Socket) reinitTimeout() (err protocol.Error) {
	err = sk.readTimer.Reinit(sk)
	if err != nil {
		return
	}
	err = sk.writeTimer.Reinit(sk)
	return
}
func (sk *Socket) deinitTimeout() (err protocol.Error) {
	err = sk.readTimer.Deinit()
	if err != nil {
		return
	}
	err = sk.writeTimer.Deinit()
	return
}

// TimerHandler or NotifyChannel does a non-blocking send the signal on sk.signal
func (sk *Socket) TimerHandler() {
	// TODO:::
}

//memar:impl memar/protocol.Timeout
func (sk *Socket) SetTimeout(d protocol.Duration) (err protocol.Error) {
	err = sk.SetReadTimeout(d)
	if err != nil {
		return
	}
	err = sk.SetWriteTimeout(d)
	return
}
func (sk *Socket) SetReadTimeout(d protocol.Duration) (err protocol.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	if d < 0 {
		// no timeout
		// TODO::: is it ok??
		err = sk.readTimer.Stop()
		return
	}
	sk.readTimer.Reset(d)
	return
}
func (sk *Socket) SetWriteTimeout(d protocol.Duration) (err protocol.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	if d < 0 {
		// no timeout
		err = sk.writeTimer.Stop()
		return
	}
	sk.writeTimer.Reset(d)
	return
}

//memar:impl memar/protocol.Deadline
func (sk *Socket) SetDeadline(d protocol.Time) (err protocol.Error) {

	return
}
