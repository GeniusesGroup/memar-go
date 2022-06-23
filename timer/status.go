/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"sync/atomic"
)

// addtimer:
//   status_Unset   -> status_Waiting
//   anything else   -> panic: invalid value
// deltimer:
//   status_Waiting         -> status_Modifying -> status_Deleted
//   status_ModifiedEarlier -> status_Modifying -> status_Deleted
//   status_ModifiedLater   -> status_Modifying -> status_Deleted
//   status_Unset        -> do nothing
//   status_Deleted         -> do nothing
//   status_Removing        -> do nothing
//   status_Removed         -> do nothing
//   status_Running         -> wait until status changes
//   status_Moving          -> wait until status changes
//   status_Modifying       -> wait until status changes
// modtimer:
//   status_Waiting    -> status_Modifying -> timerModifiedXX
//   timerModifiedXX -> status_Modifying -> timerModifiedYY
//   status_Unset   -> status_Modifying -> status_Waiting
//   status_Removed    -> status_Modifying -> status_Waiting
//   status_Deleted    -> status_Modifying -> timerModifiedXX
//   status_Running    -> wait until status changes
//   status_Moving     -> wait until status changes
//   status_Removing   -> wait until status changes
//   status_Modifying  -> wait until status changes
// cleantimers (looks in P's timer heap):
//   status_Deleted    -> status_Removing -> status_Removed
//   timerModifiedXX -> status_Moving -> status_Waiting
// adjusttimers (looks in P's timer heap):
//   status_Deleted    -> status_Removing -> status_Removed
//   timerModifiedXX -> status_Moving -> status_Waiting
// runtimer (looks in P's timer heap):
//   status_Unset   -> panic: uninitialized timer
//   status_Waiting    -> status_Waiting or
//   status_Waiting    -> status_Running -> status_Unset or
//   status_Waiting    -> status_Running -> status_Waiting
//   status_Modifying  -> wait until status changes
//   timerModifiedXX -> status_Moving -> status_Waiting
//   status_Deleted    -> status_Removing -> status_Removed
//   status_Running    -> panic: concurrent runtimer calls
//   status_Removed    -> panic: inconsistent timer heap
//   status_Removing   -> panic: inconsistent timer heap
//   status_Moving     -> panic: inconsistent timer heap

// Values for the timer status field.
const (
	// Timer has no status set yet.
	status_Unset status = iota

	// Waiting for timer to fire.
	// The timer is in some P's heap.
	status_Waiting

	// Running the timer function.
	// A timer will only have this status briefly.
	status_Running

	// The timer is deleted and should be removed.
	// It should not be run, but it is still in some P's heap.
	status_Deleted

	// The timer is being removed.
	// The timer will only have this status briefly.
	status_Removing

	// The timer has been stopped.
	// It is not in any P's heap.
	status_Removed

	// The timer is being modified.
	// The timer will only have this status briefly.
	status_Modifying

	// The timer has been modified to an earlier time.
	// The new when value is in the nextwhen field.
	// The timer is in some P's heap, possibly in the wrong place.
	status_ModifiedEarlier

	// The timer has been modified to the same or a later time.
	// The new when value is in the nextwhen field.
	// The timer is in some P's heap, possibly in the wrong place.
	status_ModifiedLater

	// The timer has been modified and is being moved.
	// The timer will only have this status briefly.
	status_Moving
)

// due to use atomic must use uint32
type status uint32

func (s *status) Load() status {
	return status(atomic.LoadUint32((*uint32)(s)))
}
func (s *status) Store(status status) {
	atomic.StoreUint32((*uint32)(s), uint32(status))
}
func (s *status) CompareAndSwap(old, new status) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(s), uint32(old), uint32(new))
}
