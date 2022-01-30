/* For license and copyright information please see LEGAL file in repository */

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

// App is default global protocol.Application
// You must assign to it by any object implement protocol.Application on your main.go file
// Suggestion: protocol.App = &achaemenid.App
var App Application

// Application is the interface that must implement by any Application.
type Application interface {
	SoftwareStatus() SoftwareStatus
	Status() ApplicationState
	// Listen to the app state changes. Can return the channel instead of get as arg, but a channel listener can lost very fast app state changing.
	// This is because when the first goroutine blocks the channel all other goroutines must wait in line. When the channel is unblocked,
	// the state has already been received and removed from the channel so the next goroutine in line gets the next state value.
	NotifyState(notifyBy chan ApplicationState)
	Shutdown()

	Cluster
	Storages
	Logger
	Services
	Errors
	Connections
	NetworkApplicationMultiplexer
}

// ApplicationManifest is the interface that must implement by any Application.
type ApplicationManifest interface {
	DomainName() string
	Email() string

	UserUUID() (userUUID [32]byte)
	UserID() (userID uint32)
	AppUUID() (appUUID [32]byte)
	AppID() (appID uint16) // local OS application ID
}

// ApplicationState indicate application state
// Force to use 32 bit length due to it is minimum atomic helper functions size.
type ApplicationState uint32

// Application State
const (
	ApplicationStateUnset    ApplicationState = iota // State not set yet!
	ApplicationStateStarting                         // plan to start
	ApplicationStateRunning
	ApplicationStateStopping
	ApplicationStateStoped

	ApplicationStateLocalNode
	ApplicationStateStable
	ApplicationStateNotResponse

	ApplicationStateSplitting
	ApplicationStateReAllocate
	// ApplicationStateAcceptWrite
)
