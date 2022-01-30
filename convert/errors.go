/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "Convert"
const domainPersian = "تبدیل"

// Errors
var (
	ErrEmptyValue = er.New("urn:giti:convert.protocol:error:empty-value").SetDetail(protocol.LanguageEnglish, domainEnglish, "Empty Value",
		"Empty value pass to convert function that it is illegal",
		"",
		"").Save()

	ErrValueOutOfRange = er.New("urn:giti:convert.protocol:error:value out of range").SetDetail(protocol.LanguageEnglish, domainEnglish, "Value Out of Range",
			"indicates that a value is out of range for the target type, e.g. 270 for uint8",
			"",
			"").Save()

	ErrBadValue = er.New("urn:giti:convert.protocol:error:bad-value").SetDetail(protocol.LanguageEnglish, domainEnglish, "Bad Value",
		"Bad value pass to convert function that it is illegal e.g. pass '1b2' to convert to number",
		"",
		"").Save()
)
