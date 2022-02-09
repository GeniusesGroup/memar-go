/* For license and copyright information please see LEGAL file in repository */

package service

import (
	er "../error"
	"../mediatype"
	"../protocol"
)

const domainEnglish = "Service"
const domainPersian = "سرویس"

// Declare Errors Details
var (
	ErrNotFound = er.New(mediatype.New("domain/service.protocol.error; name=not-found").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested service by given identifier not found in this application",
		"",
		"",
		nil).
		Expired(0, nil))

	ErrServiceNotAcceptSRPC = er.New(mediatype.New("domain/service.protocol.error; name=service-not-accept-srpc").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept sRPC",
		"Requested service by given ID not accept sRPC protocol in this server",
		"Try other server or contact support of the software",
		"It is so easy to implement sRPC handler for a service! Take a time and do it!",
		nil).
		SetDetail(protocol.LanguagePersian, domainPersian,
			"سرویس پروتکل sRPC را پشتیبانی نمی کند",
			"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
			"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
			"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید",
			nil).
		Expired(0, nil))

	ErrServiceNotAcceptSRPCDirect = er.New(mediatype.New("domain/service.protocol.error; name=service-not-accept-direct-srpc").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept Direct sRPC",
		"Requested service by given ID not accept direct sRPC protocol in this server",
		"Try other server or contact support of the software",
		"",
		nil).
		SetDetail(protocol.LanguagePersian, domainPersian,
			"پرسش مستقیم پشتیبانی نمی شود",
			"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
			"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
			"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، در صورت تمایل به پشتیبانی وقتی برای پیاده سازی اختصاص دهید",
			nil).
		Expired(0, nil))

	ErrServiceNotAcceptHTTP = er.New(mediatype.New("domain/service.protocol.error; name=service-not-accept-http").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept HTTP",
		"Requested service by given ID not accept HTTP protocol in this server",
		"Try other server or contact support of the software",
		"It is so easy to implement HTTP handler for a service! Take a time and do it!",
		nil).
		SetDetail(protocol.LanguagePersian, domainPersian,
			"پروتکل HTTP پشتیبانی نمی شود",
			"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
			"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
			"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید",
			nil).
		Expired(0, nil))
)
