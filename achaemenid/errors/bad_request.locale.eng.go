//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errBadRequest) Domain() string  { return domainEnglish }
func (d *errBadRequest) Summary() string { return "Bad Request" }
func (d *errBadRequest) Overview() string {
	return "Some given data in request must be invalid or peer not accept them"
}
func (d *errBadRequest) UserNote() string { return "" }
func (d *errBadRequest) DevNote() string  { return "" }
func (d *errBadRequest) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errBadRequest) Name() string         { return "" }
func (d *errBadRequest) Abbreviation() string { return "" }
func (d *errBadRequest) Aliases() []string    { return []string{} }
