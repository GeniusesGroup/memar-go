//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package http

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
}
