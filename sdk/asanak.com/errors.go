/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	er "../../error"
	lang "../../language"
)

const errorEnglishDomain = "Asanak.com"
const errorPersianDomain = "آسانک"

// Errors
var (
	ErrNotRecivied = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Not Recivied",
		"Sent SMS not recivied to destination").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم دریافت پیام",
			"پیام ارسال شده توسط مقصد یا مقصدهای مورد نظر دریافت نشد").Save()

	// ErrSMSProviderError = er.New("SMSProviderError", "Our SMS provider API can't proccess send OTP message")
)
