//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"memar/detail"
	"memar/protocol"
)

const domainPersian = "چاپار"

func init() {
	ErrShortFrameLength.SetDetail(protocol.LanguagePersian, domainPersian, "اندازه کوتاه فریم",
		"ابعاد و اندازه فریم چاپار بررسی شده کمتر از 12 بایت می باشد که مجاز نمی باشد",
		"",
		"",
		nil)
	ErrLongFrameLength.SetDetail(protocol.LanguagePersian, domainPersian, "اندازه بلند فریم",
		"ابعاد و اندازه فریم چاپار بررسی شده بیشتر از 8192 بایت می باشد که مجاز نمی باشد",
		"",
		"",
		nil)
	ErrMTU.SetDetail(protocol.LanguagePersian, domainPersian, "حداکثر طول قابل ارسال",
		"اندازه فریم چاپار بررسی شده بدلیل عدم رعایت طول در بار مفید فریم مجاز نمی باشد",
		"",
		"",
		nil)
	ErrPortNotExist.SetDetail(protocol.LanguagePersian, domainPersian, "پورت وجود ندارد",
		"فریم چاپار قابلیت مسیریابی ندارد و دلیل هم درخواست به مسیری که وجود ندارد می باشد",
		"",
		"",
		nil)
	ErrPathAlreadyUse.SetDetail(protocol.LanguagePersian, domainPersian, "مسیر در حال استفاده می باشد",
		"مسیری که برای مسیر اصلی ارتباط انتخاب شده است با مسیر فعلی یکسان می باشد",
		"",
		"",
		nil)
	ErrPathAlreadyExist.SetDetail(protocol.LanguageEnglish, domainEnglish, "Path Already Exist",
		"Path already exist in chapar connection alternative paths",
		"",
		"",
		nil)
	ErrPathAlreadyExist.SetDetail(protocol.LanguagePersian, domainPersian, "مسیر موجود می باشد",
		"مسیری که برای اضافه کردن به مسیرهای جایگزین به ارتباط انتخاب شده است قبلا اضافه شده است",
		"",
		"",
		nil)
}
