//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	"libgo/detail"
	"libgo/protocol"
)

const domainEnglish = "Command"

func init() {
	ErrServiceNotAcceptCLI.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Not Accept CLI").
		SetOverview("Requested service not accept CLI protocol in this server").
		SetUserNote("Try other server or contact support of the software").
		SetDevNote("It is so easy to implement CLI handler for a service! Take a time and do it!").
		SetTAGS([]string{})
	)
	ErrServiceCallByAlias.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Call By Alias").
		SetOverview("We found given service name in aliases that don't support to serve by it.").
		SetUserNote("").
		SetDevNote("Try to call the command by its name or its abbreviation").
		SetTAGS([]string{})
	)
}
