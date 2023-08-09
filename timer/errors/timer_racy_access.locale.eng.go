//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errTimerRacyAccess) Domain() string  { return domainEnglish }
func (d *errTimerRacyAccess) Summary() string { return "Racy Access" }
func (d *errTimerRacyAccess) Overview() string {
	return "Timer fields must not change illegally or called it's method concurrently."
}
func (d *errTimerRacyAccess) UserNote() string { return "data corruption, maybe racy use of timers" }
func (d *errTimerRacyAccess) DevNote() string {
	return `The timer data structures have been corrupted, presumably due to racy use by the program.
dispatch log event here rather than panicing due to invalid slice access while holding locks.
See issue https://github.com/golang/go/issues/25686`
}
func (d *errTimerRacyAccess) TAGS() []string { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errTimerRacyAccess) Name() string         { return "" }
func (d *errTimerRacyAccess) Abbreviation() string { return "" }
func (d *errTimerRacyAccess) Aliases() []string    { return []string{} }
