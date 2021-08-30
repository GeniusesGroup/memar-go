/* For license and copyright information please see LEGAL file in repository */

package protocol

type Language interface {
	ID() LanguageID
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
