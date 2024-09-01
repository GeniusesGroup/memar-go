/* For license and copyright information please see the LEGAL file in the code repository */

package lang_p

type Language interface {
	LanguageID() LanguageID
}

// LanguageID indicate
type LanguageID uint32

// Languages code
const (
	LanguageUnset LanguageID = iota
	LanguageMachine
	LanguageMathematics
	LanguageMusical
	LanguageSign // https://en.wikipedia.org/wiki/Sign_language

	LanguagePersian
	LanguageEnglish
	LanguageRussian
	LanguageArabic
)
