//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errNotFound) Domain() string  { return domainEnglish }
func (d *errNotFound) Summary() string { return "Not Found" }
func (d *errNotFound) Overview() string {
	return "Can't find requested compression||decompression algorithm"
}
func (d *errNotFound) UserNote() string { return "" }
func (d *errNotFound) DevNote() string  { return "" }
func (d *errNotFound) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNotFound) Name() string         { return "" }
func (d *errNotFound) Abbreviation() string { return "" }
func (d *errNotFound) Aliases() []string    { return []string{} }
