//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	"libgo/protocol"
)

const domainPersian = "HTTP Services"

func init() {
	ErrNoConnection.SetDetail(protocol.LanguagePersian, domainPersian,
		"عدم وجود ارتباط",
		"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد",
		"",
		"",
		nil)
}
