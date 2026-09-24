//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package file

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "File"

func init() {
	MediaType.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("File URI").
		SetOverview("").
		SetUserActionNote("").
		SetDevActionNote("").
		SetTAGS([]string{})
	)
}
