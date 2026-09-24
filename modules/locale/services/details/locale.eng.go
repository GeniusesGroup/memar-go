//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package locale

const (
	domainEnglish = "locale"
)

func init() {
	LocaleCommand.Service.SetDetail(lang_p.LanguageEnglish,
		domainEnglish,
		"locale add an empty locale file to given directory",
		``,
		"",
		"",
		nil)
}
