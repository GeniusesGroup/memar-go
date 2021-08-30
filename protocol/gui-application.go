/* For license and copyright information please see LEGAL file in repository */

package protocol

// GUIApplication is the interface that must implement by any UI (GUI, VUI, ...) Application!
type GUIApplication interface {
	Application
	DOM
	GUIPages

	History() GUIHistory
}
