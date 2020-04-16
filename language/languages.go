/* For license and copyright information please see LEGAL file in repository */

package lang

const (
	// EnglishLanguage language code
	EnglishLanguage uint32 = iota
	// PersianLanguage language code
	PersianLanguage
)

// LanguageDetail use to store other detail for a language
type LanguageDetail struct {
	Code        uint32
	Description []string
	ScriptsID   []uint32 //
}

// LanguagesDetails store all available language in
var LanguagesDetails = map[uint32]LanguageDetail{
	EnglishLanguage: LanguageDetail{
		Code:        EnglishLanguage,
		Description: []string{"English", "انگلیسی"},
	},
	PersianLanguage: LanguageDetail{
		Code:        PersianLanguage,
		Description: []string{"Persian", "پارسی"},
	},
}
