//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainPersian = "سرویس"

func init() {
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainPersian,
		"یافت نشد",
		"سرویس درخواست شده برای پردازش یافت نشد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"",
		nil)

	ErrServiceNotAcceptSRPC.SetDetail(protocol.LanguagePersian, domainPersian,
		"سرویس پروتکل sRPC را پشتیبانی نمی کند",
		"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید",
		nil)

	ErrServiceNotAcceptSRPCDirect.SetDetail(protocol.LanguagePersian, domainPersian,
		"پرسش مستقیم پشتیبانی نمی شود",
		"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، در صورت تمایل به پشتیبانی، وقتی برای پیاده سازی اختصاص دهید",
		nil)

	ErrServiceNotAcceptHTTP.SetDetail(protocol.LanguagePersian, domainPersian,
		"پروتکل HTTP پشتیبانی نمی شود",
		"درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید",
		nil)
}
