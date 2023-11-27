//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "URI"

func init() {
	ErrParse.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse URI",
		"Parsing received HTTP packet encounter unknown situation in URI part",
		"",
		"",
		nil)
	ErrQueryBadKey.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Query bad key",
		"invalid semicolon separator in query",
		"",
		"",
		nil)
	ErrInvalidURLEscape.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Invalid URL escape",
		"invalid URL escape in given uri",
		"",
		"",
		nil)
	ErrInvalidHostEscape.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Invalid host escape",
		"invalid character in host name",
		"",
		"",
		nil)
}

func init() {
	MediaType.SetDetail(protocol.LanguageEnglish, "URI", "", "", "", "", []string{})
}
