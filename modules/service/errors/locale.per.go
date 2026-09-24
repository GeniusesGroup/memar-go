//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

const domainPersian = "فرمانده"

func init() {
	ErrSDKNotFound.SetDetail(lang_p.LanguagePersian, domainPersian, "بسته توسعه یافت نشد",
		"بسته توسعه نرم افزار برای خطاهای این نرم افزار برای زبان انسانی مورد نظر شما یافت نشد",
		"برای بدست آوردن داده های مورد نظر خود با پشتیبان نرم افزار تماس حاصل فرمایید",
		"شما می توانید با استفاده از سرویس های ترجمه آنلاین اطلاعات مورد نیاز ارورها را در زبان انسانی مورد نظر اضافه کنید")
}
