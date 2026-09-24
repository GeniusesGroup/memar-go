//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package chapar_errs

const domainPersian = "چاپار"

func init() {
	ErrShortFrameLength.SetDetail(lang_p.LanguagePersian, domainPersian, "اندازه کوتاه فریم",
		"ابعاد و اندازه فریم چاپار بررسی شده کمتر از 12 بایت می باشد که مجاز نمی باشد",
		"",
		"",
		nil)
	ErrLongFrameLength.SetDetail(lang_p.LanguagePersian, domainPersian, "اندازه بلند فریم",
		"ابعاد و اندازه فریم چاپار بررسی شده بیشتر از 8192 بایت می باشد که مجاز نمی باشد",
		"",
		"",
		nil)
	ErrMTU.SetDetail(lang_p.LanguagePersian, domainPersian, "حداکثر طول قابل ارسال",
		"اندازه فریم چاپار بررسی شده بدلیل عدم رعایت طول در بار مفید فریم مجاز نمی باشد",
		"",
		"",
		nil)
	ErrPortNotExist.SetDetail(lang_p.LanguagePersian, domainPersian, "پورت وجود ندارد",
		"فریم چاپار قابلیت مسیریابی ندارد و دلیل هم درخواست به مسیری که وجود ندارد می باشد",
		"",
		"",
		nil)
	ErrPathAlreadyUse.SetDetail(lang_p.LanguagePersian, domainPersian, "مسیر در حال استفاده می باشد",
		"مسیری که برای مسیر اصلی ارتباط انتخاب شده است با مسیر فعلی یکسان می باشد",
		"",
		"",
		nil)
	ErrPathAlreadyExist.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Path Already Exist",
		"Path already exist in chapar connection alternative paths",
		"",
		"",
		nil)
	ErrPathAlreadyExist.SetDetail(lang_p.LanguagePersian, domainPersian, "مسیر موجود می باشد",
		"مسیری که برای اضافه کردن به مسیرهای جایگزین به ارتباط انتخاب شده است قبلا اضافه شده است",
		"",
		"",
		nil)
}
