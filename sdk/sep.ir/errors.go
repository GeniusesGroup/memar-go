/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"../../achaemenid"
	er "../../error"
	lang "../../language"
	"../../price"
)

const errorEnglishDomain = "SEP.ir"
const errorPersianDomain = "پرداخت الکترونیک سامان"

// Errors
var (
	ErrBadPOSSettings = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Bad POS Settings",
		"Can't find 'sep.ir-pos.json' file in 'secret' folder in top of repository or has not enough information").Save()

	ErrBadTerminalID = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Bad Terminal ID",
		"Requested Terminal ID is not valid by platform settings").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - پایانه غیر مجاز",
		"عملیات بدلیل درخواست به پایانه غیر مجاز لغو شد").Save()

	ErrPOSAuthenticationError = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS - SEP Authentication Error",
		"POS server encounter bank authentication error, Try again later").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - اشکال در تایید هوییت",
		"عملیات بدلیل وجود اشکال در سرور تایید هوییت بانک لغو شد").Save()

	ErrPOSInternalError = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS - SEP Server Error",
		"POS server encounter internal error, Try again later").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - اشکال در سرور مقصد",
		"عملیات بدلیل وجود اشکال در سرور مقصد لغو شد").Save()

	// SEP.ir Platform errors

	ErrNoActionAfterReadCard = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "No Action After Read Card",
		"Transaction canceled due to no action received after read card by POS device").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - عدم عمل پس از کشیدن کارت ",
		"عملیات بدلیل عدم عمل توسط کاربر پس از کشیدن کارت بر روی دستگاه پایانه لغو شد").Save()

	ErrAmountMinimum = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Amount Minimum",
		"Transaction canceled due to below legal minimum amount sent").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - مبلغ کمتر از مجاز",
		"عملیات بدلیل درخواست با مبلغ کمتر از حد مجاز لغو شد").Save()

	ErrPOSNotReachable = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS Not Reachable",
		"Transaction canceled due to selected POS not reachable").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - عدم دسترسی به پایانه",
		"عملیات بدلیل عدم دسترسی به پایانه مدنظر انجام نشد، دوباره تلاش کنید").Save()

	ErrPOSNotValidData = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS Not Valid Data",
		"Transaction canceled due to not valid data received").Save()

	ErrPOSAnotherRequest = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS - Request is in progress",
		"Another request is in progress, Complete or Cancel it before send new request").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - درخواست جاری",
		"عملیات بدلیل وجود درخواست جاری لغو شد، درخواست جاری را کامل یا رد کنید و دوباره تلاش کنید").Save()

	ErrPOSCardPassword = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS - Wrong Password",
		"Transaction canceled due to user not enter right password of its card").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - رمز اشتباه",
		"عملیات بدلیل ورود رمز اشتباه توسط کاربر بر روی دستگاه پوز لغو شد").Save()

	ErrPOSNotIndicated = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS Not Indicated",
		"Transaction canceled due to not indicate POS order service error").Save()

	ErrPOSCanceledByUser = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS - Canceled By User",
		"Transaction canceled due to user canceled transaction in POS device").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - لغو عملیات توسط کاربر",
		"عملیات توسط کاربر بر روی دستگاه پوز لغو شد").Save()

	ErrPOSNotResponse = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "POS Not Response",
		"Transaction canceled due to POS not response in proper time").SetDetail(
		lang.LanguagePersian, errorPersianDomain, "پوز - عدم پاسخ در زمان مناسب",
		"عملیات بدلیل عدم پاسخ پایانه در زمان مناسب لغو شد").Save()
)

func getErrorByResCode(code int64) (err *er.Error) {
	switch code {
	case 0:
		return nil
	case 1:
		return ErrNoActionAfterReadCard
	case 2:
		return ErrAmountMinimum
	case 3:
		return ErrPOSNotReachable
	case 4:
		return ErrPOSNotValidData
	case 9:
		return price.ErrInvalidTerminalID
	case 30:
		return ErrPOSAnotherRequest
	case 51:
		return price.ErrNotEnoughBalance
	case 55:
		return ErrPOSCardPassword
	case 93, 96:
		return ErrPOSNotIndicated
	case 98:
		return ErrPOSCanceledByUser
	case 99:
		return ErrPOSNotResponse
	default:
		return achaemenid.ErrInternalError
	}
}
