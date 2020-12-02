/* For license and copyright information please see LEGAL file in repository */

package error

import (
	lang "../language"
)

const errorEnglishDomain = "Error"
const errorPersianDomain = "خطا"

// package errors
var (
	ErrErrorNotFound = New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Not Found",
		"Given ErrorID not exist or not registered yet to show more detail to you!").
		SetDetail(lang.PersianLanguage, errorPersianDomain, "یافت نشد",
			"خطایی با کد خطای داده شده یافت نشد یا هنوز ثبت نشده است که اطلاعات آن به شما داده شود").Save()

	ErrErrorIsEmpty = New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Is Empty",
		"Given Error is not exist").
		SetDetail(lang.PersianLanguage, errorPersianDomain, "وجود ندارد",
			"خطایی با آدرس حافظه داده شده یافت نشد").Save()
)
