//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/detail"
	"memar/protocol"
)

const domainPersian = "سرویس"

func init() {
	ErrNotFound.SetDetail(detail.New(protocol.LanguagePersian, domainPersian).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("یافت نشد").
		SetOverview("سرویس درخواست شده برای پردازش یافت نشد").
		SetUserNote("سرور دیگر را امتحان کنید یا با پشتیبانی نرم افزار تماس بگیرید").
		SetDevNote("").
		SetTAGS([]string{})
	)
}
