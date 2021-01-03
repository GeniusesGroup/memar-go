/* For license and copyright information please see LEGAL file in repository */

package authorization

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "Authorization"
const errorPersianDomain = "سطح دسترسی"

// Errors
var (
	/*
		User
	*/
	ErrUserNotAllow = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "User Not Allow",
		"Request service is not allow by user of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه به کاربر",
			"درخواست به سرویس مورد نظر توسط کاربر ارتباط فعلی مقدور نمی باشد").Save()

	ErrUserNotOwnRecord = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "User Not Own Record",
		"Request record is not own by user of the connection and by service rule can't access to it by other users").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم مالکیت داده",
			"درخواست به داده مورد نظر با توجه به تفاوت مالک اصلی و قوانین سرویس فراخوانی شده توسط کاربر ارتباط فعلی مقدور نمی باشد").Save()

	/*
		Society
	*/
	ErrNotAllowSociety = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Not Allow Society",
		"Request send by society that is not in allow list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه از جامعه درخواستی",
			"درخواست از جامعه ای ارسال می شود که در لیست جامعه های مجاز در ارتباط فعلی نمی باشد").Save()

	ErrDeniedSociety = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Denied Society",
		"Request send by society that is in deny list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "درخواست از جامعه منع شده",
			"درخواست از جامعه ای ارسال می شود که در لیست جامعه های منع شده ارتباط فعلی می باشد").Save()

	/*
		Router
	*/
	ErrNotAllowRouter = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Not Allow Router",
		"Request send by router that is not in allow list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه از روتر درخواستی",
			"درخواست از روتر شبکه ای ارسال می شود که در لیست روتر های مجاز ارتباط فعلی نمی باشد").Save()

	ErrDeniedRouter = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Denied Router",
		"Request send by router that is in deny list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "درخواست از روتر منع شده",
			"درخواست از روتر شبکه ای ارسال می شود که در لیست منع شده های ارتباط فعلی می باشد").Save()

	/*
		Day
	*/
	ErrDayNotAllow = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Day Not Allow",
		"Request send in the day of week that is not in allow list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه در روز درخواست",
			"درخواست در روزی از هفته ارسال شده است که در لیست مجاز در ارتباط فعلی نمی باشد").Save()

	ErrDayDenied = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Day Denied",
		"Request send in the day of week that is in deny list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه در روز درخواست",
			"درخواست در روزی از هفته ارسال شده است که در لیست غیر مجاز در ارتباط فعلی می باشد").Save()

	/*
		Hour
	*/
	ErrHourNotAllow = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Hour Not Allow",
		"Request send in the hour of day that is not in allow list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه ساعت درخواست",
			"درخواست در ساعتی از روز ارسال شده است که در لیست مجاز در ارتباط فعلی نمی باشد").Save()

	ErrHourDenied = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Hour Denied",
		"Request send in the hour of day that is in deny list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه ساعت درخواست",
			"درخواست در ساعتی از روز ارسال شده است که در لیست غیر مجاز در ارتباط فعلی می باشد").Save()

	/*
		Service
	*/
	ErrServiceNotAllow = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Service Not Allow",
		"Request service is not in allow list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه به سرویس",
			"درخواست به سرویس مورد نظر در لیست دسترسی های مجاز در ارتباط فعلی نمی باشد").Save()

	ErrServiceDenied = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Service Denied",
		"Request service is in deny list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "سرویس منع شده",
			"درخواست به سرویس مورد نظر در لیست سرویس های منع شده ارتباط فعلی می باشد").Save()

	/*
		CRUD
	*/
	ErrCrudNotAllow = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "CRUD Not Allow",
		"Request type is not in allow list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه نوع درخواست",
			"نوع درخواست در لیست مجاز در ارتباط فعلی نمی باشد").Save()

	ErrCRUDDenied = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "CRUD Denied",
		"Request type is deny list of the connection").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه نوع درخواست",
			"نوع درخواست در لیست غیر مجاز در ارتباط فعلی می باشد").Save()

	/*
		Delegate
	*/
	ErrNotAllowToDelegate = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Not Allow To Delegate",
		"Platforms rules not permit active type of user to register the delegate connection with given details").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه به وکالت دادن",
			"قوانین پلتفرم به نوع کاربر فعلی اجازه نمی دهد اتصال وکالتی مورد نظر را ایجاد نماید").Save()

	ErrNotAllowToNotDelegate = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Not Allow To Not Delegate",
		"Platforms rules not permit active type of user (usually Org type) register not delegate connection.").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم اجازه به وکالت ندادن",
			"قوانین پلتفرم به نوع کاربر فعلی (معمولا سازمان) اجازه نمی دهد اتصال غیر وکالتی ایجاد نماید").Save()
)
