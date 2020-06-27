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
	"./datastore"
	ps "./services"
	gs "./libgo/Ganjine-services"
	"./libgo/achaemenid"
	as "./libgo/achaemenid-services"
	"./libgo/assets"
	"./libgo/ganjine"
	"./libgo/www"
)

// Server is just address of Achaemenid DefaultServer for easily usage
var server = achaemenid.DefaultServer

var cluster *ganjine.Cluster

func init() {
	var err error

	server.Init()

	server.Manifest = achaemenid.Manifest{
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

	// Register stream app layer protocols
	server.StreamProtocols.SetProtocolHandler(4, achaemenid.SrpcIncomeRequestHandler)
	server.StreamProtocols.SetProtocolHandler(80, achaemenid.HTTPIncomeRequestHandler)
	server.StreamProtocols.SetProtocolHandler(443, achaemenid.HTTPSIncomeRequestHandler)

	// register networks.
	err = achaemenid.MakeGPNetwork(server)
	if err != nil {
		panic(err)
	}
	err = achaemenid.MakeTCPTLSNetwork(server, 4, achaemenid.SrpcIncomeRequestHandler)
	if err != nil {
		panic(err)
	}
	err = achaemenid.MakeTCPNetwork(server, 80, achaemenid.HTTPIncomeRequestHandler)
	if err != nil {
		panic(err)
	}
	err = achaemenid.MakeTCPTLSNetwork(server, 443, achaemenid.HTTPSIncomeRequestHandler)
	if err != nil {
		panic(err)
	}
	// Comment other networks and uncomment below network in development phase!
	// err = achaemenid.MakeTCPNetwork(server, 8080, achaemenid.HTTPSIncomeRequestHandler)
	// if err != nil {
	// 	panic(err)
	// }
	// server.Assets = assets.NewFolder(server.Manifest.Domain)
	// go www.ReloadAssetsInDevPhase(server.Assets)

	// Register default Achaemenid services
	as.Init(server)
	// Register default Ganjine services
	gs.Init(server)
	// Register platform defined custom service in ./services/ folder
	ps.Init(server)

	err = cluster.Init(server)
	if err != nil {
		panic(err)
	}

	// Initialize datastore
	datastore.Init(server, cluster)
}

func main() {
	var err error
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
`))
