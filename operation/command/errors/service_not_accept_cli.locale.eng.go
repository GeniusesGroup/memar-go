//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errServiceNotAcceptCLI) Domain() string  { return domainEnglish }
func (d *errServiceNotAcceptCLI) Summary() string { return "Service Not Accept CLI" }
func (d *errServiceNotAcceptCLI) Overview() string {
	return "Requested service not accept CLI protocol in this server"
}
func (d *errServiceNotAcceptCLI) UserNote() string {
	return "Try other server or contact support of the software"
}
func (d *errServiceNotAcceptCLI) DevNote() string {
	return "It is so easy to implement CLI handler for a service! Take a time and do it!"
}
func (d *errServiceNotAcceptCLI) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errServiceNotAcceptCLI) Name() string         { return "" }
func (d *errServiceNotAcceptCLI) Abbreviation() string { return "" }
func (d *errServiceNotAcceptCLI) Aliases() []string    { return []string{} }
