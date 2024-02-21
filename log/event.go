/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"runtime/debug"

	"memar/convert"
	"memar/event"
	"memar/protocol"
	"memar/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	event.Event

	level   protocol.LogLevel
	message string
	stack   string
}

//memar:impl memar/protocol.LogEvent
func (e *Event) Level() protocol.LogLevel { return e.level }
func (e *Event) Message() string          { return e.message }
func (e *Event) Stack() string            { return e.stack }

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Event) Init(domain protocol.MediaType, level protocol.LogLevel, message string, stack bool) {
	e.level = level
	e.message = message
	if stack {
		e.stack = convert.UnsafeByteSliceToString(debug.Stack())
	}
	// TODO::: set NodeID???
	e.Event.Init(domain, unix.Now())
}
