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
	"./libgo/achaemenid"
	as "./libgo/achaemenid-services"
	"./libgo/ganjine"
	gs "./libgo/ganjine-services"
	"./libgo/letsencrypt"
	"./libgo/log"
	ps "./services"
)

// Server is just address of Achaemenid DefaultServer for easily usage
var server *achaemenid.Server = achaemenid.DefaultServer

var cluster *ganjine.Cluster

func init() {
	var err error

	server.Manifest = achaemenid.Manifest{
		SocietyID:           0,
		AppID:               [16]byte{},
		DomainID:            [16]byte{},
		DomainName:              "",
		Email:               "",
		Icon:                "",
		AuthorizedAppDomain: [][16]byte{},
		SupportedLanguages:  []uint32{},
		ManifestLanguages:   []uint32{},
		Organization:        []string{
			"",
		},
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
		TechnicalInfo:       achaemenid.TechnicalInfo{
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
			MaxNodeNumber:                3,

			TransactionTimeOut: 500,
			NodeFailureTimeOut: 60,
		},
	}

	// Initialize server
	server.Init()

	// Register stream app layer protocols. Dev can remove below and register only needed protocols handlers.
	server.StreamProtocols.Init()

	// register networks.
	err = server.Networks.Init(server)
	if err != nil {
		log.Fatal(err)
	}

	// Register some selectable networks. Check to add or delete networks to this function.
	selectableNetworks()

	// Connect to other nodes or become first node!
	server.Nodes.Init(server)
	if err != nil {
		log.Fatal(err)
	}

	// Register default Achaemenid services
	as.Init(server)
	// Register default Ganjine services
	gs.Init(server)
	// Register platform defined custom service in ./services/ folder
	ps.Init(server)

	// Register some other services for Achaemenid
	server.Connections.GetConnByID = getConnectionsByID
	server.Connections.GetConnByUserID = getConnectionsByUserID

	// Ganjine initialize
	err = cluster.Init(server)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize datastore
	datastore.Init(server, cluster)
}

func main() {
	var err error
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func selectableNetworks() {
	var err error

	// Check LetsEncrypt Certificate
	err = letsencrypt.CheckByAchaemenid(server)
	if err != nil {
		log.Fatal(err)
	}

	// Delete below network if you don't want to listen for HTTPs protocol.
	log.Info("try to register TCP on port 80 to listen for HTTP protocol")
	server.StreamProtocols.SetProtocolHandler(80, achaemenid.HTTPToHTTPSHandler)
	err = achaemenid.MakeTCPNetwork(server, 80)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("try to register TCP/TLS on port 443 to listen for HTTPs protocol")
	server.StreamProtocols.SetProtocolHandler(443, achaemenid.HTTPIncomeRequestHandler)
	err = achaemenid.MakeTCPTLSNetwork(server, 443)
	if err != nil {
		log.Fatal(err)
	}

	// Delete below network if you don't want to listen for DNS protocol.
	log.Info("try to register UDP on port 53 to listen for DNS protocol")
	server.StreamProtocols.SetProtocolHandler(53, achaemenid.DNSIncomeRequestHandler)
	err = achaemenid.MakeUDPNetwork(server, 53)
	if err != nil {
		log.Fatal(err)
	}
}
`))
