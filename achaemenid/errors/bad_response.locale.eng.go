//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errBadResponse) Domain() string  { return domainEnglish }
func (d *errBadResponse) Summary() string { return "Bad Response" }
func (d *errBadResponse) Overview() string {
	return "Response data from peer is not valid"
}
func (d *errBadResponse) UserNote() string { return "" }
func (d *errBadResponse) DevNote() string  { return "" }
func (d *errBadResponse) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errBadResponse) Name() string         { return "" }
func (d *errBadResponse) Abbreviation() string { return "" }
func (d *errBadResponse) Aliases() []string    { return []string{} }
