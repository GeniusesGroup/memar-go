/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"fmt"
	"os"
	"path/filepath"

	"../assets"
	"../log"
)

// ReloadAssetsFromStorage use to reload assets from gui folder usually in development phase!
// It block function and must call by seprate goroutine, otherwise it can block other app logic!
func ReloadAssetsFromStorage(ass *assets.Folder) {
	// Indicate repoLocation
	var ex, err = os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	var repoLocation = filepath.Dir(ex)

	var repo = assets.NewFolder("")
reload:
	readRepositoryFromFileSystem(repoLocation, repo)
	addRepoToAsset(ass, repo)
	log.Info("GUI changes successfully updated in server and ready to serve")

	log.Info("Press '''Enter''' key to reload GUI changes")
	var non string
	fmt.Scanln(&non)
	goto reload
}

// ReadRepositoryFromFileSystem use to get all repository by its name!
func readRepositoryFromFileSystem(dirname string, repo *assets.Folder) (err error) {
	var gui, sdk *assets.Folder

	// gui folder
	gui = repo.GetDependency("gui")
	if gui == nil {
		gui = assets.NewFolder("gui")
		repo.SetDependency(gui)
	}
	err = gui.ReadRepositoryFromFileSystem(dirname + "/gui")

	// sdk-js folder
	sdk = repo.GetDependency("sdk-js")
	if sdk == nil {
		sdk = assets.NewFolder("sdk-js")
		repo.SetDependency(sdk)
	}
	err = sdk.ReadRepositoryFromFileSystem(dirname + "/sdk-js")

	return
}

// addRepoToAsset use to add needed repo files that get from disk or network to the assets!!
func addRepoToAsset(ass, repo *assets.Folder) {
	var c = combine{}
	c.init(repo)
	c.do()
	ass.SetFile(c.mainJS)
	ass.SetFiles(c.initsJS)
	ass.SetFiles(c.landings)

	// set images from gui
	var images = c.repoGUI.GetDependency("images")
	for _, file := range images.Files {
		ass.SetFile(file)
	}

	// set main.html & sw.js
	ass.SetFile(c.repoGUI.GetFile("main.html"))
	ass.SetFile(c.repoGUI.GetDependency("libjs").GetDependency("gui-engine").GetFile("sw.js"))

	// set design-languages
	var dl = c.repoGUI.GetDependency("design-languages")
	var file *assets.File
	for _, file = range dl.Files {
		if file.Extension == "css" {
			ass.SetFile(file)
		}
	}
}
