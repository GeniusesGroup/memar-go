//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errTimerBadStatus) Domain() string  { return domainEnglish }
func (d *errTimerBadStatus) Summary() string { return "Bad Status" }
func (d *errTimerBadStatus) Overview() string {
	return "Timer or Ticker is in a bad status and we don't process your request to Start, Stop or Reset it"
}
func (d *errTimerBadStatus) UserNote() string { return "" }
func (d *errTimerBadStatus) DevNote() string  { return "" }
func (d *errTimerBadStatus) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errTimerBadStatus) Name() string         { return "" }
func (d *errTimerBadStatus) Abbreviation() string { return "" }
func (d *errTimerBadStatus) Aliases() []string    { return []string{} }
