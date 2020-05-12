/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	"../assets"
)

// MakeMainFile use to make main file to start ChaparKhane sever
func MakeMainFile(file *assets.File) (err error) {
	file.FullName = "main.go"
	file.Name = "main"
	file.Extension = "go"
	file.Status = assets.StateChanged

	var mt = new(bytes.Buffer)
	if err = mainFileTemplate.Execute(mt, ""); err != nil {
		return
	}

	file.Data, err = format.Source(mt.Bytes())
	if err != nil {
		return
	}

	return
}

var mainFileTemplate = template.Must(template.New("main").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"../libgo/achaemenid"
	ss "../libgo/achaemenid-services"
	"./services"
)

func init() {
	achaemenid.DefaultServer.Init()

	ss.Init(achaemenid.DefaultServer)
	services.Init(achaemenid.DefaultServer)

	achaemenid.DefaultServer.Manifest = achaemenid.Manifest{
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
		TechnicalInfo:       achaemenid.TechnicalInfo{},
	}
}

func main() {
	var err error
	err = achaemenid.DefaultServer.Start()
	if err != nil {
		panic(err)
	}
}
`))
