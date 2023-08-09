//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errTimerAlreadyStarted) Domain() string  { return domainEnglish }
func (d *errTimerAlreadyStarted) Summary() string { return "Timer Already Started" }
func (d *errTimerAlreadyStarted) Overview() string {
	return "Start called with started timer"
}
func (d *errTimerAlreadyStarted) UserNote() string { return "" }
func (d *errTimerAlreadyStarted) DevNote() string  { return "" }
func (d *errTimerAlreadyStarted) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errTimerAlreadyStarted) Name() string         { return "" }
func (d *errTimerAlreadyStarted) Abbreviation() string { return "" }
func (d *errTimerAlreadyStarted) Aliases() []string    { return []string{} }
