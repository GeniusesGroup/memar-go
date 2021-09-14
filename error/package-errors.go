/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"../protocol"
)

const errorEnglishDomain = "Error"
const errorPersianDomain = "خطا"

// package errors
var (
	ErrNotFound = New("urn:giti:error.giti:error:not-found").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Not Found",
		"An error occurred but it is not registered yet to show more detail to you!",
		"Sorry it's us not your fault! Contact administrator of platform!",
		"Find error by its URN and save it for further use by any UserInterfaces").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "یافت نشد",
			"خطایی رخ داده است ولی جزییات آن خطا برای نمایش به شما ثبت نشده است",
			"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی پلتفرم برای رفع این مشکل در تماس باشید",
			"خطای رخ داده شده را با استفاده از URN آن پیدا کرده و با استفاده از متد ذخیره آن را برای هر نوع استفاده رابط کاربری آماده کنید").Save()

	ErrIsEmpty = New("urn:giti:error.giti:error:is-empty").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Is Empty",
		"Given Error is not exist",
		"Sorry it's us not your fault! Contact administrator of platform!",
		"Trace error by enable panic recovery to find nil error detection problem").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "وجود ندارد",
			"خطایی با آدرس حافظه داده شده یافت نشد",
			"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی پلتفرم برای رفع این مشکل در تماس باشید",
			"ارور بوجود آمده را با استفاده از فعال سازی قابلیت کامپایلر زبان برنامه نویسی خود، منشا خطا ناموجود را پیدا کنید").Save()

	ErrSDKNotFound = New("urn:giti:error.giti:error:sdk-not-found").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "SDK Not Found",
		"Requested software errors SDK in desire human language not found",
		"Contact administrator of software and ask add needed data",
		"You can use online translator and desire human language to errors details").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "بسته توسعه یافت نشد",
			"بسته توسعه نرم افزار برای خطاهای این نرم افزار برای زبان انسانی مورد نظر شما یافت نشد",
			"برای بدست آوردن داده های مورد نظر خود با پشتیبان نرم افزار تماس حاصل فرمایید",
			"شما می توانید با استفاده از سرویس های ترجمه آنلاین اطلاعات مورد نیاز ارورها را در زبان انسانی مورد نظر اضافه کنید").Save()
)
