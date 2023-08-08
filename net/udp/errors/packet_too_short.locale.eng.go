//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl memar/protocol.Detail
func (d *errPacketTooShort) Domain() string  { return domainEnglish }
func (d *errPacketTooShort) Summary() string { return "Packet Too Short" }
func (d *errPacketTooShort) Overview() string {
	return "UDP packet is empty or too short than standard header. It must include at least 20Byte header"
}
func (d *errPacketTooShort) UserNote() string { return }
func (d *errPacketTooShort) DevNote() string  { return }
func (d *errPacketTooShort) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errPacketTooShort) Name() string         { return "" }
func (d *errPacketTooShort) Abbreviation() string { return "" }
func (d *errPacketTooShort) Aliases() []string    { return []string{} }
