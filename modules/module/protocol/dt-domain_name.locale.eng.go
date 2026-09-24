//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package mp

//memar:impl DomainName memar/protocol.Detail
func (d *DomainName) Domain() string  { return DomainEnglish }
func (d *DomainName) Summary() string { return "Domain name" }
func (d *DomainName) Overview() string {
	return "module domain name e.g. user-name"
}
func (d *DomainName) UserActionNote() string {
	return ""
}
func (d *DomainName) DevActionNote() string {
	return ""
}
func (d *DomainName) TAGS() []string { return []string{} }

//memar:impl DomainName memar/protocol.Quiddity
func (d *DomainName) Name() string         { return "name" }
func (d *DomainName) Abbreviation() string { return "n" }
func (d *DomainName) Aliases() []string    { return []string{} }
