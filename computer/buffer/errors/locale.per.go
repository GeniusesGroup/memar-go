//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_errs

import (
	"memar/detail"
	"memar/protocol"
)

const domainPersian = "بافر"

func init() {
	TooLarge.SetDetail(detail.New(protocol.LanguagePersian, domainPersian).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("خیلی بزرگ").
		SetOverview("").
		SetUserActionNote("").
		SetDevActionNote("").
		SetTAGS([]string{})
	)
}
