//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "Service"

func init() {
	ErrNotFound.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Not Found").
		SetOverview("Requested service by given identifier not found in this application").
		SetUserActionNote("").
		SetDevActionNote("").
		SetTAGS([]string{})
	)

	ErrServiceNotProvideIdentifier.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Not Provide Identifier").
		SetOverview("Service must have a valid ID and mediatype. It is rule to add more detail about service.").
		SetUserActionNote("").
		SetDevActionNote("Initialize inner s.MediaType.Init() first if use memar/service package").
		SetTAGS([]string{})
	)

	ErrServiceDuplicateIdentifier.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Duplicate Identifier").
		SetOverview("ID or MediaType associated for requested register service used before for other service and not legal to reuse same identifier for other services").
		SetUserActionNote("").
		SetDevActionNote("Check the services for duplicate identifier assigned").
		SetTAGS([]string{})
	)
}
