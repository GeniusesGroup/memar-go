//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package app

const (
	domainEnglish = "app"
)

func init() {
	RootCommand.Service.SetDetail(lang_p.LanguageEnglish,
		domainEnglish,
		"",
		``,
		"",
		"",
		nil)
}
