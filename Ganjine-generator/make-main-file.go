/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	"../assets"
)

// MakeMainFile use to make main file to start Achaemenid sever for www!
func MakeMainFile(file *assets.File) (err error) {
	file.FullName = "main.go"
	file.Name = "main"
	file.Extension = "go"
	file.State = assets.StateChanged

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
	"./libgo/achaemenid"
	as "../libgo/achaemenid-services"
	"../libgo/Ganjine"
	gs "../libgo/Ganjine-services/services"
)

const (
	replicationNumber = 3
	NodeNumber = 3
)

// Server is just address of Achaemenid DefaultServer for easily usage
var server = achaemenid.DefaultServer

var Cluster *ganjine.Cluster

func init() {
	var err error

	server.Init()

	server.Manifest = achaemenid.Manifest{
		AppID:               [16]byte{},
		Domain:              "db.",
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

	// Register stream app layer protocols
	server.StreamProtocols.SetProtocolHandler(4, achaemenid.SrpcIncomeRequestHandler)

	// register networks.
	err = achaemenid.MakeGPNetwork(server)
	if err != nil {
		panic(err)
	}
	err = achaemenid.MakeTCPTLSNetwork(server, 4, achaemenid.SrpcIncomeRequestHandler)
	if err != nil {
		panic(err)
	}

	// Register default Achaemenid services
	as.Init(server)
	// Register default Ganjine services
	gs.Init(server)

	err = Clueter.Get(server.Manifest.AppID)
	if err != nil {
		err = Cluster.GetBySDK()
		if err != nil {
			Cluster.Init(replicationNumber, nodeNumber)
		}
	}
}

func main() {
	var err error
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
`))
