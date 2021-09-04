/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"fmt"

	"../log"
	"../protocol"
)

const (
	guiDirectoryName      = "gui"
	mainHTMLDirectoryName = "main-html"
)

type Assets struct {
	GUI      protocol.FileDirectory
	MainHTML protocol.FileDirectory // files name is just language in iso format e.g. "en", "fa",
	// OldBrowsers protocol.FileDirectory // files name is just language in iso format e.g. "en", "fa",
}

func (a *Assets) Init() {
	var err protocol.Error
	a.GUI, err = protocol.App.FileDirectory().Directory(guiDirectoryName)
	log.Fatal(err)
	a.MainHTML, err = a.GUI.Directory(mainHTMLDirectoryName)
	log.Fatal(err)
	a.update()
}

// ReloadByCLI block function and must call by seprate goroutine, otherwise it can block other app logic!
func (a *Assets) ReloadByCLI() {
	// defer Server.PanicHandler()
reload:
	log.Info("Write '''R''' & press '''Enter''' key to reload GUI changes")
	var non string
	fmt.Scanln(&non)
	if non == "R" || non == "r" {
		a.update()
	} else {
		log.Warn("Requested command not found")
	}
	goto reload
}

// Update use to add needed repo files that get from disk or network to the assets!!
func (a *Assets) update() {
	var c combine
	c.update()
	log.Info("WWW - GUI assets successfully updated and ready to serve")
}
