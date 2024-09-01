/* For license and copyright information please see the LEGAL file in the code repository */

package runtime_p

import (
	"memar/protocol"
)

type Thread interface {
	// Defer(where ??, v any)

	// Thread_Execution_Exception
	Thread_SchedulerWaiting

	Runtime
}

type Thread_SchedulerWaiting interface {
	AddToWaitList(id int) (err protocol.Error)
	NotifyWaitList(id int) (err protocol.Error)
	Wait() (id int, er protocol.Error)
}

// Even when a package uses panic internally, its external API still presents explicit error return values.
// TODO::: Is it good idea to have these logics? Why not just Errors??
// type Thread_Execution_Exception interface {
// 	Panic(e Event)
// 	Recover() (e Event)
// }
