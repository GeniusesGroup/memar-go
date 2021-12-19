/* For license and copyright information please see LEGAL file in repository */

package protocol

// GUIApp is default global protocol.GUIApplication
// You must assign to it by any object implement protocol.GUIApplication on your main.go file
// Suggestion: protocol.GUIApp = &gui.Application
var GUIApp GUIApplication

// GUIApplication is the interface that must implement by any UI (GUI, VUI, ...) Application!
type GUIApplication interface {
	Application
	DOM
	GUIPages
	GUIHistory
}
