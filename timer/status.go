/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"sync/atomic"

	"memar/protocol"
)

// timer.Start:
//   protocol.TimerStatus_Unset   -> protocol.TimerStatus_Waiting
//   anything else                -> panic: invalid value
//
// timer.Stop:
//   protocol.TimerStatus_Waiting         -> protocol.TimerStatus_Modifying -> protocol.TimerStatus_Deleted
//   protocol.TimerStatus_ModifiedEarlier -> protocol.TimerStatus_Modifying -> protocol.TimerStatus_Deleted
//   protocol.TimerStatus_ModifiedLater   -> protocol.TimerStatus_Modifying -> protocol.TimerStatus_Deleted
//   protocol.TimerStatus_Unset           -> do nothing
//   protocol.TimerStatus_Deleted         -> do nothing
//   protocol.TimerStatus_Removing        -> do nothing
//   protocol.TimerStatus_Removed         -> do nothing
//   protocol.TimerStatus_Running         -> wait until status changes
//   protocol.TimerStatus_Moving          -> wait until status changes
//   protocol.TimerStatus_Modifying       -> wait until status changes
//
// timer.Reset:
//   protocol.TimerStatus_Waiting    -> protocol.TimerStatus_Modifying -> timerModifiedXX
//   timerModifiedXX                 -> protocol.TimerStatus_Modifying -> timerModifiedYY
//   protocol.TimerStatus_Unset      -> protocol.TimerStatus_Modifying -> protocol.TimerStatus_Waiting
//   protocol.TimerStatus_Removed    -> protocol.TimerStatus_Modifying -> protocol.TimerStatus_Waiting
//   protocol.TimerStatus_Deleted    -> protocol.TimerStatus_Modifying -> timerModifiedXX
//   protocol.TimerStatus_Running    -> wait until status changes
//   protocol.TimerStatus_Moving     -> wait until status changes
//   protocol.TimerStatus_Removing   -> wait until status changes
//   protocol.TimerStatus_Modifying  -> wait until status changes
//
// timing.cleanTimers (looks in timers heap):
//   protocol.TimerStatus_Deleted    -> protocol.TimerStatus_Removing -> protocol.TimerStatus_Removed
//   timerModifiedXX                 -> protocol.TimerStatus_Moving   -> protocol.TimerStatus_Waiting
//
// timing.adjustTimers (looks in timers heap):
//   protocol.TimerStatus_Deleted    -> protocol.TimerStatus_Removing -> protocol.TimerStatus_Removed
//   timerModifiedXX                 -> protocol.TimerStatus_Moving   -> protocol.TimerStatus_Waiting
//
// timing.runTimer (looks in timers heap):
//   protocol.TimerStatus_Unset      -> panic: uninitialized timer
//   protocol.TimerStatus_Waiting    -> protocol.TimerStatus_Waiting or
//   protocol.TimerStatus_Waiting    -> protocol.TimerStatus_Running  -> protocol.TimerStatus_Unset or
//   protocol.TimerStatus_Waiting    -> protocol.TimerStatus_Running  -> protocol.TimerStatus_Waiting
//   protocol.TimerStatus_Modifying  -> wait until status changes
//   timerModifiedXX                 -> protocol.TimerStatus_Moving   -> protocol.TimerStatus_Waiting
//   protocol.TimerStatus_Deleted    -> protocol.TimerStatus_Removing -> protocol.TimerStatus_Removed
//   protocol.TimerStatus_Running    -> panic: concurrent runTimer calls
//   protocol.TimerStatus_Removed    -> panic: inconsistent timer heap
//   protocol.TimerStatus_Removing   -> panic: inconsistent timer heap
//   protocol.TimerStatus_Moving     -> panic: inconsistent timer heap

// due to use atomic must use uint32
type status uint32

func (s *status) Load() protocol.TimerStatus {
	return protocol.TimerStatus(atomic.LoadUint32((*uint32)(s)))
}
func (s *status) Store(status protocol.TimerStatus) {
	atomic.StoreUint32((*uint32)(s), uint32(status))
}
func (s *status) CompareAndSwap(old, new protocol.TimerStatus) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(s), uint32(old), uint32(new))
}
