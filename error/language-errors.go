/* For license and copyright information please see LEGAL file in repository */

package error

// DUE TO GOLANG IMPORT CYCLE WE NEED DO THIS IN THIS PACKAGE!!

import (
	lang "../language"
)

const languageErrorEnglishDomain = "Language"
const languageErrorPersianDomain = "زبان"

// Declare Errors Details
var (
	ErrBadLanguage = New("urn:giti:language.libgo:error:bad-language").
		SetDetail(lang.LanguageEnglish, languageErrorEnglishDomain, "Bad Language", "Not supported or bad language selected!").
		SetDetail(lang.LanguagePersian, languageErrorPersianDomain, "زبان بد", "زبان انتخاب شده صحیح نمی باشد یا توسط نرم افزار پشتیبانی نمی شود").
		Save()
)
