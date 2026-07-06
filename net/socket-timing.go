/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	error_p "memar/process/error/protocol"
	net_p "memar/net/protocol"
	"memar/time/duration"
	"memar/time/monotonic"
)

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (sk *Socket[BUF]) initTimeout(timeout duration.NanoSecond) (err error_p.Error) {
	err = sk.socketTimer.Init(sk)
	err = sk.socketTimer.Start(timeout)
	return
}
func (sk *Socket[BUF]) reinitTimeout(timeout duration.NanoSecond) (err error_p.Error) {
	err = sk.socketTimer.Reset(timeout)
	return
}
func (sk *Socket[BUF]) deinitTimeout() (err error_p.Error) {
	err = sk.socketTimer.Deinit()
	return
}

// Don't block the caller
func (sk *Socket[BUF]) TimerHandler() {
	var timerWhen = sk.socketTimer.When()

	if sk.readDeadline.Load() <= timerWhen {
		sk.SetStatus(net_p.Status_Timeout_Read)
	} else if sk.writeDeadline.Load() <= timerWhen {
		sk.SetStatus(net_p.Status_Timeout_Write)
	} else {
		// TODO::: Is it possible??
	}
}

//memar:impl memar/process/operation/protocol.Timeout
func (sk *Socket[BUF]) SetTimeout(d duration.NanoSecond) (err error_p.Error) {
	err = sk.SetReadTimeout(d)
	if err != nil {
		return
	}
	err = sk.SetWriteTimeout(d)
	return
}
func (sk *Socket[BUF]) SetReadTimeout(d duration.NanoSecond) (err error_p.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	err = sk.setWriteTimeout(d)
	return
}
func (sk *Socket[BUF]) SetWriteTimeout(d duration.NanoSecond) (err error_p.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	err = sk.setReadTimeout(d)
	return
}

func (sk *Socket[BUF]) setReadTimeout(d duration.NanoSecond) (err error_p.Error) {
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
func (sk *Socket[BUF]) setWriteTimeout(d duration.NanoSecond) (err error_p.Error) {
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
