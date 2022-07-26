/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainEnglish = "Error"
const domainPersian = "خطا"

// package errors
var (
	ErrNotFound Error
	ErrNotExist Error
)

func init() {
	ErrNotFound.Init("domain/error.protocol.error; name=not-found")
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"An error occurred but it is not registered yet to show more detail to you!",
		"Sorry it's us not your fault! Contact administrator of platform!",
		"Find error by its URN and save it for further use by any UserInterfaces",
		nil)
	ErrNotFound.SetDetail(protocol.LanguagePersian, domainPersian,
		"یافت نشد",
		"خطایی رخ داده است ولی جزییات آن خطا برای نمایش به شما ثبت نشده است",
		"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی نرم افزار برای رفع این مشکل در تماس باشید",
		"خطای رخ داده شده را با استفاده از URN آن پیدا کرده و با استفاده از متد ذخیره آن را برای هر نوع استفاده رابط کاربری آماده کنید",
		nil)

	ErrNotExist.Init("domain/error.protocol.error; name=not-exist")
	ErrNotExist.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Exist",
		"Given Error is not exist",
		"Sorry it's us not your fault! Contact administrator of platform",
		"Trace error by enable panic recovery to find nil error detection problem",
		nil)
	ErrNotExist.SetDetail(protocol.LanguagePersian, domainPersian,
		"وجود ندارد",
		"خطایی با آدرس حافظه داده شده یافت نشد",
		"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی نرم افزار برای رفع این مشکل در تماس باشید",
		"خطای بوجود آمده را با استفاده از فعال سازی قابلیت کامپایلر زبان برنامه نویسی خود، منشا خطا ناموجود را پیدا کنید",
		nil)
}
