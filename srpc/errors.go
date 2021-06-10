/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "sRPC"
const errorPersianDomain = "sRPC"

// Errors
var (
	ErrPacketTooShort = er.New("urn:giti:srpc.giti:error:packet-too-short").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Packet Too Short",
		"Received sRPC Packet size is smaller than expected and can't use",
		"",
		"").Save()

	ErrServiceNotFound = er.New("urn:giti:srpc.giti:error:service-not-found").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Service Not Found",
		"Requested service by given ID not found in this server",
		"",
		"").Save()

	ErrServiceNotAcceptSRPC = er.New("urn:giti:srpc.giti:error:service-not-accept-srpc").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Service Not Accept sRPC",
		"Requested service by given ID not accept sRPC protocol in this server",
		"Try other server or contact support of the platform",
		"It is so easy to implement sRPC handler for a service! Take a time and do it!").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "سرویس پروتکل sRPC را پشتیبانی نمی کند",
			"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
			"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
			"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید").Save()
)
