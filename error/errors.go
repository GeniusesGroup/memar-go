/* For license and copyright information please see LEGAL file in repository */

package error

import (
	lang "../language"
)

const errorEnglishDomain = "Error"
const errorPersianDomain = "خطا"

// package errors
var (
	ErrNotFound = New("urn:giti:error.libgo:error:not-found").
			SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Not Found", "An error occurred but it is not registered yet to show more detail to you!").
			SetDetail(lang.LanguagePersian, errorPersianDomain, "یافت نشد", "خطایی رخ داده است ولی جزییات آن خطا برای نمایش به شما ثبت نشده است").
			Save()

	ErrIsEmpty = New("urn:giti:error.libgo:error:is-empty").
			SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Is Empty", "Given Error is not exist").
			SetDetail(lang.LanguagePersian, errorPersianDomain, "وجود ندارد", "خطایی با آدرس حافظه داده شده یافت نشد").
			Save()
)
