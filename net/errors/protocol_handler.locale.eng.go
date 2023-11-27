//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl errProtocolHandler memar/protocol.Detail
func (dt *errProtocolHandler) Domain() string  { return domainEnglish }
func (dt *errProtocolHandler) Summary() string { return "Protocol Handler" }
func (dt *errProtocolHandler) Overview() string {
	return "Protocol handler not exist to complete the request"
}
func (dt *errProtocolHandler) UserNote() string {
	return ""
}
func (dt *errProtocolHandler) DevNote() string {
	return ""
}
func (dt *errProtocolHandler) TAGS() []string { return []string{} }

//memar:impl errProtocolHandler memar/protocol.Quiddity
func (dt *errProtocolHandler) Name() string         { return "" }
func (dt *errProtocolHandler) Abbreviation() string { return "" }
func (dt *errProtocolHandler) Aliases() []string    { return []string{} }
