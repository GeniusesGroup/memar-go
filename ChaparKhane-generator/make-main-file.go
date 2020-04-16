/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"
)

// MakeMainFileReq is request structure of MakeMainFile()
type MakeMainFileReq struct{}

// MakeMainFileRes is response structure of MakeMainFile()
type MakeMainFileRes struct {
	MainFileName string
	MainFile     []byte
}

// MakeMainFile use to make main file to start ChaparKhane sever
func MakeMainFile(req *MakeMainFileReq) (res *MakeMainFileRes, err error) {
	res = &MakeMainFileRes{
		MainFileName: "main.go",
		MainFile:     []byte{},
	}

	var mt = new(bytes.Buffer)
	if err = mainFileTemplate.Execute(mt, ""); err != nil {
		return nil, err
	}

	res.MainFile, err = format.Source(mt.Bytes())
	if err != nil {
		return nil, err
	}

	return res, nil
}

var mainFileTemplate = template.Must(template.New("main").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"encoding/json"

	chaparkhane "./libgo/ChaparKhane"
)

func init() {
	if manifest, ok := Assets.Files["server.manifest.json"]; ok {
		// Manifest : All needed Data to operation the service(application)
		err := json.Unmarshal(manifest.Data, &chaparkhane.Manifest)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Chaparkhane manifest not exist!!")
	}
}

func main() {
	if err := StartServer(); err != nil {
		panic(err)
	}
}
`))
