//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl errNoConnection memar/protocol.Detail
func (dt *errNoConnection) Domain() string  { return domainPersian }
func (dt *errNoConnection) Summary() string { return "ارتباط قطع" }
func (dt *errNoConnection) Overview() string {
	return "ارتباطی جهت انجام رخواست مورد نظر بدلیل وجود مشکل موقت یا دایم وجود ندارد"
}
func (dt *errNoConnection) UserNote() string {
	return ""
}
func (dt *errNoConnection) DevNote() string {
	return ""
}
func (dt *errNoConnection) TAGS() []string { return []string{} }

//memar:impl errNoConnection memar/protocol.Quiddity
func (dt *errNoConnection) Name() string         { return "" }
func (dt *errNoConnection) Abbreviation() string { return "" }
func (dt *errNoConnection) Aliases() []string    { return []string{} }
