/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"fmt"

	"../codec"
	"../file"
	"../log"
)

type Assets struct {
	Folder       file.Folder
	MainHTML     file.Folder // files name is just language in iso format e.g. "en", "fa",
	repoLocation string
	Repo         *file.Folder
}

// LoadAssetsFromStorage use to load assets from gui folder.
func (a *Assets) Init(dirLocation string) {
	a.repoLocation = dirLocation
	a.Folder.Init("www", "/", false)
}

// It block function and must call by seprate goroutine, otherwise it can block other app logic!
func (a *Assets) ReLoadFromStorageByCLI() {
	// defer Server.PanicHandler()
reload:
	log.Info("Write '''R''' & press '''Enter''' key to reload GUI changes")
	var non string
	fmt.Scanln(&non)
	if non == "R" || non == "r" {
		a.LoadAssetsFromStorage()
	} else {
		log.Warn("Requested command not found")
	}
	goto reload
}

// LoadAssetsFromStorage use to load assets from gui folder.
func (a *Assets) LoadAssetsFromStorage() {
	a.Repo = file.NewFolder("repo", "", true)
	a.readRepositoryFromFileSystem()
	a.Update()
	log.Info("WWW - GUI assets successfully updated in HTTP asset and ready to serve")
	return
}

// ReadRepositoryFromFileSystem use to get all repository by its name!
func (a *Assets) readRepositoryFromFileSystem() (err error) {
	// gui folder
	var gui *file.Folder = file.NewFolder("gui", a.repoLocation+"/gui", false)
	gui.Folder = a.Repo
	a.Repo.SetDependency(gui)
	err = gui.Update()
	return
}

// Update use to add needed repo files that get from disk or network to the assets!!
func (a *Assets) Update() {
	var c = combine{}
	c.init(a.Repo)

	for _, fi := range c.swFiles {
		a.Folder.MinifyCompressSet(fi, codec.CompressTypeGZIP)
	}
	for _, fi := range c.mainHTMLFiles {
		a.Folder.MinifyCompressSet(fi, codec.CompressTypeGZIP)
	}
	for _, fi := range c.mainJSFiles {
		a.Folder.MinifyCompressSet(fi, codec.CompressTypeGZIP)
	}
	for _, files := range c.landingsFiles {
		for lang, fi := range files {
			fi.Rename(fi.Name + "-" + lang)
			a.Folder.MinifyCompressSet(fi, codec.CompressTypeGZIP)
		}
	}
	a.MainHTML.Files = c.mainHTMLFiles
	a.Folder.MinifyCompressSets(c.designLanguages, codec.CompressTypeGZIP)

	// set images from gui
	var images = c.repoGUI.GetDependency("images")
	for _, imageFile := range images.Files {
		a.Folder.MinifyCompressSet(imageFile, codec.CompressTypeGZIP)
	}

	for _, otherFile := range c.otherFiles {
		a.Folder.MinifyCompressSet(otherFile, codec.CompressTypeGZIP)
	}
}
