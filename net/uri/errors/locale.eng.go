//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

const domainEnglish = "URI"

func init() {
	ErrParse.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Parse URI",
		"Parsing received HTTP packet encounter unknown situation in URI part",
		"",
		"",
		nil)
	ErrQueryBadKey.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Query bad key",
		"invalid semicolon separator in query",
		"",
		"",
		nil)
	ErrInvalidURLEscape.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Invalid URL escape",
		"invalid URL escape in given uri",
		"",
		"",
		nil)
	ErrInvalidHostEscape.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Invalid host escape",
		"invalid character in host name",
		"",
		"",
		nil)
}

func init() {
	MediaType.SetDetail(lang_p.LanguageEnglish, "URI", "", "", "", "", []string{})
}
