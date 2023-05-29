/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Application immutable runtime settings
const (
	// AppLanguage store global language to use by any locale text selector.
	AppLanguage = LanguageEnglish

	// AppMode_Dev use to indicate that app can do some more logic e.g.
	// - Save more logs
	// - Add more services like net/http/pprof for better debugging
	// - Add more pages that just need only for developers
	AppMode_Dev = true
)

// App is default global protocol.Application like window global variable in browsers.
// You must assign to it by any object implement protocol.Application on your main.go file. Suggestion:
// protocol.App = &achaemenid.App
var App Application

// Application is the interface that must implement by any Application.
// It introduce just local computing mechanism not network, storage, distributed, gui, ...
type Application interface {
	Engine() ApplicationEngine

	SoftwareStatus() SoftwareStatus

	Status() ApplicationState
	// Listen to the app state changes. Can return the channel instead of get as arg, but a channel listener can lost very fast app state changing.
	// This is because when the first goroutine blocks the channel all other goroutines must wait in line. When the channel is unblocked,
	// the state has already been received and removed from the channel so the next goroutine in line gets the next state value.
	NotifyState(notifyBy chan ApplicationState)

	OS_Signal_Listener
	Logger
	Services
	Errors
	EventTarget
	NetworkApplication_Multiplexer

	ObjectLifeCycle
}

// ApplicationState indicate application state
// Force to use 32 bit length due to it is minimum atomic helper functions size.
type ApplicationState uint32

// Application State
const (
	ApplicationState_Unset    ApplicationState = iota // State not set yet.
	ApplicationState_Starting                         // plan to start e.g. wait for hardware to be available
	ApplicationState_Running
	ApplicationState_Stopping
	ApplicationState_Stopped

	ApplicationState_PowerFailure

	ApplicationState_Stable
	ApplicationState_NotResponse

	ApplicationState_Splitting
	ApplicationState_ReAllocate
	// ApplicationState_AcceptWrite
)
