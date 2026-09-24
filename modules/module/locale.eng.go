//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package module

const (
	domainEnglish = "module"
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
