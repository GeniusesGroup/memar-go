//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errSourceNotChangeable) Domain() string  { return domainEnglish }
func (d *errSourceNotChangeable) Summary() string { return "Source not Changeable" }
func (d *errSourceNotChangeable) Overview() string {
	return "Can't read from other source than source given in compression||decompression creation"
}
func (d *errSourceNotChangeable) UserNote() string { return "" }
func (d *errSourceNotChangeable) DevNote() string  { return "" }
func (d *errSourceNotChangeable) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errSourceNotChangeable) Name() string         { return "" }
func (d *errSourceNotChangeable) Abbreviation() string { return "" }
func (d *errSourceNotChangeable) Aliases() []string    { return []string{} }
