//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl errNoConnection memar/protocol.Detail
func (dt *errNoConnection) Domain() string  { return domainEnglish }
func (dt *errNoConnection) Summary() string { return "No Connection" }
func (dt *errNoConnection) Overview() string {
	return "No connection exist to complete request due to temporary or long term problem"
}
func (dt *errNoConnection) UserNote() string {
	return ""
}
func (dt *errNoConnection) DevNote() string {
	return ""
}
func (dt *errNoConnection) TAGS() []string { return []string{} }

//memar:impl errNoConnection memar/protocol.Quiddity
func (dt *errNoConnection) Name() string         { return "" }
func (dt *errNoConnection) Abbreviation() string { return "" }
func (dt *errNoConnection) Aliases() []string    { return []string{} }
