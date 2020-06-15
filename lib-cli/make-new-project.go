/* For license and copyright information please see LEGAL file in repository */

package main

import (
	ag "../Achaemenid-generator"
	gg "../Ganjine-generator"
	"../assets"
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
	var apisF = assets.NewFolder(FolderNameAPIs)
	apisF.State = assets.StateChanged
	req.Repo.SetDependency(apisF)
	var dbF = assets.NewFolder(FolderNameDB)
	dbF.State = assets.StateChanged
	req.Repo.SetDependency(dbF)
	var guiF = assets.NewFolder(FolderNameGUI)
	guiF.State = assets.StateChanged
	req.Repo.SetDependency(guiF)

	/* APIs */
	// Folders
	var services = assets.NewFolder(FolderNameServices)
	services.State = assets.StateChanged
	apisF.SetDependency(services)
	var datastore = assets.NewFolder(FolderNameDataStore)
	datastore.State = assets.StateChanged
	apisF.SetDependency(datastore)

	/* DB */
	// /db/main.go
	var DBMainGo = &assets.File{}
	err = gg.MakeMainFile(DBMainGo)
	if err != nil {
		return nil, err
	}
	dbF.SetFile(DBMainGo)

	/* GUI */
	// Folders
	var pages = assets.NewFolder(FolderNameGGPages)
	pages.State = assets.StateChanged
	guiF.SetDependency(pages)
	var landings = assets.NewFolder(FolderNameGUILandings)
	landings.State = assets.StateChanged
	guiF.SetDependency(landings)
	var widgets = assets.NewFolder(FolderNameGUIWidgets)
	widgets.State = assets.StateChanged
	guiF.SetDependency(widgets)

	// /gui/main.html

	// /gui/main.js

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

	// .gitignore
	var ob3 assets.File
	ob3.Name = ".gitignore"
	ob3.Data = []byte(gitignore)
	ob3.State = assets.StateChanged
	req.Repo.SetFile(&ob3)

	return res, nil
}

const gitignore = `
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
vendor/
*node_modules*
*bower_components*
*bundle.*
*yarn-error\.log
*yarn\.lock

# Log data
*.log
`
