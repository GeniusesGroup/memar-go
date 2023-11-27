//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl errGuestNotAllow memar/protocol.Detail
func (dt *errGuestNotAllow) Domain() string  { return domainEnglish }
func (dt *errGuestNotAllow) Summary() string { return "Guest Not Allow" }
func (dt *errGuestNotAllow) Overview() string {
	return "Guest users don't allow to make new connection"
}
func (dt *errGuestNotAllow) UserNote() string {
	return ""
}
func (dt *errGuestNotAllow) DevNote() string {
	return ""
}
func (dt *errGuestNotAllow) TAGS() []string { return []string{} }

//memar:impl errGuestNotAllow memar/protocol.Quiddity
func (dt *errGuestNotAllow) Name() string         { return "" }
func (dt *errGuestNotAllow) Abbreviation() string { return "" }
func (dt *errGuestNotAllow) Aliases() []string    { return []string{} }
