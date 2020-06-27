/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	"../assets"
)

// MakeServicesInitFile use to make init.go file to register Achaemenid services files!
func MakeServicesInitFile(file *assets.File) (err error) {
	file.FullName = "init.go"
	file.Name = "init"
	file.Extension = "go"
	file.State = assets.StateChanged

	var mt = new(bytes.Buffer)
	if err = servicesInitFile.Execute(mt, nil); err != nil {
		return
	}

	file.Data, err = format.Source(mt.Bytes())
	if err != nil {
		return
	}

	return
}

var servicesInitFile = template.Must(template.New("main").Parse(`
/* For license and copyright information please see LEGAL file in repository */
// Auto-generated, edits will be overwritten

package services

import "../../libgo/achaemenid"

// Init use to register all available services to given achaemenid.
func Init(s *achaemenid.Server) {

	// Common Services
	// s.Services.RegisterService(&)
}

`))
