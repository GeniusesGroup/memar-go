/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"os"
	"path/filepath"

	as "../assets"
	"../log"
)

var repoLocation string

// Any data files to serve or use by server!
type assets struct {
	GUI     *as.Folder
	Objects *as.Folder
	Secret  *as.Folder
	WWW     *as.Folder
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

	a.Secret.ReadRepositoryFromFileSystem(repoLocation + "/secret")
}

func (a *assets) shutdown() {
	// write secret files to storage device if any change made
	a.Secret.WriteRepositoryToFileSystem(repoLocation + "/secret")
}
