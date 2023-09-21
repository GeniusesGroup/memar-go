//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/detail"
	"memar/protocol"
)

const domainPersian = "HTTP"

func init() {
	ErrNoConnection.SetDetail(protocol.LanguagePersian, domainPersian,
		"عدم وجود ارتباط",
		"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد",
		"",
		"",
		nil)
		
		ErrServiceNotAcceptHTTP.SetDetail(detail.New(protocol.LanguagePersian, domainPersian).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("پروتکل HTTP پشتیبانی نمی شود").
			SetOverview("درخواست برای سرویس با شماره داده شده بدلیل عدم پشتیبانی پروتکل مدنظر قابلیت انجام روی سرور فعلی را ندارد").
			SetUserNote("سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید").
			SetDevNote("پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید").
			SetTAGS([]string{})
		)
}
