//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errServiceNotAcceptCLI) Domain() string { return domainPersian }
func (d *errServiceNotAcceptCLI) Summary() string {
	return "پروتکل CLI پشتیبانی نمی شود"
}
func (d *errServiceNotAcceptCLI) Overview() string {
	return "درخواست برای سرویس مدنظر بدلیل عدم پشتیبانی پروتکل مورد نیاز قابلیت انجام روی سرور فعلی را ندارد"
}
func (d *errServiceNotAcceptCLI) UserNote() string {
	return "سرور دیگر را امتحان کنید یا با پشتیبانی پلتفرم تماس بگیرید"
}
func (d *errServiceNotAcceptCLI) DevNote() string {
	return "پیاده سازی این پروتکل برای پاسخ گویی به سرویس ها به شدت ساده است، وقتی برای پیاده سازی اختصاص دهید"
}
func (d *errServiceNotAcceptCLI) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errServiceNotAcceptCLI) Name() string         { return "" }
func (d *errServiceNotAcceptCLI) Abbreviation() string { return "" }
func (d *errServiceNotAcceptCLI) Aliases() []string    { return []string{} }
