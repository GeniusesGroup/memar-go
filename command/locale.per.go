//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	er "github.com/GeniusesGroup/libgo/error"
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainPersian = "فرمان"

func init() {
	ErrServiceNotAcceptCLI.SetDetail(protocol.LanguagePersian, domainPersian, "پروتکل CLI پشتیبانی نمی شود",
		"درخواست برای سرویس مدنظر بدلیل عدم پشتیبانی پروتکل مورد نیاز قابلیت انجام روی سرور فعلی را ندارد",
		"سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید",
		"پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید",
		nil)
}
