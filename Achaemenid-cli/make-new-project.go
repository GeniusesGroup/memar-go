/* For license and copyright information please see LEGAL file in repository */

package main

import (
	generator "../Achaemenid-generator"
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
	var apis = assets.NewFolder(FolderNameAPIs)
	apis.Status = assets.StateChanged
	req.Repo.SetDependency(apis)
	var db = assets.NewFolder(FolderNameDB)
	db.Status = assets.StateChanged
	req.Repo.SetDependency(db)
	var gui = assets.NewFolder(FolderNameGUI)
	gui.Status = assets.StateChanged
	req.Repo.SetDependency(gui)
	var www = assets.NewFolder(FolderNameWWW)
	www.Status = assets.StateChanged
	req.Repo.SetDependency(www)

	/* APIs */
	// Folders
	var services = assets.NewFolder(FolderNameAPIsServices)
	services.Status = assets.StateChanged
	apis.SetDependency(services)
	var datastore = assets.NewFolder(FolderNameAPIsDataStore)
	datastore.Status = assets.StateChanged
	apis.SetDependency(datastore)

	// /apis/main.go
	var mainGo = &assets.File{}
	err = generator.MakeMainFile(mainGo)
	if err != nil {
		return nil, err
	}
	apis.SetFile(mainGo)

	/* GUI */
	// Folders
	var pages = assets.NewFolder(FolderNameGGPages)
	pages.Status = assets.StateChanged
	gui.SetDependency(pages)
	var landings = assets.NewFolder(FolderNameGUILandings)
	landings.Status = assets.StateChanged
	gui.SetDependency(landings)
	var widgets = assets.NewFolder(FolderNameGUIWidgets)
	widgets.Status = assets.StateChanged
	gui.SetDependency(widgets)

	// /gui/main.html

	// /gui/main.js

	/* Common files */
	// .gitignore
	var ob3 assets.File
	ob3.Name = ".gitignore"
	ob3.Data = []byte(gitignore)
	ob3.Status = assets.StateChanged
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

# ChaparKhane files
assets--g.go
`
