//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_errs

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "Buffer"

func init() {
	TooLarge.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Too Large").
		SetOverview("").
		SetUserActionNote("").
		SetDevActionNote("").
		SetTAGS([]string{})
	)
}
