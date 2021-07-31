/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	er "../error"
	"../giti"
)

const errorEnglishDomain = "Chapar"
const errorPersianDomain = "چاپار"

// Declare Errors Details
var (
	ErrShortFrameLength = er.New("urn:giti:chapar.giti:error:short-frame-length").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Short Frame Length",
		"Chapar frame is too short(<12) than standard",
		"",
		"").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "اندازه کوتاه فریم",
			"ابعاد و اندازه فریم چاپار بررسی شده کمتر از 12 بایت می باشد که مجاز نمی باشد",
			"",
			"").Save()

	ErrLongFrameLength = er.New("urn:giti:chapar.giti:error:long-frame-length").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Long Frame Length",
		"Chapar frame is too long(>8192) than standard",
		"",
		"").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "اندازه بلند فریم",
			"ابعاد و اندازه فریم چاپار بررسی شده بیشتر از 8192 بایت می باشد که مجاز نمی باشد",
			"",
			"").Save()

	ErrMTU = er.New("urn:giti:chapar.giti:error:maximum-transmission-unit").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Maximum Transmission Unit - MTU",
		"Chapar frame isn't legal due to MTU is not respected by payload!",
		"",
		"").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "حداکثر طول قابل ارسال",
			"اندازه فریم چاپار بررسی شده بدلیل عدم رعایت طول در بار مفید فریم مجاز نمی باشد",
			"",
			"").Save()

	ErrPortNotExist = er.New("urn:giti:chapar.giti:error:port-not-exist").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Port Not Exist",
		"Chapar frame can't be handle due to frame want to switch to a port that not exist in network",
		"",
		"").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "پورت وجود ندارد",
			"فریم چاپار قابلیت مسیریابی ندارد و دلیل هم درخواست به مسیری که وجود ندارد می باشد",
			"",
			"").Save()

	ErrPathAlreadyUse = er.New("urn:giti:chapar.giti:error:path-already-use").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Path Already Use",
		"Path already use as main chapar connection path",
		"",
		"").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "مسیر در حال استفاده می باشد",
			"مسیری که برای مسیر اصلی ارتباط انتخاب شده است با مسیر فعلی یکسان می باشد",
			"",
			"").Save()

	ErrPathAlreadyExist = er.New("urn:giti:chapar.giti:error:path-already-exist").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Path Already Exist",
		"Path already exist in chapar connection alternative paths",
		"",
		"").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "مسیر موجود می باشد",
			"مسیری که برای اضافه کردن به مسیرهای جایگزین به ارتباط انتخاب شده است قبلا اضافه شده است",
			"",
			"").Save()
)
