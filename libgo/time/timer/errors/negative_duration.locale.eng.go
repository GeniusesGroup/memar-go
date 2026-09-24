//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

//memar:impl memar/protocol.Detail
func (d *negativeDuration) Domain() string  { return domainEnglish }
func (d *negativeDuration) Summary() string { return "Negative Duration" }
func (d *negativeDuration) Overview() string {
	return "Timer or Ticker must have positive duration or interval."
}
func (d *negativeDuration) UserActionNote() string { return "" }
func (d *negativeDuration) DevActionNote() string  { return "" }
func (d *negativeDuration) TAGS() []string         { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *negativeDuration) Name() string         { return "" }
func (d *negativeDuration) Abbreviation() string { return "" }
func (d *negativeDuration) Aliases() []string    { return []string{} }
