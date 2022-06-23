/* For license and copyright information please see LEGAL file in repository */

package protocol

// GUIApplication is UI (GUI, VUI, ...) specific protocols that include in Application interface
// All below projects have many problems.
// https://docs.flutter.dev/		>> Mix content(HTML) and style(CSS) in to logic(JS, Dart, Go, ...) language
// https://github.com/maxence-charriere/go-app/blob/master/pkg/ui/scroll.go		>> indicate scroll as a content not behavior of elements
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

/*
## Reference:
- https://webvision.mozilla.org/full/
- https://extensiblewebmanifesto.org/
- https://open-ui.org/
- https://www.chromium.org/teams/web-capabilities-fugu/
- https://fugu-tracker.web.app/
*/
