/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"hash/crc32"
	"strconv"
	"strings"

	"../assets"
	"../log"
)

// LoadAssetsFromStorage use to load assets from gui folder.
func LoadAssetsFromStorage(ass, repo *assets.Folder, repoLocation string) (main *assets.File) {
	readRepositoryFromFileSystem(repoLocation, repo)
	main = addRepoToAsset(ass, repo)
	log.Info("GUI changes successfully updated in server and ready to serve")
	return
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
func addRepoToAsset(ass, repo *assets.Folder) (main *assets.File) {
	var c = combine{}
	c.init(repo)
	c.do()

	// set design-languages
	var dl = c.repoGUI.GetDependency("design-languages")
	var file *assets.File
	for _, file = range dl.Files {
		if file.Extension == "css" {
			var fullName = file.FullName
			file.Name = strconv.FormatUint(uint64(crc32.ChecksumIEEE(file.Data)), 10)
			file.FullName = file.Name + ".css"
			ass.SetAndCompressFile(file, assets.CompressTypeGZIP)

			c.mainJS.Data = bytes.ReplaceAll(c.mainJS.Data, []byte(fullName), []byte(file.FullName))
		}
	}

	// ass.SetAndCompressFiles(c.initsJS, assets.CompressTypeGZIP)
	var initHashName = strconv.FormatUint(uint64(crc32.ChecksumIEEE(c.initsJS[0].Data)), 10)
	for _, file := range c.initsJS {
		file.Name = initHashName + "-" + strings.Split(file.Name, "-")[1] // strings.Split(file.Name, "-")[1] as file lang
		file.FullName = file.Name + ".js"
		ass.SetAndCompressFile(file, assets.CompressTypeGZIP)
	}

	c.mainJS.Data = bytes.ReplaceAll(c.mainJS.Data, []byte("/init-"), []byte("/"+initHashName+"-"))
	c.mainJS.Name = strconv.FormatUint(uint64(crc32.ChecksumIEEE(c.mainJS.Data)), 10)
	c.mainJS.FullName = c.mainJS.Name + ".js"
	ass.SetAndCompressFile(c.mainJS, assets.CompressTypeGZIP)

	// TODO::: Need to change landings file name to hash of data??
	ass.SetAndCompressFiles(c.landings, assets.CompressTypeGZIP)

	// set /main.html
	main = c.repoGUI.GetFile("main.html")
	main.Data = bytes.ReplaceAll(main.Data, []byte("main.js"), []byte(c.mainJS.FullName))
	main.Name = strconv.FormatUint(uint64(crc32.ChecksumIEEE(main.Data)), 10)
	main.FullName = main.Name + ".html"
	ass.SetAndCompressFile(main, assets.CompressTypeGZIP)

	// set /sw.js
	var swFile = c.repoGUI.GetDependency("libjs").GetDependency("gui-engine").GetFile("sw.js")
	swFile.Data = bytes.ReplaceAll(swFile.Data, []byte("main.html"), []byte(main.FullName))
	ass.SetAndCompressFile(swFile, assets.CompressTypeGZIP)

	// set images from gui
	var images = c.repoGUI.GetDependency("images")
	for _, file := range images.Files {
		ass.SetAndCompressFile(file, assets.CompressTypeGZIP)
	}

	return
}
