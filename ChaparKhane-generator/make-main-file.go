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
	chaparkhane "../libgo/ChaparKhane"
	ss "../libgo/ChaparKhane-services"
	services "./platform-services"
)

func init() {
	chaparkhane.DefaultServer.Init()

	ss.Init(chaparkhane.DefaultServer)
	services.Init(chaparkhane.DefaultServer)

	chaparkhane.DefaultServer.Manifest = chaparkhane.Manifest{
		AppID:               [16]byte{},
		Domain:              "",
		Email:               "",
		Icon:                "",
		AuthorizedAppDomain: []string{},
		SupportedLanguages:  []uint32{},
		ManifestLanguages:   []uint32{},
		Name: []string{
			"",
			"",
		},
		Description: []string{
			"",
			"",
		},
		TermsOfService: []string{
			"",
		},
		Licence: []string{
			"",
		},
		TAGS: []string{
			"",
		},
		RequestedPermission: []uint32{},
		TechnicalInfo:       chaparkhane.TechnicalInfo{},
	}
}

func main() {
	var err error
	err = chaparkhane.DefaultServer.Start()
	if err != nil {
		panic(err)
	}
}
`))
