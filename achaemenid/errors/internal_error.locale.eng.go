//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errInternalError) Domain() string  { return domainEnglish }
func (d *errInternalError) Summary() string { return "Internal Error" }
func (d *errInternalError) Overview() string {
	return "Peer encounter problem due to temporary or long term problem!"
}
func (d *errInternalError) UserNote() string { return "" }
func (d *errInternalError) DevNote() string  { return "" }
func (d *errInternalError) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errInternalError) Name() string         { return "" }
func (d *errInternalError) Abbreviation() string { return "" }
func (d *errInternalError) Aliases() []string    { return []string{} }
