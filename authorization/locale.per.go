//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package authorization

import (
	"libgo/protocol"
)

const domainPersian = "سطح دسترسی"

func init() {
	ErrUserNotAllow.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه به کاربر",
		"درخواست به سرویس مورد نظر توسط کاربر ارتباط فعلی مقدور نمی باشد",
		"",
		"",
		nil)
	ErrUserNotOwnRecord.SetDetail(protocol.LanguagePersian, domainPersian, "عدم مالکیت داده",
		"درخواست به داده مورد نظر با توجه به تفاوت مالک اصلی و قوانین سرویس فراخوانی شده توسط کاربر ارتباط فعلی مقدور نمی باشد",
		"",
		"",
		nil)

	ErrNotAllowSociety.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه از جامعه درخواستی",
		"درخواست از جامعه ای ارسال می شود که در لیست جامعه های مجاز در ارتباط فعلی نمی باشد",
		"",
		"",
		nil)
	ErrDeniedSociety.SetDetail(protocol.LanguagePersian, domainPersian, "درخواست از جامعه منع شده",
		"درخواست از جامعه ای ارسال می شود که در لیست جامعه های منع شده ارتباط فعلی می باشد",
		"",
		"",
		nil)

	ErrNotAllowRouter.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه از روتر درخواستی",
		"درخواست از روتر شبکه ای ارسال می شود که در لیست روتر های مجاز ارتباط فعلی نمی باشد",
		"",
		"",
		nil)
	ErrDeniedRouter.SetDetail(protocol.LanguagePersian, domainPersian, "درخواست از روتر منع شده",
		"درخواست از روتر شبکه ای ارسال می شود که در لیست منع شده های ارتباط فعلی می باشد",
		"",
		"",
		nil)

	ErrDayNotAllow.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه در روز درخواست",
		"درخواست در روزی از هفته ارسال شده است که در لیست مجاز در ارتباط فعلی نمی باشد",
		"",
		"",
		nil)
	ErrDayDenied.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه در روز درخواست",
		"درخواست در روزی از هفته ارسال شده است که در لیست غیر مجاز در ارتباط فعلی می باشد",
		"",
		"",
		nil)

	ErrHourNotAllow.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه ساعت درخواست",
		"درخواست در ساعتی از روز ارسال شده است که در لیست مجاز در ارتباط فعلی نمی باشد",
		"",
		"",
		nil)
	ErrHourDenied.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه ساعت درخواست",
		"درخواست در ساعتی از روز ارسال شده است که در لیست غیر مجاز در ارتباط فعلی می باشد",
		"",
		"",
		nil)

	ErrServiceNotAllow.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه به سرویس",
		"درخواست به سرویس مورد نظر در لیست دسترسی های مجاز در ارتباط فعلی نمی باشد",
		"",
		"",
		nil)
	ErrServiceDenied.SetDetail(protocol.LanguagePersian, domainPersian, "سرویس منع شده",
		"درخواست به سرویس مورد نظر در لیست سرویس های منع شده ارتباط فعلی می باشد",
		"",
		"",
		nil)

	ErrCrudNotAllow.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه نوع درخواست",
		"نوع درخواست در لیست مجاز در ارتباط فعلی نمی باشد",
		"",
		"",
		nil)
	ErrCRUDDenied.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه نوع درخواست",
		"نوع درخواست در لیست غیر مجاز در ارتباط فعلی می باشد",
		"",
		"",
		nil)

	ErrNotAllowToDelegate.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه به وکالت دادن",
		"قوانین پلتفرم به نوع کاربر فعلی اجازه نمی دهد اتصال وکالتی مورد نظر را ایجاد نماید",
		"",
		"",
		nil)
	ErrNotAllowToNotDelegate.SetDetail(protocol.LanguagePersian, domainPersian, "عدم اجازه به وکالت ندادن",
		"قوانین پلتفرم به نوع کاربر فعلی (معمولا سازمان) اجازه نمی دهد اتصال غیر وکالتی ایجاد نماید",
		"",
		"",
		nil)
}
