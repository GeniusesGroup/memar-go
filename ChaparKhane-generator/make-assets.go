/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"../assets"
)

// MakeAssetsReq is request structure of MakeAssets()
type MakeAssetsReq struct {
	Repo *assets.Folder
}

// MakeAssetsRes is response structure of MakeAssets()
type MakeAssetsRes struct {
	AssetsFileName string
	AssetsFile     []byte
}

// MakeAssets :
func MakeAssets(req *MakeAssetsReq) (res *MakeAssetsRes, err error) {
	res = &MakeAssetsRes{
		AssetsFileName: "assets--g.go",
	}

	var mt = new(bytes.Buffer)
	err = assetsTemplate.Execute(mt, req.Repo)
	if err != nil {
		return nil, err
	}

	res.AssetsFile, err = format.Source(mt.Bytes())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func printAssets(repo *assets.Folder, depLocation string) (printed string) {
	for fileName, file := range repo.Files {
		filePrinted := printFile(fileName, file.Data)
		printed = printed + depLocation + `.Files["` + fileName + `"] =` + filePrinted + "\n"
	}

	printed = printed + "\n"

	for _, repo := range repo.Dependencies {
		innerLocation := depLocation + `.Dependencies["` + repo.Name + `"]`
		printed = printed + innerLocation + ".Files = make(map[string]*chaparkhane.AssetsFile)" + "\n"
		printed = printed + innerLocation + ".Dependencies = make(map[string]*chaparkhane.Assets)" + "\n"
		innerAssets := printAssets(repo, innerLocation)
		printed = printed + innerAssets + "\n"
	}

	return printed
}

func printFile(fileName string, fileDate []byte) (print string) {
	data := "[]byte{"
	for _, b := range fileDate {
		data = data + fmt.Sprintf("%v", b) + ","
	}
	data = data + "}"

	print = fmt.Sprintf(`&chaparkhane.AssetsFile{Name: "%s", Data: %s}`, fileName, data)

	return print
}

var assetsTemplate = template.Must(template.New("assetsTemplate").Funcs(template.FuncMap{
	"printAssets": printAssets,
	"printFile":   printFile,
}).Parse(`
// Code generated .* DO NOT EDIT.$
/* For license and copyright information please see LEGAL file in ChaparKhane repository */

package main

import (
	chaparkhane "./ChaparKhane"
)

// Assets use to store all assets data and structure
var Assets = chaparkhane.NewAssets()

func init(){
	chaparkhane.DefaultServer.Assets = Assets.Dependencies["assets"]

	Assets.Dependencies["assets"].Files = make(map[string]*chaparkhane.AssetsFile)
	Assets.Dependencies["assets"].Dependencies = make(map[string]*chaparkhane.Assets)
	{{range .Dependencies}}
		{{if (eq .Name "assets")}} 
			{{printAssets . "Assets.Dependencies[` + "`" + `assets` + "`" + `]"}}
		{{- end}}
	{{end}}
}

`))
