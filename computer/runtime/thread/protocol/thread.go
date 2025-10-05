/* For license and copyright information please see the LEGAL file in the code repository */

package thread_p

import (
	error_p "memar/process/error/protocol"
	"memar/time/duration"
)

type Thread interface {
	Sleep(d duration.NanoSecond)

	// Defer(where ??, v any)

	// Thread_Execution_Exception
	Thread_SchedulerWaiting

	Runtime
}

type Thread_SchedulerWaiting interface {
	AddToWaitList(id int) (err error_p.Error)
	NotifyWaitList(id int) (err error_p.Error)
	Wait() (id int, er error_p.Error)
}

// Even when a package uses panic internally, its external API still presents explicit error return values.
// TODO::: Is it good idea to have these logics? Why not just Errors??
// type Thread_Execution_Exception interface {
// 	Panic(e Event)
// 	Recover() (e Event)
// }
