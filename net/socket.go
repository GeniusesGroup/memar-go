/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
	operation_p "memar/process/operation/protocol"
	"memar/time/duration"
	"memar/time/monotonic"
	"memar/time/timer"
	// "memar/uuid/16byte"
)

type Socket[BUF buffer_p.Buffer] struct {
	/* Connection data */
	weight operation_p.Weight

	sendBuffer    BUF
	receiveBuffer BUF

	// socketTiming
	socketTimer   timer.Async
	readDeadline  monotonic.Atomic
	writeDeadline monotonic.Atomic

	STATUS
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (sk *Socket[BUF]) Init(timeout duration.NanoSecond) (err error_p.Error) {
	err = sk.initTimeout(timeout)
	return
}
func (sk *Socket[BUF]) Reinit(timeout duration.NanoSecond) (err error_p.Error) {
	err = sk.reinitTimeout(timeout)
	return
}
func (sk *Socket[BUF]) Deinit() (err error_p.Error) {
	err = sk.deinitTimeout()

	// first closing open listener for income frame and refuse all new frame,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
	return
}

//memar:impl memar/protocol.Socket_Buffer
func (sk *Socket[BUF]) SendBuffer() buffer_p.Buffer    { return sk.sendBuffer }
func (sk *Socket[BUF]) ReceiveBuffer() buffer_p.Buffer { return sk.receiveBuffer }

func (sk *Socket[BUF]) Weight() operation_p.Weight { return sk.weight }

//memar:impl memar/net/protocol.Field_NetworkAddresses
// func (s *Socket) LocalAddr() net_p.NetworkAddress  { return &s.localAddr }
// func (s *Socket) RemoteAddr() net_p.NetworkAddress { return &s.remoteAddr }

//memar:impl memar/protocol.Session
func (sk *Socket[BUF]) Close() (err error_p.Error)  { return }
func (sk *Socket[BUF]) Revoke() (err error_p.Error) { return }

func (sk *Socket[BUF]) Check() (err error_p.Error) {
	// TODO:::
	return
}

// ScheduleProcessingSocket is Non-Blocking means It must not block the caller in any ways.
// Stream must start with NetworkStatus_NeedMoreData if it doesn't need to call the service when the state changed for the first time
func (sk *Socket[BUF]) ScheduleProcessingSocket() {
	// decide by stream odd or even
	// TODO::: check better performance as "streamID%2 == 0" to check odd id
	// if streamID&1 == 0 {
	// 	// TODO::: easily call by "go" or call by workers pool or what??
	// 	go f.callService(conn, stream)
	// } else {
	// 	// income response
	// 	stream.SetState(net_p.Status_Ready)
	// }

	// if st.State == net_p.Status_Open {
	// TODO::: easily call by "go" or call by workers pool or what??
	// go st.callService()
	// return
	// }
	// st.SetState(net_p.Status_ReceivedCompletely)
}
