/* For license and copyright information please see LEGAL file in repository */

package service

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "Service"
const domainPersian = "سرویس"

// Declare package errors
var (
	ErrNotFound                   er.Error
	ErrServiceNotAcceptSRPC       er.Error
	ErrServiceNotAcceptSRPCDirect er.Error
	ErrServiceNotAcceptHTTP       er.Error
)

func init() {
	ErrNotFound.Init("domain/service.protocol.error; name=not-found")
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested service by given identifier not found in this application",
		"",
		"",
		nil)
	ErrNotFound.RegisterError()

	ErrServiceNotAcceptSRPC.Init("domain/service.protocol.error; name=service-not-accept-srpc")
	ErrServiceNotAcceptSRPC.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept sRPC",
		"Requested service by given ID not accept sRPC protocol in this server",
		"Try other server or contact support of the software",
		"It is so easy to implement sRPC handler for a service! Take a time and do it!",
		nil)
	ErrServiceNotAcceptSRPC.SetDetail(protocol.LanguagePersian, domainPersian,
		"سرویس پروتکل sRPC را پشتیبانی نمی کند",
		"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید",
		nil)
	ErrServiceNotAcceptSRPC.RegisterError()

	ErrServiceNotAcceptSRPCDirect.Init("domain/service.protocol.error; name=service-not-accept-direct-srpc")
	ErrServiceNotAcceptSRPCDirect.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept Direct sRPC",
		"Requested service by given ID not accept direct sRPC protocol in this server",
		"Try other server or contact support of the software",
		"",
		nil)
	ErrServiceNotAcceptSRPCDirect.SetDetail(protocol.LanguagePersian, domainPersian,
		"پرسش مستقیم پشتیبانی نمی شود",
		"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، در صورت تمایل به پشتیبانی وقتی برای پیاده سازی اختصاص دهید",
		nil)
	ErrServiceNotAcceptSRPCDirect.RegisterError()

	ErrServiceNotAcceptHTTP.Init("domain/service.protocol.error; name=service-not-accept-http")
	ErrServiceNotAcceptHTTP.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Not Accept HTTP",
		"Requested service by given ID not accept HTTP protocol in this server",
		"Try other server or contact support of the software",
		"It is so easy to implement HTTP handler for a service! Take a time and do it!",
		nil)
	ErrServiceNotAcceptHTTP.SetDetail(protocol.LanguagePersian, domainPersian,
		"پروتکل HTTP پشتیبانی نمی شود",
		"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید",
		nil)
	ErrServiceNotAcceptHTTP.RegisterError()
}
