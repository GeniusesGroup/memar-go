/* For license and copyright information please see LEGAL file in repository */

package giti

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

// Application is the interface that must implement by any Application!
type Application interface {
	Services

	Errors

	StorageEngine

	NetworkApplicationMultiplexer
}
