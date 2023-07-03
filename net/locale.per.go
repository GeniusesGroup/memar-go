//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"libgo/protocol"
)

const domainPersian = "ارتباط"

func init() {
	ErrNoConnection.SetDetail(protocol.LanguagePersian, domainPersian, "ارتباط قطع",
		"ارتباطی جهت انجام رخواست مورد نظر بدلیل وجود مشکل موقت یا دایم وجود ندارد",
		"",
		"",
		nil)
}
