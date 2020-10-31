/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strings"

	"../assets"
	"../convert"
	"../log"
)

// LoadAssetsFromStorage use to load assets from gui folder.
func LoadAssetsFromStorage(ass, repo *assets.Folder, repoLocation string) (main *assets.File) {
	readRepositoryFromFileSystem(repoLocation, repo)
	main = AddRepoToAsset(ass, repo)
	log.Info("WWW - GUI assets successfully updated in server and ready to serve")
	return
}

// ReadRepositoryFromFileSystem use to get all repository by its name!
func readRepositoryFromFileSystem(dirname string, repo *assets.Folder) (err error) {
	var gui, sdk *assets.Folder

	// gui folder
	gui = repo.GetDependency("gui")
	if gui == nil {
		gui = assets.NewFolder("gui")
		gui.Dep = repo
		repo.SetDependency(gui)
	}
	err = gui.ReadRepositoryFromFileSystem(dirname+"/gui", false)

	// sdk-js folder
	sdk = repo.GetDependency("sdk-js")
	if sdk == nil {
		sdk = assets.NewFolder("sdk-js")
		sdk.Dep = repo
		repo.SetDependency(sdk)
	}
	err = sdk.ReadRepositoryFromFileSystem(dirname+"/sdk-js", false)

	return
}

// AddRepoToAsset use to add needed repo files that get from disk or network to the assets!!
func AddRepoToAsset(ass, repo *assets.Folder) (main *assets.File) {
	var c = combine{}
	c.init(repo)
	c.readyLandingsFiles()
	c.readyAppJSFiles()

	// set design-languages
	var dl = c.repoGUI.GetDependency("design-languages")
	var file *assets.File
	for _, file = range dl.Files {
		if file.Extension == "css" {
			var fullName = file.FullName
			file.AddHashToName()
			ass.MinifyCompressSet(file, assets.CompressTypeGZIP)

			c.mainJS.Data = bytes.ReplaceAll(c.mainJS.Data, convert.UnsafeStringToByteSlice(fullName), convert.UnsafeStringToByteSlice(file.FullName))
		}
	}

	ass.MinifyCompressSets(c.initsJS, assets.CompressTypeGZIP)

	var initJSHashName = "init-" + strings.Split(c.initsJS[0].Name, "-")[1]
	c.mainJS.Data = bytes.ReplaceAll(c.mainJS.Data, []byte("/init"), []byte("/"+initJSHashName))
	c.mainJS.AddHashToName()
	ass.MinifyCompressSet(c.mainJS, assets.CompressTypeGZIP)

	// TODO::: Need to change landings file name to hash of data??
	ass.MinifyCompressSets(c.landings, assets.CompressTypeGZIP)

	// set /main.html
	c.readyMainHTMLFile(c.mainJS.Name)
	ass.MinifyCompressSet(c.mainHTML, assets.CompressTypeGZIP)

	// set /sw.js
	var swFile = c.repoGUI.GetDependency("libjs").GetFile("sw.js")
	swFile.Data = bytes.ReplaceAll(swFile.Data, []byte("main.html"), []byte(c.mainHTML.FullName))
	ass.MinifyCompressSet(swFile, assets.CompressTypeGZIP)

	// set images from gui
	var images = c.repoGUI.GetDependency("images")
	for _, file := range images.Files {
		ass.MinifyCompressSet(file, assets.CompressTypeGZIP)
	}

	return c.mainHTML
}
