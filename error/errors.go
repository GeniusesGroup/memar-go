/* For license and copyright information please see LEGAL file in repository */

package error

import (
	lang "../language"
)

// package errors
var (
	ErrErrorNotFound = New().
				SetDetail(lang.EnglishLanguage, "Error - Not Found",
			"Given ErrorID not exist").
		SetDetail(lang.PersianLanguage, "خطا یافت نشد",
			"خطایی با کد خطای داده شده یافت نشد").Save()

	ErrErrorIsEmpty = New().
			SetDetail(lang.EnglishLanguage, "Error - Is Empty",
			"Given Error is not exist").
		SetDetail(lang.PersianLanguage, "خطایی وجود ندارد",
			"خطایی با آدرس حافظه داده شده یافت نشد").Save()
)
