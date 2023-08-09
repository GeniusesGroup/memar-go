//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errNegativeDuration) Domain() string  { return domainEnglish }
func (d *errNegativeDuration) Summary() string { return "Negative Duration" }
func (d *errNegativeDuration) Overview() string {
	return "Timer or Ticker must have positive duration or interval."
}
func (d *errNegativeDuration) UserNote() string { return "" }
func (d *errNegativeDuration) DevNote() string  { return "" }
func (d *errNegativeDuration) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNegativeDuration) Name() string         { return "" }
func (d *errNegativeDuration) Abbreviation() string { return "" }
func (d *errNegativeDuration) Aliases() []string    { return []string{} }
