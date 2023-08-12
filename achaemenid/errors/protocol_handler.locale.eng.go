//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errProtocolHandler) Domain() string  { return domainEnglish }
func (d *errProtocolHandler) Summary() string { return "Protocol Handler" }
func (d *errProtocolHandler) Overview() string {
	return "Protocol handler not exist to complete the request"
}
func (d *errProtocolHandler) UserNote() string { return "" }
func (d *errProtocolHandler) DevNote() string  { return "" }
func (d *errProtocolHandler) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errProtocolHandler) Name() string         { return "" }
func (d *errProtocolHandler) Abbreviation() string { return "" }
func (d *errProtocolHandler) Aliases() []string    { return []string{} }
