/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"fmt"

	"../assets"
)

// MakeAssetsFile generate a go file with combine & inline all assets files!
func MakeAssetsFile(Repo *assets.Folder, file *assets.File) (err error) {
	file.FullName = "assets--www.go"
	file.Name = "assets--www"
	file.Extension = "go"

	var ass = assets.NewFolder("")
	addRepoToAsset(ass, Repo)

	for _, file := range ass.Files {
		addFile(file)
	}

	// Add ending file
	assetsFile += "\n}\n"

	file.Data = []byte(assetsFile)
	file.State = assets.StateChanged

	return
}

func addFile(file *assets.File) {
	var ln = len(file.Data)
	var data = make([]byte, 0, ln*3)
	for i := 0; i < ln; i++ {
		data = append(data, fmt.Sprintf("%v", file.Data[i])+","...)
	}

	assetsFile += `Server.Assets.Files["` + file.FullName +
		`"] = &assets.File{FullName: "` + file.FullName +
		`", Name: "` + file.Name +
		`", Extension: "` + file.Extension +
		`", MimeType: "` + file.MimeType +
		`", Data: []byte{` + string(data) + "}, } \n"
}

var assetsFile = `
// Code generated .* DO NOT EDIT.$
/* For license and copyright information please see LEGAL file in Achaemenid repository */

package main

import (
	"./libgo/assets"
)

func init(){
	Server.Assets = assets.NewFolder("")

`
