//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainPersian = "خطا"

func init() {
	ErrNotFound.SetDetail(protocol.LanguagePersian, domainPersian,
		"یافت نشد",
		"خطایی رخ داده است ولی جزییات آن خطا برای نمایش به شما ثبت نشده است",
		"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی نرم افزار برای رفع این مشکل در تماس باشید",
		"خطای رخ داده شده را با استفاده از URN آن پیدا کرده و با استفاده از متد ذخیره آن را برای هر نوع استفاده رابط کاربری آماده کنید",
		nil)

	ErrNotExist.SetDetail(protocol.LanguagePersian, domainPersian,
		"وجود ندارد",
		"خطایی با آدرس حافظه داده شده یافت نشد",
		"اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی نرم افزار برای رفع این مشکل در تماس باشید",
		"خطای بوجود آمده را با استفاده از فعال سازی قابلیت کامپایلر زبان برنامه نویسی خود، منشا خطا ناموجود را پیدا کنید",
		nil)
}
