//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl memar/protocol.Detail
func (d *errNotExist) Domain() string   { return domainEnglish }
func (d *errNotExist) Summary() string  { return "Not Exist" }
func (d *errNotExist) Overview() string { return "Given Error is not exist" }
func (d *errNotExist) UserNote() string {
	return "Sorry it's us not your fault! Contact administrator of platform"
}
func (d *errNotExist) DevNote() string {
	return "Trace error by enable panic recovery to find nil error detection problem"
}
func (d *errNotExist) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNotExist) Name() string         { return "" }
func (d *errNotExist) Abbreviation() string { return "" }
func (d *errNotExist) Aliases() []string    { return []string{} }
