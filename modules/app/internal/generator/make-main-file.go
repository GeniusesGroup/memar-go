/* For license and copyright information please see the LEGAL file in the code repository */

package generator

import (
	"bytes"
	"go/format"
	file_p "memar/storage/file/protocol"
	"text/template"
)

// MakeMainFile use to make main file to start Achaemenid sever for www!
func MakeMainFile(file file_p.File) (fileData []byte, err error) {
	// file.Rename("main.go")

	var mt = new(bytes.Buffer)
	if err = mainFileTemplate.Execute(mt, ""); err != nil {
		return
	}

	fileData, err = format.Source(mt.Bytes())
	return
}

var mainFileTemplate = template.Must(template.New("main").Parse(`
/* For license and copyright information please see the LEGAL file in the code repository */

package main

import (
	"time"

	"./services"

	"memar/achaemenid"
	lang "memar/language"
	"memar/process/log"
	letsencrypt "github.com/GeniusesGroup/SDK/letsencrypt.org"
)

func init() {
	achaemenid.App.Manifest = achaemenid.Manifest{
		Society:    "",
		DomainName: "",
		Email:      "ict@??DomainName??",
		Icon:       "",

		Organization: map[lang_p.LanguageID]string{
			lang_p.LanguageEnglish: "",
		},
		Name: map[lang_p.LanguageID]string{
			lang_p.LanguageEnglish: "",
		},
		Description: map[lang_p.LanguageID]string{
			lang_p.LanguageEnglish: "",
		},
		TermsOfService: map[lang_p.LanguageID]string{
			lang_p.LanguageEnglish: "https://??DomainName??/terms?hl=en",
		},
		Licence: map[lang_p.LanguageID]string{
			lang_p.LanguageEnglish: "https://??DomainName??/licence?hl=en",
		},
		TAGS: []string{
			"", "", "",
		},
	
		RequestedPermission: []uint32{},
	}

	var err error_p.Error

	// Initialize server
	achaemenid.App.Init()

	// Register stream app layer protocols. Dev can remove below and register only needed protocols handlers.
	achaemenid.App.StreamProtocols.Init()

	services.Init(&server)

	if log.CNF_DevelopingMode {
		devInit()
	} else {
		productionInit()
	}

	// Connect to other nodes or become first node!
	fatal(achaemenid.App.Nodes.Init(&server))

	cluster.Manifest = ganjine.Manifest{
		DataCentersClass: 0,
		TotalZones:       3,

		TransactionTimeOut: 500,
		NodeFailureTimeOut: 60,

		CachePercent: 20,
	}

	// Ganjine initialize
	err = cluster.Init(&server)
	if err != nil {
		fatal(err)
	}
	// Initialize datastore
	data.Init(&server, &cluster)

	// Register some other services for Achaemenid
	achaemenid.App.Connections.GetConnByID = getConnectionsByID
	achaemenid.App.Connections.GetConnByUserIDThingID = getConnectionsByUserIDThingID
	achaemenid.App.Connections.SaveConn = saveConnection
}

func main() {
	achaemenid.App.Start()
}

func productionInit() {
	var err error

	// register networks.
	err = achaemenid.App.Networks.Init(&server)
	if err != nil {
		fatal(err)
	}

	// Check LetsEncrypt Certificate
	err = letsencrypt.CheckByAchaemenid(&server)
	if err != nil {
		fatal(err)
	}

	// Register some selectable networks. Check to add or delete networks.
	log.Info(domainEnglish, "try to register TCP on port 80 to listen for HTTP protocol")
	achaemenid.App.StreamProtocols.SetProtocolHandler(80, http.HTTPToHTTPSHandler)
	err = achaemenid.MakeTCPNetwork(server, 80)
	if err != nil {
		fatal(err)
	}
	log.Info(domainEnglish, "try to register TCP/TLS on port 443 to listen for HTTPs protocol")
	achaemenid.App.StreamProtocols.SetProtocolHandler(443, http.IncomeRequestHandler)
	err = achaemenid.MakeTCPTLSNetwork(&server, 443)
	if err != nil {
		fatal(err)
	}

	log.Info(domainEnglish, "try to register UDP on port 53 to listen for DNS protocol")
	achaemenid.App.StreamProtocols.SetProtocolHandler(53, dns.IncomeRequestHandler)
	err = achaemenid.MakeUDPNetwork(server, 53)
	if err != nil {
		fatal(err)
	}
}

func devInit() {
	var err error

	log.Info(domainEnglish, "try to register TCP on port 8080 to listen for HTTP protocol in dev phase")
	achaemenid.App.StreamProtocols.SetProtocolHandler(8080, http.IncomeRequestHandler)
	err = achaemenid.MakeTCPNetwork(&server, 8080)
	if err != nil {
		fatal(err)
	}

	// log.Info(domainEnglish, "try to register TCP/TLS on port 443 to listen for HTTPs protocol")
	// achaemenid.App.StreamProtocols.SetProtocolHandler(443, http.IncomeRequestHandler)
	// err = achaemenid.MakeTCPTLSNetwork(&server, 443)
	// if err != nil {
	// 	fatal(err)
	// }

	go achaemenid.App.Assets.ReLoadByCLI()
}

// registerNetworkApplicationHandlers use to register supported protocols!
// Usually Dev must register needed stream protocol not use this method to register all available protocols!
// Not rule but suggest to use even port number to listen||receive||response||server and odd for send||request||client
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
// https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml
func registerNetworkApplicationHandlers() {
	// sRPC
	nam.handlers[protocol.NetworkApplicationSRPC] = srpc.SRPCSyllab{}

	// SMTP
	// nam.handlers[25] = smtpHandler

	// DNS
	nam.handlers[protocol.NetworkApplicationDNS] = dns.IncomeRequestHandler

	// HTTP & HTTPS
	nam.handlers[protocol.NetworkApplicationHTTP] = http.HTTPHandler{}
}
`))
