/* For license and copyright information please see LEGAL file in repository */

package main

import (
	generator "../ChaparKhane-generator"
	parser "../ChaparKhane-parser"
)

// MakeNewProjectReq :
type MakeNewProjectReq struct{
	Repo *parser.Repository
}

// MakeNewProjectRes :
type MakeNewProjectRes struct {}

// MakeNewProject :
func MakeNewProject(req *MakeNewProjectReq) (res *MakeNewProjectRes, err error) {
	// Folders
	var platformServices = parser.NewRepository()
	req.Repo.Dependencies[PlatformServicesFolderName] = platformServices
	var platformDatastore = parser.NewRepository()
	req.Repo.Dependencies[PlatfromDataStoreFolderName] = platformDatastore

	// /main.go
	var mainRes *generator.MakeMainFileRes
	mainRes, err = generator.MakeMainFile(&generator.MakeMainFileReq{})
	if err != nil {
		return nil, err
	}
	var ob2 parser.File
	ob2.Name = mainRes.MainFileName
	ob2.Data = mainRes.MainFile

	req.Repo.Files[ob2.Name] = &ob2

	// .gitignore
	var ob3 parser.File
	ob3.Name = ".gitignore"
	ob3.Data = []byte(gitignore)
	req.Repo.Files[ob3.Name] = &ob3

	return res, nil
}

var gitignore = `
# Compiled Object files, Static and Dynamic libs (Shared Objects)
*.o
*.a
*.so

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

# Log data
*.log

# Folders
libgo
_obj
_test
build
bin
pkg

# ChaparKhane code generated
assets.go

# use this file to store temporarily commit message due vsc clear message on exit!
___CommitMessage___
`
