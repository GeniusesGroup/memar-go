/* For license and copyright information please see LEGAL file in repository */

package protocol

// App is default global protocol.Application
// You must assign to it by any object implement protocol.Application on your main.go file
// Suggestion: protocol.App = &achaemenid.App
var App Application

// Application is the interface that must implement by any Application!
type Application interface {
	SoftwareStatus() SoftwareStatus
	Status() ApplicationState
	ObjectDirectory() ObjectDirectory // Distributed object storage
	FileDirectory() FileDirectory     // Distributed file storage
	Shutdown()

	Logger
	ApplicationNodes
	Services
	Errors
	Connections
	NetworkApplicationMultiplexer
}

// ApplicationManifest is the interface that must implement by any Application!
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
