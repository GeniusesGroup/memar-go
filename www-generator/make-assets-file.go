/* For license and copyright information please see LEGAL file in repository */

package wg

import (
	"fmt"
	"unsafe"

	"../assets"
	"../www"
)

// MakeAssetsFile generate a go file with combine & inline all assets files!
func MakeAssetsFile(Repo *assets.Folder, file *assets.File) (err error) {
	file.FullName = "assets--www.go"
	file.Name = "assets--www"
	file.Extension = "go"

	var ass = assets.NewFolder("")
	www.AddRepoToAsset(ass, Repo)

	for _, file := range ass.Files {
		addFile(file)
	}

	// Add ending file
	assetsFile += "\n}\n"

	file.Data = *(*[]byte)(unsafe.Pointer(&assetsFile))
	file.State = assets.StateChanged

	return
}

func addFile(file *assets.File) {
	var ln = len(file.CompressData)
	var data = make([]byte, 0, ln*3)
	for i := 0; i < ln; i++ {
		data = append(data, fmt.Sprintf("%v", file.CompressData[i])+","...)
	}

	assetsFile += `Server.Assets.Files["` + file.FullName +
		`"] = &assets.File{FullName: "` + file.FullName +
		`", Name: "` + file.Name +
		`", Extension: "` + file.Extension +
		`", MimeType: "` + file.MimeType +
		`", CompressType: ` + file.CompressType +
		`", CompressData: []byte{` + *(*string)(unsafe.Pointer(&data)) + "}, } \n"
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
