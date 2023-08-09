//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errTimerAlreadyInit) Domain() string  { return domainEnglish }
func (d *errTimerAlreadyInit) Summary() string { return "Timer Already Initialized" }
func (d *errTimerAlreadyInit) Overview() string {
	return "Don't initialize a timer twice. Use Reset() method to change the timer."
}
func (d *errTimerAlreadyInit) UserNote() string { return "" }
func (d *errTimerAlreadyInit) DevNote() string  { return "" }
func (d *errTimerAlreadyInit) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errTimerAlreadyInit) Name() string         { return "" }
func (d *errTimerAlreadyInit) Abbreviation() string { return "" }
func (d *errTimerAlreadyInit) Aliases() []string    { return []string{} }
