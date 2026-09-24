/* For license and copyright information please see the LEGAL file in the code repository */

package generator

import (
	"bytes"
	"text/template"
)

// MakeDataStoreInitFile make datastore init file to initialize some data to start a ganjine node!
func MakeDataStoreInitFile(file file_p.File) (err error) {
	file.Rename("init.go")

	var mt = new(bytes.Buffer)
	if err = initFileTemplate.Execute(mt, ""); err != nil {
		return
	}

	// file.Data, err = format.Source(mt.Bytes())
	if err != nil {
		return
	}

	return
}

var initFileTemplate = template.Must(template.New("initFile").Parse(`
/* For license and copyright information please see the LEGAL file in the code repository */

package data

import (
	"memar/memar/achaemenid"
	"memar/memar/ganjine"
	"memar/memar/matn"
	"memar/memar/pehrest"
)

// Init must call in main file before use any methods!
func Init() (err error_p.Error) {
	server = s
	cluster = c

	pehrest.Init(s, c)
	matn.Init(s, c)
}

`))
