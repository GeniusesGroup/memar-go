//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

//memar:impl memar/protocol.Detail
func (d *TimerAlreadyInit) Domain() string  { return domainEnglish }
func (d *TimerAlreadyInit) Summary() string { return "Timer Already Initialized" }
func (d *TimerAlreadyInit) Overview() string {
	return "Don't initialize a timer twice. Use Reset() method to change the timer."
}
func (d *TimerAlreadyInit) UserActionNote() string { return "" }
func (d *TimerAlreadyInit) DevActionNote() string  { return "" }
func (d *errTimerAlreadyInit) TAGS() []string      { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *TimerAlreadyInit) Name() string         { return "" }
func (d *TimerAlreadyInit) Abbreviation() string { return "" }
func (d *TimerAlreadyInit) Aliases() []string    { return []string{} }
