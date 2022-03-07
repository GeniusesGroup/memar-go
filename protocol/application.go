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

// App is default global protocol.Application like window global variable in browsers.
// You must assign to it by any object implement protocol.Application on your main.go file. Suggestion:
// - Server App	>> protocol.App = &achaemenid.App
// - GUI App	>> protocol.App = &gui.Application
var App Application

// Application is the interface that must implement by any Application.
type Application interface {
	Engine() ApplicationEngine

	SoftwareStatus() SoftwareStatus
	Status() ApplicationState
	// Listen to the app state changes. Can return the channel instead of get as arg, but a channel listener can lost very fast app state changing.
	// This is because when the first goroutine blocks the channel all other goroutines must wait in line. When the channel is unblocked,
	// the state has already been received and removed from the channel so the next goroutine in line gets the next state value.
	NotifyState(notifyBy chan ApplicationState)
	Shutdown()

	Cluster
	StoragesLocal
	StoragesCache
	Logger
	Services
	Errors
	Connections
	NetworkApplicationMultiplexer
	EventTarget

	// Server specific applications
	StoragesDistributed

	GUIApplication
}

// ApplicationEngine is the interface that return some useful data about the engine that implement Application protocol
// In many ways it is like window.navigator in web ecosystem
type ApplicationEngine interface {
	Name() string
	CharacterSet() string
}

// ApplicationManifest is the interface that must implement by any Application.
type ApplicationManifest interface {
	Icon() []byte
	DomainName() string
	Email() string

	UserUUID() (userUUID [32]byte)
	UserID() (userID uint32)
	AppUUID() (appUUID [32]byte)
	AppID() (appID uint16) // local OS application ID

	ContentPreferences()
	PresentationPreferences()
}

// ApplicationState indicate application state
// Force to use 32 bit length due to it is minimum atomic helper functions size.
type ApplicationState uint32

// Application State
const (
	ApplicationState_Unset    ApplicationState = iota // State not set yet!
	ApplicationState_Starting                         // plan to start
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
