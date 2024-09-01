//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errServiceCallByAlias) Domain() string  { return domainEnglish }
func (d *errServiceCallByAlias) Summary() string { return "Service Call By Alias" }
func (d *errServiceCallByAlias) Overview() string {
	return "We found given service name in aliases that don't support to serve by it."
}
func (d *errServiceCallByAlias) UserNote() string {
	return ""
}
func (d *errServiceCallByAlias) DevNote() string {
	return "Try to call the command by its name or its abbreviation"
}
func (d *errServiceCallByAlias) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errServiceCallByAlias) Name() string         { return "" }
func (d *errServiceCallByAlias) Abbreviation() string { return "" }
func (d *errServiceCallByAlias) Aliases() []string    { return []string{} }
