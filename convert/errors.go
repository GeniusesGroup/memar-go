/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "Convert"
const errorPersianDomain = "تبدیل"

// Errors
var (
	ErrEmptyValue = er.New("urn:giti:convert.libgo:error:empty-value").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Empty Value",
		"Empty value pass to convert function that it is illegal",
		"",
		"").Save()

	ErrBadValue = er.New("urn:giti:convert.libgo:error:bad-value").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Bad Value",
		"Bad value pass to convert function that it is illegal e.g. pass '1b2' to convert to number",
		"",
		"").Save()
)
