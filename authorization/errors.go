/* For license and copyright information please see LEGAL file in repository */

package authorization

import (
	errorr "../error"
	lang "../language"
)

// Errors
var (
	ErrAuthorizationServiceNotAllow = errorr.New().
					SetDetail(lang.EnglishLanguage, "Authorization - Service Not Allow",
			"Request service is not in allow list of connection").
		SetDetail(lang.PersianLanguage, "سطح دسترسی - عدم اجازه به سرویس",
			"درخواست به سرویس مورد نظر در لیست دسترسی های مجاز در ارتباط فعلی نمی باشد").Save()

	ErrAuthorizationServiceDenied = errorr.New().
					SetDetail(lang.EnglishLanguage, "Authorization - Service Denied",
			"Request service is in deny list of connection").
		SetDetail(lang.PersianLanguage, "سطح دسترسی - سرویس منع شده",
			"درخواست به سرویس مورد نظر در لیست سرویس های منع شده ارتباط فعلی می باشد").Save()

	ErrAuthorizationNotAllowSociety = errorr.New().
					SetDetail(lang.EnglishLanguage, "Authorization - Not Allow Society",
			"Request send by society that is not in allow list of connection").
		SetDetail(lang.PersianLanguage, "سطح دسترسی - عدم اجازه از جامعه درخواستی",
			"درخواست از جامعه ای ارسال می شود که در لیست جامعه های مجاز در ارتباط فعلی نمی باشد").Save()

	ErrAuthorizationDeniedSociety = errorr.New().
					SetDetail(lang.EnglishLanguage, "Authorization - Denied Society",
			"Request send by society that is in deny list of connection").
		SetDetail(lang.PersianLanguage, "سطح دسترسی - درخواست از جامعه منع شده",
			"درخواست از جامعه ای ارسال می شود که در لیست جامعه های منع شده ارتباط فعلی می باشد").Save()

	ErrAuthorizationNotAllowRouter = errorr.New().
					SetDetail(lang.EnglishLanguage, "Authorization - Not Allow Router",
			"Request send by router that is not in allow list of connection").
		SetDetail(lang.PersianLanguage, "سطح دسترسی - عدم اجازه از روتر درخواستی",
			"درخواست از روتر شبکه ای ارسال می شود که در لیست روتر های مجاز ارتباط فعلی نمی باشد").Save()

	ErrAuthorizationDeniedRouter = errorr.New().
					SetDetail(lang.EnglishLanguage, "Authorization - Denied Router",
			"Request send by router that is in deny list of connection").
		SetDetail(lang.PersianLanguage, "سطح دسترسی - درخواست از روتر منع شده",
			"درخواست از روتر شبکه ای ارسال می شود که در لیست روتر های منع شده ارتباط فعلی می باشد").Save()
)
