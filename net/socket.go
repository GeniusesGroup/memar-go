/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
	"memar/timer"
)

type Socket struct {
	buf
	STATUS

	// timeout logic
	readTimer  timer.Async // read deadline timer
	writeTimer timer.Async // write deadline timer
}

//memar:impl memar/protocol.ObjectLifeCycle
func (sk *Socket) Init(timeout protocol.Duration) (err protocol.Error) {
	err = sk.buf.Init()
	if err != nil {
		return
	}
	err = sk.initTimeout()
	if err != nil {
		return
	}
	err = sk.SetTimeout(timeout)
	return
}
func (sk *Socket) Reinit() (err protocol.Error) {
	err = sk.reinitTimeout()
	return
}
func (sk *Socket) Deinit() (err protocol.Error) {
	err = sk.deinitTimeout()
	return
}

func (sk *Socket) Check() (err protocol.Error) {
	// TODO:::
	return
}
