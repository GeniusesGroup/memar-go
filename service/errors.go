/* For license and copyright information please see LEGAL file in repository */

package service

import (
	er "../error"
	"../protocol"
)

const errorEnglishDomain = "Service"
const errorPersianDomain = "سرویس"

// Declare Errors Details
var (
	ErrServiceNotAcceptHTTP = er.New("urn:giti:service.giti:error:service-not-accept-http").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Service Not Accept HTTP",
		"Requested service by given ID not accept HTTP protocol in this server",
		"Try other server or contact support of the platform",
		"It is so easy to implement HTTP handler for a service! Take a time and do it!").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "پروتکل HTTP پشتیبانی نمی شود",
			"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
			"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
			"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید").Save()

	ErrServiceNotAcceptSRPC = er.New("urn:giti:service.giti:error:service-not-accept-srpc").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Service Not Accept sRPC",
		"Requested service by given ID not accept sRPC protocol in this server",
		"Try other server or contact support of the platform",
		"It is so easy to implement sRPC handler for a service! Take a time and do it!").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "سرویس پروتکل sRPC را پشتیبانی نمی کند",
			"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
			"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
			"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید").Save()

	ErrServiceNotAcceptCLI = er.New("urn:giti:service.giti:error:service-not-accept-cli").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Service Not Accept CLI",
		"Requested service not accept CLI protocol in this server",
		"Try other server or contact support of the platform",
		"It is so easy to implement CLI handler for a service! Take a time and do it!").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "پروتکل CLI پشتیبانی نمی شود",
			"درخواست برای سرویس مدنظر بدلیل عدم پشتیبانی پروتکل مورد نیاز قابلیت انجام روی سرور فعلی را ندارد",
			"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
			"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید").Save()
)
