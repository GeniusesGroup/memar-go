/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	"../assets"
)

// MakeDataStoreInitFile make datastore init file to initialize some data to start a ganjine node!
func MakeDataStoreInitFile(file *assets.File) (err error) {
	file.FullName = "init.go"
	file.Name = "init"
	file.Extension = "go"
	file.State = assets.StateChanged

	var mt = new(bytes.Buffer)
	if err = initFileTemplate.Execute(mt, ""); err != nil {
		return
	}

	file.Data, err = format.Source(mt.Bytes())
	if err != nil {
		return
	}

	return
}

var initFileTemplate = template.Must(template.New("initFile").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package datastore

import (
	"../libgo/achaemenid"
	"../libgo/ganjine"
)

var (
	// Cluster store cluster data to use by services!
	cluster *ganjine.Cluster
	// Server store address location to server use by other part of app!
	server *achaemenid.Server
)

// Init must call in main file before use any methods!
func Init(s *achaemenid.Server, c *ganjine.Cluster) {
	server = s
	cluster = c
}

`))
