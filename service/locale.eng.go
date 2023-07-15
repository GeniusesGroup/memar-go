//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"libgo/detail"
	"libgo/protocol"
)

const domainEnglish = "Service"

func init() {
	ErrNotFound.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Not Found").
		SetOverview("Requested service by given identifier not found in this application").
		SetUserNote("").
		SetDevNote("").
		SetTAGS([]string{})
	)

	ErrServiceNotAcceptSRPC.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Not Accept sRPC").
		SetOverview("Requested service by given ID not accept sRPC protocol in this server").
		SetUserNote("Try other server or contact support of the software").
		SetDevNote("It is so easy to implement sRPC handler for a service! Take a time and do it!").
		SetTAGS([]string{})
	)

	ErrServiceNotAcceptSRPCDirect.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Not Accept Direct sRPC").
		SetOverview("Requested service by given ID not accept direct sRPC protocol in this server").
		SetUserNote("Try other server or contact support of the software").
		SetDevNote("").
		SetTAGS([]string{})
	)

	ErrServiceNotAcceptHTTP.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Not Accept HTTP").
		SetOverview("Requested service by given ID not accept HTTP protocol in this server").
		SetUserNote("Try other server or contact support of the software").
		SetDevNote("It is so easy to implement HTTP handler for a service! Take a time and do it!").
		SetTAGS([]string{})
	)

	ErrServiceNotProvideIdentifier.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Not Provide Identifier").
		SetOverview("Service must have a valid URI or mediatype. It is rule to add more detail about service.").
		SetUserNote("").
		SetDevNote("Initialize inner s.MediaType.Init() first if use libgo/service package").
		SetTAGS([]string{})
	)

	ErrServiceDuplicateIdentifier.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Duplicate Identifier").
		SetOverview("ID or MediaType or URI associated for requested register service used before for other service and not legal to reuse same identifier for other services").
		SetUserNote("").
		SetDevNote("Check the services for duplicate identifier assigned").
		SetTAGS([]string{})
	)
}
