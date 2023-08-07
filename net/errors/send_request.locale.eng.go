//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl errSendRequest memar/protocol.Detail
func (dt *errSendRequest) Domain() string  { return domainEnglish }
func (dt *errSendRequest) Summary() string { return "Send Request" }
func (dt *errSendRequest) Overview() string {
	return "Send request encounter problem due to temporary or long term problem!"
}
func (dt *errSendRequest) UserNote() string {
	return ""
}
func (dt *errSendRequest) DevNote() string {
	return ""
}
func (dt *errSendRequest) TAGS() []string { return []string{} }

//memar:impl errSendRequest memar/protocol.Quiddity
func (dt *errSendRequest) Name() string         { return "" }
func (dt *errSendRequest) Abbreviation() string { return "" }
func (dt *errSendRequest) Aliases() []string    { return []string{} }
