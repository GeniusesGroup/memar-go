//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errors

//memar:impl memar/protocol.Detail
func (d *errPacketWrongLength) Domain() string  { return domainEnglish }
func (d *errPacketWrongLength) Summary() string { return "Packet Wrong Length" }
func (d *errPacketWrongLength) Overview() string {
	return "Data offset set in UDP packet header is not set correctly"
}
func (d *errPacketWrongLength) UserNote() string { return }
func (d *errPacketWrongLength) DevNote() string  { return }
func (d *errPacketWrongLength) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errPacketWrongLength) Name() string         { return "" }
func (d *errPacketWrongLength) Abbreviation() string { return "" }
func (d *errPacketWrongLength) Aliases() []string    { return []string{} }
