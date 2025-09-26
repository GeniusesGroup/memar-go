//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errNegativePeriodNumber) Domain() string  { return domainEnglish }
func (d *errNegativePeriodNumber) Summary() string { return "Negative Period Number" }
func (d *errNegativePeriodNumber) Overview() string {
	return "periodNumber must be more than one on LimitTicker."
}
func (d *errNegativePeriodNumber) UserNote() string { return "" }
func (d *errNegativePeriodNumber) DevNote() string  { return "" }
func (d *errNegativePeriodNumber) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errNegativePeriodNumber) Name() string         { return "" }
func (d *errNegativePeriodNumber) Abbreviation() string { return "" }
func (d *errNegativePeriodNumber) Aliases() []string    { return []string{} }
