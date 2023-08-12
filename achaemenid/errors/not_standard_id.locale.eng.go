//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errNotStandardID) Domain() string  { return domainEnglish }
func (d *errNotStandardID) Summary() string { return "Not Standard ID" }
func (d *errNotStandardID) Overview() string {
	return "You set non standard ID for error||service||data-structure||..., It can cause some bad situation in your platform"
}
func (d *errNotStandardID) UserNote() string { return "" }
func (d *errNotStandardID) DevNote() string  { return "" }
func (d *errNotStandardID) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNotStandardID) Name() string         { return "" }
func (d *errNotStandardID) Abbreviation() string { return "" }
func (d *errNotStandardID) Aliases() []string    { return []string{} }
