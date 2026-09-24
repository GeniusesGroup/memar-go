//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package modules

const (
	domainEnglish = "memar"
)

func init() {
	RootCommand.Service.SetDetail(lang_p.LanguageEnglish,
		domainEnglish,
		"Memar is a tool to develop software in Go by the Geniuses.Group architecture framework",
		"",
		"",
		"",
		nil)
}
