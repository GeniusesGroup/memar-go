//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl memar/protocol.Detail
func (d *errNotFound) Domain() string  { return domainEnglish }
func (d *errNotFound) Summary() string { return "Not Found" }
func (d *errNotFound) Overview() string {
	return "An error occurred but it is not registered yet to show more detail to you!"
}
func (d *errNotFound) UserNote() string {
	return "Sorry it's us not your fault! Contact administrator of platform!"
}
func (d *errNotFound) DevNote() string {
	return "Find error by its URN and save it for further use by any UserInterfaces"
}
func (d *errNotFound) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNotFound) Name() string         { return "" }
func (d *errNotFound) Abbreviation() string { return "" }
func (d *errNotFound) Aliases() []string    { return []string{} }
