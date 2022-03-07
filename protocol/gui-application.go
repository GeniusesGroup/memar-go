/* For license and copyright information please see LEGAL file in repository */

package protocol

// GUIApplication is UI (GUI, VUI, ...) specific protocols that include in Application interface
// All below projects have many problems.
// e.g. https://github.com/maxence-charriere/go-app/blob/master/pkg/ui/scroll.go indicate scroll as content not behavior of any DOM content
// https://github.com/maxence-charriere/go-app
// https://github.com/gioui/gio
// https://github.com/asticode/go-astilectron
// https://github.com/zserge/lorca
// https://github.com/wailsapp/wails
// https://github.com/sciter-sdk/go-sciter
type GUIApplication interface {
	GUIPages
	GUINavigator
	GUIHistory
}
