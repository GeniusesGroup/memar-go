/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	"memar/protocol"
	"memar/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	domain protocol.MediaType
	time   unix.Time
}

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Event) Init(domain protocol.MediaType, time unix.Time) {
	e.domain = domain
	e.time = time
}

//memar:impl memar/protocol.Event
func (e *Event) Domain() protocol.MediaType { return e.domain }
func (e *Event) Time() protocol.Time        { return &e.time }
func (e *Event) Cancelable() bool           { return false }
func (e *Event) DefaultPrevented() bool     { return false }
func (e *Event) Bubbles() bool              { return false }

//memar:impl memar/protocol.Event_Methods
func (e *Event) PreventDefault() {}
