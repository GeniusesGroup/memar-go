/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Application is the interface that must implement by any Application.
// It introduce just local computing mechanism not network, storage, distributed, gui, ...
type Application_Status interface {
	Status() ApplicationStatus
	// Listen to the app state changes. Can return the channel instead of get as arg, but a channel listener can lost very fast app state changing.
	// This is because when the first goroutine blocks the channel all other goroutines must wait in line. When the channel is unblocked,
	// the state has already been received and removed from the channel so the next goroutine in line gets the next state value.
	NotifyStatus(notifyBy chan ApplicationStatus)
}

// ApplicationStatus indicate application state
// Force to use 32 bit length due to it is minimum atomic helper functions size.
type ApplicationStatus uint32

const (
	ApplicationStatus_Unset    ApplicationStatus = 0           // Status not set yet.
	ApplicationStatus_Starting ApplicationStatus = (1 << iota) // plan to start e.g. wait for hardware to be available
	ApplicationStatus_Running
	ApplicationStatus_Stopping
	ApplicationStatus_Stopped

	ApplicationStatus_Power_Failure
	ApplicationStatus_Power_Green // provide by any recyclable sources
	ApplicationStatus_Power_UPS
	ApplicationStatus_Power_Diesel

	ApplicationStatus_Stable
	ApplicationStatus_NotResponse

	ApplicationStatus_Splitting
	ApplicationStatus_ReAllocate
	// ApplicationStatus_AcceptWrite
)
