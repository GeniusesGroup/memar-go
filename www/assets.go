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
func LoadAssetsFromStorage(repoLocation string) (ass *assets.Folder, main *assets.File) {
	var repo = assets.NewFolder("repo")
	ass = assets.NewFolder("www")
	readRepositoryFromFileSystem(repoLocation, repo)
	main = AddRepoToAsset(ass, repo)
	log.Info("WWW - GUI assets successfully updated in server and ready to serve")
	return
}

// ReadRepositoryFromFileSystem use to get all repository by its name!
func readRepositoryFromFileSystem(dirname string, repo *assets.Folder) (err error) {
	// gui folder
	var gui *assets.Folder
	gui = assets.NewFolder("gui")
	gui.Dep = repo
	repo.SetDependency(gui)
	err = gui.ReadRepositoryFromFileSystem(dirname+"/gui", false)

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
	for _, folder := range dl.Dependencies {
		var combinedFile = assets.File{
			Name:      "design-language--" + folder.Name,
			Extension: "css",
		}
		combinedFile.CheckAndFix()
		var fullName = combinedFile.FullName
		for _, file := range folder.Files {
			if file.Extension == "css" {
				combinedFile.Data = append(combinedFile.Data, file.Data...)
			}
		}
		combinedFile.AddHashToName()
		ass.MinifyCompressSet(&combinedFile, assets.CompressTypeGZIP)

		c.mainJS.Data = bytes.ReplaceAll(c.mainJS.Data, convert.UnsafeStringToByteSlice(fullName), convert.UnsafeStringToByteSlice(combinedFile.FullName))
	}
	for _, file := range dl.Files {
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

	for _, file := range c.otherFiles {
		// file.AddHashToName()
		ass.MinifyCompressSet(file, assets.CompressTypeGZIP)
	}

	return c.mainHTML
}
