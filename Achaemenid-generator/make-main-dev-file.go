/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	"../assets"
)

// MakeMainDevFile use to make main-dev.go file to start Achaemenid sever for development phase!
func MakeMainDevFile(file *assets.File) (err error) {
	file.FullName = "main-dev.go"
	file.Name = "main-dev"
	file.Extension = "go"
	file.State = assets.StateChanged

	var mt = new(bytes.Buffer)
	if err = mainDevFileTemplate.Execute(mt, ""); err != nil {
		return
	}

	file.Data, err = format.Source(mt.Bytes())
	if err != nil {
		return
	}

	return
}

var mainDevFileTemplate = template.Must(template.New("main-dev").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"../libgo/achaemenid"
	as "../libgo/achaemenid-services"
	gs "../libgo/ganjine-services"
	"../libgo/log"
	"../libgo/www"
	ps "../services"
)

// Server is just address of Achaemenid DefaultServer for easily usage
var devServer *achaemenid.Server = achaemenid.DefaultServer

func init() {
	var err error

	devServer.Manifest = achaemenid.Manifest{
		RequestedPermission: []uint32{},
		TechnicalInfo: achaemenid.TechnicalInfo{
			GuestMaxConnections: 2000000,

			CPUCores: 1,                        // one core
			CPUSpeed: 1 * 1000 * 1000,          // 1 GHz
			RAM:      4 * 1024 * 1024 * 1024,   // 4 GB
			GPU:      0,                        // 0 Ghz
			Network:  100 * 1024 * 1024,        // 100 MB/S
			Storage:  100 * 1024 * 1024 * 1024, // 100 GB

			DistributeOutOfSociety:       false,
			DataCentersClass:             5,
			DataCentersClassForDataStore: 0,
			ReplicationNumber:            3,
			MaxNodeNumber:                30,

			TransactionTimeOut: 500,
			NodeFailureTimeOut: 60,
		},
	}

	// Initialize devServer
	devServer.Init()

	// Register stream app layer protocols. Dev can remove below and register only needed protocol.
	devServer.StreamProtocols.Init()

	log.Info("try to register TCP on port 8080 to listen for HTTP protocol in dev phase")
	devServer.StreamProtocols.SetProtocolHandler(8080, achaemenid.HTTPIncomeRequestHandler)
	err = achaemenid.MakeTCPNetwork(devServer, 8080)
	if err != nil {
		log.Fatal(err)
	}
	go www.ReloadAssetsFromStorage(devServer.Assets.WWW)

	// Register default Achaemenid devServer
	as.Init(devServer)
	// Register platform defined custom service in ./services/
	ps.Init(devServer)
	// Register default Ganjine services
	gs.Init(devServer)

	// Register some other services for Achaemenid
	// devServer.Connections.GetConnByID = getConnectionsByID
	// devServer.Connections.GetConnByUserID = getConnectionsByUserID
}

func main() {
	var err error
	err = devServer.Start()
	if err != nil {
		log.Fatal(err)
	}
}

`))
