//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package math_errs

import (
	"memar/detail"
)

const domainEnglish = "Math"

func init() {
	ErrEmptyValue.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Empty Value").
		SetOverview("Empty value pass to convert function that it is illegal").
		SetUserActionNote("").
		SetDevActionNote("").
		SetTAGS([]string{}))

	ErrValueOutOfRange.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Value Out of Range").
		SetOverview("indicates that a value is out of range for the target type, e.g. 270 for uint8").
		SetUserActionNote("").
		SetDevActionNote("").
		SetTAGS([]string{}))

	ErrBadValue.SetDetail(detail.New(lang_p.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Bad Value").
		SetOverview("Bad value pass to convert function that it is illegal e.g. pass '1b2' to convert to number").
		SetUserActionNote("").
		SetDevActionNote("").
		SetTAGS([]string{}))
}
