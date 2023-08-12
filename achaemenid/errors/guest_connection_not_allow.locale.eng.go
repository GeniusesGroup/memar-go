//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errGuestConnectionNotAllow) Domain() string  { return domainEnglish }
func (d *errGuestConnectionNotAllow) Summary() string { return "Guest Connection Not Allow" }
func (d *errGuestConnectionNotAllow) Overview() string {
	return "Guest users don't allow to make new connection"
}
func (d *errGuestConnectionNotAllow) UserNote() string { return "" }
func (d *errGuestConnectionNotAllow) DevNote() string  { return "" }
func (d *errGuestConnectionNotAllow) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errGuestConnectionNotAllow) Name() string         { return "" }
func (d *errGuestConnectionNotAllow) Abbreviation() string { return "" }
func (d *errGuestConnectionNotAllow) Aliases() []string    { return []string{} }
