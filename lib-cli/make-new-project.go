/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"unsafe"

	ag "../Achaemenid-generator"
	gg "../Ganjine-generator"
	"../assets"
	wg "../www-generator"
)

// MakeNewProjectReq :
type MakeNewProjectReq struct {
	Repo *assets.Folder
}

// MakeNewProjectRes :
type MakeNewProjectRes struct{}

// MakeNewProject :
func MakeNewProject(req *MakeNewProjectReq) (res *MakeNewProjectRes, err error) {
	/* Folders */
	var servicesF = assets.NewFolder(FolderNameServices)
	servicesF.State = assets.StateChanged
	req.Repo.SetDependency(servicesF)
	var datastoreF = assets.NewFolder(FolderNameDataStore)
	datastoreF.State = assets.StateChanged
	req.Repo.SetDependency(datastoreF)
	var guiF = assets.NewFolder(FolderNameGUI)
	guiF.State = assets.StateChanged
	req.Repo.SetDependency(guiF)
	var secretF = assets.NewFolder(FolderNameSecret)
	secretF.State = assets.StateChanged
	req.Repo.SetDependency(secretF)

	/* Services */
	// /services/init.go
	var servicesInitGO = &assets.File{}
	err = ag.MakeServicesInitFile(servicesInitGO)
	if err != nil {
		return nil, err
	}
	servicesF.SetFile(servicesInitGO)

	/* Datastore */
	// /datastore/init.go
	var datastoreInitGO = &assets.File{}
	err = gg.MakeDataStoreInitFile(datastoreInitGO)
	if err != nil {
		return nil, err
	}
	datastoreF.SetFile(datastoreInitGO)

	/* GUI */
	// Folders
	var pages = assets.NewFolder(FolderNameGUIPages)
	pages.State = assets.StateChanged
	guiF.SetDependency(pages)
	var landings = assets.NewFolder(FolderNameGUILandings)
	landings.State = assets.StateChanged
	guiF.SetDependency(landings)
	var widgets = assets.NewFolder(FolderNameGUIWidgets)
	widgets.State = assets.StateChanged
	guiF.SetDependency(widgets)

	// /gui/main.js

	// /gui/init.js
	// /gui/init.json

	// /gui/landings/splash.html
	// /gui/landings/splash.css
	// /gui/landings/splash.json
	var splashHTML = &assets.File{}
	var splashCSS = &assets.File{}
	var splashJSON = &assets.File{}
	err = wg.MakeSplashFiles(splashHTML, splashCSS, splashJSON)
	if err != nil {
		return nil, err
	}
	req.Repo.SetFile(splashHTML)
	req.Repo.SetFile(splashCSS)
	req.Repo.SetFile(splashJSON)

	/* SDK */
	var goSDKRepo = assets.NewFolder("sdk-go")
	goSDKRepo.State = assets.StateChanged
	req.Repo.SetDependency(goSDKRepo)
	var jsSDKRepo = assets.NewFolder("sdk-js")
	jsSDKRepo.State = assets.StateChanged
	req.Repo.SetDependency(jsSDKRepo)

	/* Common files */
	// main.go
	var MainGo = &assets.File{}
	err = ag.MakeMainFile(MainGo)
	if err != nil {
		return nil, err
	}
	req.Repo.SetFile(MainGo)
	// connections.go
	var connectionsGo = &assets.File{}
	err = ag.MakeConnectionsFile(connectionsGo)
	if err != nil {
		return nil, err
	}
	req.Repo.SetFile(connectionsGo)

	// .gitignore
	var ob3 assets.File
	ob3.Name = ".gitignore"
	ob3.Data = *(*[]byte)(unsafe.Pointer(&gitignore))
	ob3.State = assets.StateChanged
	req.Repo.SetFile(&ob3)

	return res, nil
}

var gitignore = `
# Secrets files in secret folders
secret/

# Prerequisites
*.d

# Object files
*.o
*.ko
*.obj
*.elf

# Linker output
*.ilk
*.map
*.exp

# Precompiled Headers
*.gch
*.pch

# Libraries
*.lib
*.a
*.la
*.lo
assets--g.go

# Shared objects (inc. Windows DLLs)
*.dll
*.so
*.so.*
*.dylib

# Executables
*.exe
*.out
*.app
*.i*86
*.x86_64
*.hex

# Debug files
*.dSYM/
*.su
*.idb
*.pdb

# Kernel Module Compile Results
*.mod*
*.cmd
.tmp_versions/
modules.order
Module.symvers
Mkfile.old
dkms.conf

# Architecture specific extensions/prefixes
*.[568vq]
[568vq].out

*.cgo1.go
*.cgo2.c
_cgo_defun.c
_cgo_gotypes.go
_cgo_export.*

_testmain.go

*.exe
*.prof

# external packages folder
datastore-files/
images/
vendor/
*node_modules*
*bower_components*
*bundle.*
*yarn-error\.log
*yarn\.lock

# Log data
/log*
*.log
`
