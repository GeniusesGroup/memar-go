//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl GuestMaxReached memar/protocol.Detail
func (dt *guestMaxReached) Domain() string  { return domainEnglish }
func (dt *guestMaxReached) Summary() string { return "Guest Max Reached" }
func (dt *guestMaxReached) Overview() string {
	return "Server not have enough resource to make new guest connection, try few minutes later or try other server"
}
func (dt *guestMaxReached) UserActionNote() string {
	return ""
}
func (dt *guestMaxReached) DevActionNote() string {
	return ""
}
func (dt *guestMaxReached) TAGS() []string { return []string{} }

//memar:impl GuestMaxReached memar/protocol.Quiddity
func (dt *guestMaxReached) Name() string         { return "" }
func (dt *guestMaxReached) Abbreviation() string { return "" }
func (dt *guestMaxReached) Aliases() []string    { return []string{} }
