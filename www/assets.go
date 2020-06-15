/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"../assets"
)

// ReloadAssetsInDevPhase use to reload assets from gui folder!
// It is blocking function, call by seprate goroutine, otherwise it can block other app logic!
func ReloadAssetsInDevPhase(a *assets.Folder) {
	// Indicate repoLocation
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information, So we can't specify service root location")
	}
	var repoLocation = path.Dir(path.Dir(path.Dir(filename)))

	var repo = assets.NewFolder("")
reload:
	readRepositoryFromFileSystem(repoLocation, repo)
	addRepo(a, repo)
	addGUIToMain(a, repo)

	fmt.Fprintf(os.Stderr, "%v\n", "Press '''Enter''' key to reload GUI changes")
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

// AddRepo use to add repo that get from disk or network to the assets!!
func addRepo(a, repo *assets.Folder) {
	var gui = repo.GetDependency("gui")

	// set images from gui
	var images = gui.GetDependency("images")
	for _, file := range images.Files {
		a.SetFile(file)
	}

	// set main.html & sw.js
	a.SetFile(gui.GetFile("main.html"))
	a.SetFile(gui.GetDependency("libjs").GetDependency("gui-engine").GetFile("sw.js"))

	// set design-languages
	var dl = gui.GetDependency("design-languages")
	var file *assets.File
	for _, file = range dl.Files {
		if file.Extension == "css" {
			a.SetFile(file)
		}
	}
}

// AddGUIToMain use to add pages, landings, widgets to main.js!
func addGUIToMain(a, repo *assets.Folder) {
	var gui = repo.GetDependency("gui")

	var inline []*assets.File

	var main = gui.GetFile("main.js")
	// Inline main.js at end of combined main.js
	inline = append(inline, main)

	var mainJS = a.GetFile("main.js")
	if mainJS == nil {
		mainJS = &assets.File{
			FullName:  "main.js",
			Name:      "main",
			Extension: "js",
			MimeType:  main.MimeType,
		}
		a.SetFile(mainJS)
	}
	mainJS.DataString = ""

	var pages = gui.GetDependency("pages")
	var page *assets.File
	for _, page = range pages.Files {
		if page.Extension == "html" {
			addHTMLToJS(pages, page)
		} else if page.Extension == "css" {
			addCSSToJS(pages, page)
		} else if page.Extension == "js" {
			inline = append(inline, page)
		}
	}

	var landings = gui.GetDependency("landings")
	var landing *assets.File
	for _, landing = range landings.Files {
		if landing.Extension == "html" {
			addHTMLToJS(landings, landing)
		} else if landing.Extension == "css" {
			addCSSToJS(landings, landing)
		} else if landing.Extension == "js" {
			inline = append(inline, landing)
		}
	}

	var widgets = gui.GetDependency("widgets")
	var widget *assets.File
	for _, widget = range widgets.Files {
		if widget.Extension == "html" {
			addHTMLToJS(widgets, widget)
		} else if widget.Extension == "css" {
			addCSSToJS(widgets, widget)
		} else if widget.Extension == "js" {
			inline = append(inline, widget)
		}
	}

	// Inline gui engine at first of main.js
	inline = append(inline, gui.GetDependency("libjs").GetDependency("gui-engine").GetFile("application.js"))

	// InlineFilesRecursively use to inline files to the given file Recursively!!
	var inlined = map[string]*assets.File{}
	var i int = len(inline) - 1
	for ; i >= 0; i-- {
		addJSToJSRecursively(repo, inline[i], mainJS, inlined)
	}

	mainJS.Data = []byte(mainJS.DataString)
}
