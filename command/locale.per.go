//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	"libgo/detail"
	"libgo/protocol"
)

const domainPersian = "فرمان"

func init() {
	ErrServiceNotAcceptCLI.SetDetail(detail.New(protocol.LanguagePersian, domainPersian).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("پروتکل CLI پشتیبانی نمی شود").
		SetOverview("درخواست برای سرویس مدنظر بدلیل عدم پشتیبانی پروتکل مورد نیاز قابلیت انجام روی سرور فعلی را ندارد").
		SetUserNote("سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید").
		SetDevNote("پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید").
		SetTAGS([]string{})
	)
}
