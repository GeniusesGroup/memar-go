/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"../giti"
)

const errorEnglishDomain = "Error"
const errorPersianDomain = "خطا"

// package errors
var (
	ErrNotFound = New("urn:giti:error.giti:error:not-found").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Not Found",
		"An error occurred but it is not registered yet to show more detail to you!",
		"Sorry it's us not your fault! Contact administrator of platform!",
		"Find error by its URN and save it for further use by any UserInterfaces").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "یافت نشد",
			"خطایی رخ داده است ولی جزییات آن خطا برای نمایش به شما ثبت نشده است",
			"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی پلتفرم برای رفع این مشکل در تماس باشید",
			"خطای رخ داده شده را با استفاده از URN آن پیدا کرده و با استفاده از متد ذخیره آن را برای هر نوع استفاده رابط کاربری آماده کنید").Save()

	ErrIsEmpty = New("urn:giti:error.giti:error:is-empty").SetDetail(giti.LanguageEnglish, errorEnglishDomain, "Is Empty",
		"Given Error is not exist",
		"Sorry it's us not your fault! Contact administrator of platform!",
		"Trace error by enable panic recovery to find nil error detection problem").
		SetDetail(giti.LanguagePersian, errorPersianDomain, "وجود ندارد",
			"خطایی با آدرس حافظه داده شده یافت نشد",
			"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی پلتفرم برای رفع این مشکل در تماس باشید",
			"ارور بوجود آمده را با استفاده از فعال سازی قابلیت کامپایلر زبان برنامه نویسی خود، منشا خطا ناموجود را پیدا کنید").Save()
)
