/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// GUI is default global protocol.GUI_Application like window global variable in browsers.
// You must assign to it by any object implement protocol.GUI_Application on your main.go file. Suggestion:
// - GUI App	>> protocol.GUI = &gui.Application
var GUI GUI_Application

// GUI_Application is UI (GUI, VUI, ...) specific protocols that include in Application interface
type GUI_Application interface {
	GUI_Pages
	GUI_Navigator
	GUI_History
}
