//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl errGuestMaxReached memar/protocol.Detail
func (dt *errGuestMaxReached) Domain() string  { return domainEnglish }
func (dt *errGuestMaxReached) Summary() string { return "Guest Max Reached" }
func (dt *errGuestMaxReached) Overview() string {
	return "Server not have enough resource to make new guest connection, try few minutes later or try other server"
}
func (dt *errGuestMaxReached) UserNote() string {
	return ""
}
func (dt *errGuestMaxReached) DevNote() string {
	return ""
}
func (dt *errGuestMaxReached) TAGS() []string { return []string{} }

//memar:impl errGuestMaxReached memar/protocol.Quiddity
func (dt *errGuestMaxReached) Name() string         { return "" }
func (dt *errGuestMaxReached) Abbreviation() string { return "" }
func (dt *errGuestMaxReached) Aliases() []string    { return []string{} }
