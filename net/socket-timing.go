/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/time/monotonic"
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

// Don't block the caller
func (sk *Socket) TimerHandler() {
	var timerWhen = sk.socketTimer.When()

	if sk.readDeadline.Load() <= timerWhen {
		sk.SetStatus(protocol.NetworkStatus_Timeout_Read)
	} else if sk.writeDeadline.Load() <= timerWhen {
		sk.SetStatus(protocol.NetworkStatus_Timeout_Write)
	} else {
		// TODO::: Is it possible??
	}
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

	err = sk.setWriteTimeout(d)
	return
}
func (sk *Socket) SetWriteTimeout(d protocol.Duration) (err protocol.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	err = sk.setReadTimeout(d)
	return
}

func (sk *Socket) setReadTimeout(d protocol.Duration) (err protocol.Error) {
	if d < 0 {
		// no timeout
		sk.readDeadline.Store(0)
		if sk.writeDeadline.Load() == 0 {
			sk.socketTimer.Stop()
		}
		return
	}

	var readDeadline = monotonic.Now()
	readDeadline.Add(d)
	sk.readDeadline.Store(readDeadline)

	if readDeadline < sk.socketTimer.When() {
		sk.socketTimer.Reset(d)
	}
	return
}
func (sk *Socket) setWriteTimeout(d protocol.Duration) (err protocol.Error) {
	if d < 0 {
		// no timeout
		sk.writeDeadline.Store(0)
		if sk.readDeadline.Load() == 0 {
			sk.socketTimer.Stop()
		}
	}

	var writeDeadline = monotonic.Now()
	writeDeadline.Add(d)
	sk.writeDeadline.Store(writeDeadline)

	if writeDeadline < sk.socketTimer.When() {
		sk.socketTimer.Reset(d)
	}
	return
}
