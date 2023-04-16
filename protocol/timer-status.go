/* For license and copyright information please see the LEGAL file in the code repository */

package protocol


type TimerStatus uint32

// Values for the timer status field.
const (
	// Timer has no status set yet.
	TimerStatus_Unset TimerStatus = iota

	// Waiting for timer to fire.
	// The timer is in some P's heap.
	TimerStatus_Waiting

	// Running the timer function.
	// A timer will only have this status briefly.
	TimerStatus_Running

	// The timer is deleted and should be removed.
	// It should not be run, but it is still in some P's heap.
	TimerStatus_Deleted

	// The timer is being removed.
	// The timer will only have this status briefly.
	TimerStatus_Removing

	// The timer has been stopped.
	// It is not in any P's heap.
	TimerStatus_Removed

	// The timer is being modified.
	// The timer will only have this status briefly.
	TimerStatus_Modifying

	// The timer has been modified to an earlier time.
	// The new when value is in the timer.when and old on in the timerBucket.when field.
	// The timer is in some P's heap, possibly in the wrong place.
	TimerStatus_ModifiedEarlier

	// The timer has been modified to the same or a later time.
	// The new when value is in the timer.when and old on in the timerBucket.when field.
	// The timer is in some P's heap, possibly in the wrong place.
	TimerStatus_ModifiedLater

	// The timer has been modified and is being moved.
	// The timer will only have this status briefly.
	TimerStatus_Moving
)