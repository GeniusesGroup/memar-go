/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"runtime/debug"

	"../protocol"
	"../time/unix"
)

func NewEvent(level protocol.LogType, domian, message string) (event *Event) {
	return &Event{
		level:   level,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func TraceEvent(level protocol.LogType, domian, message string) (event *Event) {
	return &Event{
		level:   level,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   debug.Stack(),
	}
}

func InfoEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_Information,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func NoticeEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_Notice,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func DebugEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_Debug,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func DeepDebugEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_DeepDebug,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func WarnEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_Warning,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

// FatalEvent return new event with panic level and added stack trace.
func PanicEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_Panic,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   debug.Stack(),
	}
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_Fatal,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   debug.Stack(),
	}
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.LogEvent_Confidential,
		time:    unix.Now(),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

// Event implement protocol.LogEvent
type Event struct {
	level   protocol.LogType
	time    unix.Time
	domain  string
	message string
	stack   []byte
}

func (e *Event) MainType() protocol.EventMainType { return protocol.EventMainType_Log }
func (e *Event) SubType() protocol.EventSubType   { return protocol.EventSubType(e.level) }
func (e *Event) Cancelable() bool                 { return false }
func (e *Event) DefaultPrevented() bool           { return false }
func (e *Event) Bubbles() bool                    { return false }
func (e *Event) PreventDefault()                  {}

func (e *Event) Level() protocol.LogType { return e.level }
func (e *Event) Time() protocol.Time     { return &e.time }
func (e *Event) Domain() string          { return e.domain }
func (e *Event) Message() string         { return e.message }
func (e *Event) Stack() []byte           { return e.stack }

/*
	-- protocol.Syllab interface Encoder & Decoder --
*/
func (e *Event) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(e.LenOfSyllabStack()) {
		// err = syllab.ErrShortArrayDecode
	}
	return
}
func (e *Event) FromSyllab(payload []byte, stackIndex uint32) {
}
func (e *Event) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	return
}
func (e *Event) LenAsSyllab() uint64      { return uint64(e.LenOfSyllabStack() + e.LenOfSyllabHeap()) }
func (e *Event) LenOfSyllabStack() uint32 { return 33 }
func (e *Event) LenOfSyllabHeap() (ln uint32) {
	return uint32(len(e.domain) + len(e.stack) + len(e.message))
}
