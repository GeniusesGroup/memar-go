//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

//memar:impl memar/protocol.Detail
func (d *errTimerNotInit) Domain() string  { return domainEnglish }
func (d *errTimerNotInit) Summary() string { return "Timer Not Initialized" }
func (d *errTimerNotInit) Overview() string {
	return "Timer must initialized before Start(), Reset() or Stop()"
}
func (d *errTimerNotInit) UserNote() string { return "" }
func (d *errTimerNotInit) DevNote() string  { return "" }
func (d *errTimerNotInit) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (d *errTimerNotInit) Name() string         { return "" }
func (d *errTimerNotInit) Abbreviation() string { return "" }
func (d *errTimerNotInit) Aliases() []string    { return []string{} }
