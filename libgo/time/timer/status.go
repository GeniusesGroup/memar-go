/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"sync/atomic"
)

// timer.Start:
//   Status_Unset			-> Status_Waiting
//   anything else			-> panic: invalid value
//
// timer.Stop:
//   Status_Waiting			-> Status_Modifying		-> Status_Deleted
//   Status_ModifiedEarlier	-> Status_Modifying		-> Status_Deleted
//   Status_ModifiedLater	-> Status_Modifying		-> Status_Deleted
//   Status_Unset			-> do nothing
//   Status_Deleted         -> do nothing
//   Status_Removing        -> do nothing
//   Status_Removed         -> do nothing
//   Status_Running         -> wait until status changes
//   Status_Moving          -> wait until status changes
//   Status_Modifying       -> wait until status changes
//
// timer.Reset:
//   Status_Waiting		-> Status_Modifying			-> timerModifiedXX
//   timerModifiedXX	-> Status_Modifying			-> timerModifiedYY
//   Status_Unset		-> Status_Modifying			-> Status_Waiting
//   Status_Removed		-> Status_Modifying			-> Status_Waiting
//   Status_Deleted		-> Status_Modifying			-> timerModifiedXX
//   Status_Running		-> wait until status changes
//   Status_Moving		-> wait until status changes
//   Status_Removing	-> wait until status changes
//   Status_Modifying	-> wait until status changes
//
// timing.cleanTimers (looks in timers heap):
//   Status_Deleted		-> Status_Removing			-> Status_Removed
//   timerModifiedXX	-> Status_Moving			-> Status_Waiting
//
// timing.adjustTimers (looks in timers heap):
//   Status_Deleted		-> Status_Removing			-> Status_Removed
//   timerModifiedXX	-> Status_Moving			-> Status_Waiting
//
// timing.runTimer (looks in timers heap):
//   Status_Unset		-> panic: uninitialized timer
//   Status_Waiting		-> Status_Waiting or
//   Status_Waiting		-> Status_Running			-> Status_Unset or
//   Status_Waiting		-> Status_Running			-> Status_Waiting
//   Status_Modifying	-> wait until status changes
//   timerModifiedXX	-> Status_Moving			-> Status_Waiting
//   Status_Deleted		-> Status_Removing			-> Status_Removed
//   Status_Running		-> panic: concurrent runTimer calls
//   Status_Removed		-> panic: inconsistent timer heap
//   Status_Removing	-> panic: inconsistent timer heap
//   Status_Moving		-> panic: inconsistent timer heap

// Values for the timer status field.
const (
	// Timer has no status set yet.
	Status_Unset Status = iota

	// Waiting for timer to fire.
	// The timer is in some P's heap.
	Status_Waiting

	// Running the timer function.
	// A timer will only have this status briefly.
	Status_Running

	// The timer is deleted and should be removed.
	// It should not be run, but it is still in some P's heap.
	Status_Deleted

	// The timer is being removed.
	// The timer will only have this status briefly.
	Status_Removing

	// The timer has been stopped.
	// It is not in any P's heap.
	Status_Removed

	// The timer is being modified.
	// The timer will only have this status briefly.
	Status_Modifying

	// The timer has been modified to an earlier time.
	// The new when value is in the timer.when and old on in the timerBucket.when field.
	// The timer is in some P's heap, possibly in the wrong place.
	Status_ModifiedEarlier

	// The timer has been modified to the same or a later time.
	// The new when value is in the timer.when and old on in the timerBucket.when field.
	// The timer is in some P's heap, possibly in the wrong place.
	Status_ModifiedLater

	// The timer has been modified and is being moved.
	// The timer will only have this status briefly.
	Status_Moving
)

// due to use atomic must use uint32
type Status uint32

func (s *Status) Load() Status {
	return Status(atomic.LoadUint32((*uint32)(s)))
}
func (s *Status) Store(status Status) {
	atomic.StoreUint32((*uint32)(s), uint32(status))
}
func (s *Status) CompareAndSwap(old, new Status) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(s), uint32(old), uint32(new))
}
