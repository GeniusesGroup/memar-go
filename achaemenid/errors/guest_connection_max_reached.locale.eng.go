//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errGuestConnectionMaxReached) Domain() string  { return domainEnglish }
func (d *errGuestConnectionMaxReached) Summary() string { return "Guest Connection Max Reached" }
func (d *errGuestConnectionMaxReached) Overview() string {
	return "Server not have enough resource to make new guest connection, try few minutes later or try other server"
}
func (d *errGuestConnectionMaxReached) UserNote() string { return "" }
func (d *errGuestConnectionMaxReached) DevNote() string  { return "" }
func (d *errGuestConnectionMaxReached) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errGuestConnectionMaxReached) Name() string         { return "" }
func (d *errGuestConnectionMaxReached) Abbreviation() string { return "" }
func (d *errGuestConnectionMaxReached) Aliases() []string    { return []string{} }
