//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package validators

import (
	"libgo/protocol"
)

const domainEnglish = "Validation"

func init() {
	ErrTextOverFlow.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Text OverFlow",
		"Given text size is more than it must be",
		"",
		"",
		nil)
	ErrTextLack.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Text Lack",
		"Given text size is less than it must be",
		"",
		"",
		nil)
	ErrTextIllegalCharacter.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Text Illegal Character",
		"Given text include illegal character or text!",
		"",
		"",
		nil)
}
