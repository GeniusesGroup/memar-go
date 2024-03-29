/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/buffer"
	"memar/protocol"
	"memar/time/monotonic"
	"memar/timer"
	// "memar/uuid/16byte"
)

type Socket struct {
	/* Connection data */
	weight protocol.Weight

	buf buffer.Queue

	// socketTiming
	socketTimer   timer.Async
	readDeadline  monotonic.Atomic
	writeDeadline monotonic.Atomic

	STATUS
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (sk *Socket) Init(timeout protocol.Duration) (err protocol.Error) {
	err = sk.initTimeout(timeout)
	return
}
func (sk *Socket) Reinit(timeout protocol.Duration) (err protocol.Error) {
	err = sk.reinitTimeout(timeout)
	return
}
func (sk *Socket) Deinit() (err protocol.Error) {
	err = sk.deinitTimeout()

	// first closing open listener for income frame and refuse all new frame,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
	return
}

// func (sk *Socket) Buffer() (buf protocol.Buffer)  { return &sk.buf }

func (sk *Socket) Weight() protocol.Weight { return sk.weight }

//libgo:impl libgo/protocol.NetworkAddress
// func (s *Socket) LocalAddr() protocol.Stringer  { return &s.localAddr }
// func (s *Socket) RemoteAddr() protocol.Stringer { return &s.remoteAddr }

//libgo:impl libgo/protocol.Session
func (sk *Socket) Close() (err protocol.Error)  { return }
func (sk *Socket) Revoke() (err protocol.Error) { return }

func (sk *Socket) Check() (err protocol.Error) {
	// TODO:::
	return
}

// ScheduleProcessingSocket is Non-Blocking means It must not block the caller in any ways.
// Stream must start with NetworkStatus_NeedMoreData if it doesn't need to call the service when the state changed for the first time
func (sk *Socket) ScheduleProcessingSocket() {
	// decide by stream odd or even
	// TODO::: check better performance as "streamID%2 == 0" to check odd id
	// if streamID&1 == 0 {
	// 	// TODO::: easily call by "go" or call by workers pool or what??
	// 	go f.callService(conn, stream)
	// } else {
	// 	// income response
	// 	stream.SetState(protocol.NetworkStatus_Ready)
	// }

	// if st.State == protocol.NetworkStatus_Open {
	// TODO::: easily call by "go" or call by workers pool or what??
	// go st.callService()
	// return
	// }
	// st.SetState(protocol.NetworkStatus_ReceivedCompletely)
}
