/* For license and copyright information please see LEGAL file in repository */

package error

import (
	lang "../language"
)

// package errors
var (
	ErrErrorNotFound = New().
		SetName(lang.EnglishLanguage, "Error Not Found").
		SetDescription(lang.EnglishLanguage, "Given ErrorID not exist").
		SetName(lang.PersianLanguage, "خطا یافت نشد").
		SetDescription(lang.PersianLanguage, "خطایی با کد خطای داده شده یافت نشد").Save()
)
