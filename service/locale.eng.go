//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainEnglish = "Service"

func init() {
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested service by given identifier not found in this application",
		"",
		"",
		nil)

	ErrServiceNotAcceptSRPC.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept sRPC",
		"Requested service by given ID not accept sRPC protocol in this server",
		"Try other server or contact support of the software",
		"It is so easy to implement sRPC handler for a service! Take a time and do it!",
		nil)

	ErrServiceNotAcceptSRPCDirect.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept Direct sRPC",
		"Requested service by given ID not accept direct sRPC protocol in this server",
		"Try other server or contact support of the software",
		"",
		nil)

	ErrServiceNotAcceptHTTP.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept HTTP",
		"Requested service by given ID not accept HTTP protocol in this server",
		"Try other server or contact support of the software",
		"It is so easy to implement HTTP handler for a service! Take a time and do it!",
		nil)
}
