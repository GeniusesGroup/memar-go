/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
	"memar/time/monotonic"
	"memar/timer"
	// "memar/uuid/16byte"
)

type Socket struct {
	/* Connection data */
	weight protocol.Weight

	// socketTiming
	socketTimer   timer.Async
	readDeadline  monotonic.Atomic
	writeDeadline monotonic.Atomic

	STATUS
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (sk *Socket) Init(timeout protocol.Duration) (err protocol.Error) {

	err = sk.socketTimer.Init(sk)
	err = sk.socketTimer.Start(timeout)

	return
}
func (sk *Socket) Reinit() (err protocol.Error) {
	err = sk.socketTimer.Reinit(sk)
	return
}
func (sk *Socket) Deinit() (err protocol.Error) {
	err = sk.socketTimer.Deinit()

	// first closing open listener for income frame and refuse all new frame,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
	return
}

func (sk *Socket) Weight() protocol.Weight { return sk.weight }

//libgo:impl libgo/protocol.NetworkAddress
// func (s *Socket) LocalAddr() protocol.Stringer  { return &s.localAddr }
// func (s *Socket) RemoteAddr() protocol.Stringer { return &s.remoteAddr }

//libgo:impl libgo/protocol.Session
func (sk *Socket) Close() (err protocol.Error)  { return }
func (sk *Socket) Revoke() (err protocol.Error) { return }

func (sk *Socket) Check() (err protocol.Error) { 
	// TODO::: 
	return }
