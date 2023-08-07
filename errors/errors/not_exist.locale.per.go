//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl memar/protocol.Detail
func (d *errNotExist) Domain() string  { return domainPersian }
func (d *errNotExist) Summary() string { return "وجود ندارد" }
func (d *errNotExist) Overview() string {
	return "خطایی با آدرس حافظه داده شده یافت نشد"
}
func (d *errNotExist) UserNote() string {
	return "اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی نرم افزار برای رفع این مشکل در تماس باشید"
}
func (d *errNotExist) DevNote() string {
	return "خطای بوجود آمده را با استفاده از فعال سازی قابلیت کامپایلر زبان برنامه نویسی خود، منشا خطا ناموجود را پیدا کنید"
}
func (d *errNotExist) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNotExist) Name() string         { return "" }
func (d *errNotExist) Abbreviation() string { return "" }
func (d *errNotExist) Aliases() []string    { return []string{} }
