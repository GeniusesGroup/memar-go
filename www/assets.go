/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"fmt"

	"../protocol"
)

const (
	guiDirectoryName      = "gui"
	mainHTMLDirectoryName = "main-html"
)

type Assets struct {
	GUI         protocol.FileDirectory
	MainHTMLDir protocol.FileDirectory // files name is just language in iso format e.g. "en", "fa",
	// OldBrowsers protocol.FileDirectory // files name is just language in iso format e.g. "en", "fa",
	ContentEncodings []string
}

func (a *Assets) Init() {
	var err protocol.Error
	a.GUI, err = protocol.App.Files().Directory(guiDirectoryName)
	protocol.App.LogFatal(err)
	a.MainHTMLDir, err = a.GUI.Directory(mainHTMLDirectoryName)
	protocol.App.LogFatal(err)
	a.update()
}

// ReloadByCLI block function and must call by seprate goroutine, otherwise it can block other app logic!
func (a *Assets) ReloadByCLI() {
	// defer Server.PanicHandler()
reload:
	protocol.App.Log(protocol.LogType_Information, "Write '''R''' & press '''Enter''' key to reload GUI changes")
	var non string
	fmt.Scanln(&non)
	if non == "R" || non == "r" {
		a.update()
	} else {
		protocol.App.Log(protocol.LogType_Warning, "Requested command not found")
	}
	goto reload
}

// Update use to add needed repo files that get from disk or network to the assets!!
func (a *Assets) update() {
	var c = combine{
		contentEncodings: a.ContentEncodings,
	}
	c.update()
	protocol.App.Log(protocol.LogType_Information, "WWW - GUI assets successfully updated and ready to serve")
}
