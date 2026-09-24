//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"memar/detail"
	"memar/protocol"
)

const domainPersian = "سرویس"

func init() {
	ErrNotFound.SetDetail(detail.New(lang_p.LanguagePersian, domainPersian).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("یافت نشد").
		SetOverview("سرویس درخواست شده برای پردازش یافت نشد").
		SetUserActionNote("سرور دیگر را امتحان کنید یا با پشتیبانی نرم افزار تماس بگیرید").
		SetDevActionNote("").
		SetTAGS([]string{})
	)

	ErrServiceNotAcceptSRPC.SetDetail(detail.New(lang_p.LanguagePersian, domainPersian).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("سرویس پروتکل sRPC را پشتیبانی نمی کند").
		SetOverview("درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد").
		SetUserActionNote("سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید").
		SetDevActionNote("پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید").
		SetTAGS([]string{})
	)

	ErrServiceNotAcceptSRPCDirect.SetDetail(detail.New(lang_p.LanguagePersian, domainPersian).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("پرسش مستقیم پشتیبانی نمی شود").
		SetOverview("درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد").
		SetUserActionNote("سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید").
		SetDevActionNote("پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، در صورت تمایل به پشتیبانی، وقتی برای پیاده سازی اختصاص دهید").
		SetTAGS([]string{})
	)

}
