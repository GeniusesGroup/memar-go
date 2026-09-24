/* For license and copyright information please see the LEGAL file in the code repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	file_p "memar/storage/file/protocol"
)

// MakeServicesInitFile use to make init.go file to register Achaemenid services files!
func MakeServicesInitFile(file file_p.File) (fileData []byte, err error) {
	// file.Rename("init.go")

	var mt = new(bytes.Buffer)
	if err = servicesInitFile.Execute(mt, nil); err != nil {
		return
	}

	fileData, err = format.Source(mt.Bytes())
	if err != nil {
		return
	}

	return
}

var servicesInitFile = template.Must(template.New("main").Parse(`
/* For license and copyright information please see the LEGAL file in the code repository */
// Auto-generated, edits will be overwritten

package services

import (
	"memar/protocol"
)

const (
	DomainName = ""
)

// Init use to register all available services to given app.
func Init(app protocol.Application) {

	// Common Services
	// app.RegisterService(&)
}

`))
