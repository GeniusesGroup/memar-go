//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package validators

const domainEnglish = "Validation"

func init() {
	TextOverFlow.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Text OverFlow",
		"Given text size is more than it must be",
		"",
		"",
		nil)
	TextLack.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Text Lack",
		"Given text size is less than it must be",
		"",
		"",
		nil)
	TextIllegalCharacter.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Text Illegal Character",
		"Given text include illegal character or text!",
		"",
		"",
		nil)
}
