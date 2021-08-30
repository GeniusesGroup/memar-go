/* For license and copyright information please see LEGAL file in repository */

package protocol

// Application is the interface that must implement by any Application!
type Application interface {
	OS() OS
	Manifest() ApplicationManifest
	PanicHandler()

	ObjectDirectory() ObjectDirectory // Distributed object storage
	FileDirectory() FileDirectory // Distributed file storage

	Services
	Errors

	NetworkApplicationMultiplexer
}

// ApplicationManifest is the interface that must implement by any Application!
type ApplicationManifest interface {
	DomainName() string

	UserUUID() (userUUID [32]byte)
	UserID() (userID uint32)
	AppUUID() (appUUID [32]byte)
	AppID() (appID uint16)
}

// ApplicationState indicate application state
type ApplicationState uint8

// Application State
const (
	ApplicationStateUnset ApplicationState = iota // State not set yet!
	ApplicationStateStop
	ApplicationStateRunning
	ApplicationStateStopping
	ApplicationStateStarting // plan to start
)
