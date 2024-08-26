/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

// GUI is default global protocol.Application like window global variable in browsers.
// You must assign to it by any object implement protocol.Application on your main.go file. Suggestion:
// - GUI App	>> gui_p.GUI = &gui.Application
var GUI Application

// Application is UI (GUI, VUI, ...) specific protocols that include in Application interface
type Application interface {
	Pages
	Navigator
	History
}
