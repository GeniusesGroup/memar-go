/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"fmt"
	"os"
	"path/filepath"

	as "../assets"
	"../log"
	"../www"
)

var repoLocation string

// Any data files to serve or use by server!
type assets struct {
	GUI     *as.Folder
	Objects *as.Folder
	Secret  *as.Folder
	WWW     *as.Folder
	WWWMain *as.File
}

func (a *assets) init() {
	a.GUI = as.NewFolder("gui")
	a.Objects = as.NewFolder("objects")
	a.Secret = as.NewFolder("secret")
	a.WWW = as.NewFolder("www")

	// Indicate repoLocation
	// TODO::: change to PersiaOS when it ready!
	var ex, err = os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	repoLocation = filepath.Dir(ex)
	log.Info("App start in", repoLocation)

	a.LoadFromStorage()
}

func (a *assets) shutdown() {
	// write secret files to storage device if any change made
	a.Secret.WriteRepositoryToFileSystem(repoLocation + "/secret")
}

// It block function and must call by seprate goroutine, otherwise it can block other app logic!
func (a *assets) LoadFromStorage() {
	a.Secret.ReadRepositoryFromFileSystem(repoLocation + "/secret")
	a.WWWMain = www.LoadAssetsFromStorage(a.WWW, a.GUI, repoLocation)
}

// It block function and must call by seprate goroutine, otherwise it can block other app logic!
func (a *assets) ReLoadFromStorage() {
	// defer Server.PanicHandler()
reload:
	log.Info("Press '''Enter''' key to reload GUI changes")
	var non string
	fmt.Scanln(&non)

	a.WWWMain = www.LoadAssetsFromStorage(a.WWW, a.GUI, repoLocation)
	goto reload
}
