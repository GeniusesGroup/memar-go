//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl memar/protocol.Detail
func (d *errNotFound) Domain() string  { return domainPersian }
func (d *errNotFound) Summary() string { return "یافت نشد" }
func (d *errNotFound) Overview() string {
	return "خطایی رخ داده است ولی جزییات آن خطا برای نمایش به شما ثبت نشده است"
}
func (d *errNotFound) UserNote() string {
	return "اشکال بوجود آماده بدلیل نقض عملیات توسعه ما می باشد. خواهشمندیم با پشتیبانی نرم افزار برای رفع این مشکل در تماس باشید"
}
func (d *errNotFound) DevNote() string {
	return "خطای رخ داده شده را با استفاده از URN آن پیدا کرده و با استفاده از متد ذخیره آن را برای هر نوع استفاده رابط کاربری آماده کنید"
}
func (d *errNotFound) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNotFound) Name() string         { return "" }
func (d *errNotFound) Abbreviation() string { return "" }
func (d *errNotFound) Aliases() []string    { return []string{} }
