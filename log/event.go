/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"runtime/debug"
	"time"

	"../protocol"
)

func NewEvent(level protocol.LogType, domian, message string) (event *Event) {
	return &Event{
		level:   level,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func TraceEvent(level protocol.LogType, domian, message string) (event *Event) {
	return &Event{
		level:   level,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   debug.Stack(),
	}
}

func InfoEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Information,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func NoticeEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Notice,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func DebugEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Debug,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func DeepDebugEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_DeepDebug,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

func WarnEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Warning,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

// FatalEvent return new event with panic level and added stack trace.
func PanicEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Panic,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   debug.Stack(),
	}
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Fatal,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   debug.Stack(),
	}
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Confidential,
		time:    protocol.TimeUnixMilli(time.Now().UnixMilli()),
		domain:  domian,
		message: message,
		stack:   nil,
	}
}

// Event implement protocol.LogEvent
type Event struct {
	level   protocol.LogType
	time    protocol.TimeUnixMilli
	domain  string
	message string
	stack   []byte
}

func (e *Event) Level() protocol.LogType      { return e.level }
func (e *Event) Time() protocol.TimeUnixMilli { return e.time }
func (e *Event) Domain() string               { return e.domain }
func (e *Event) Message() string              { return e.message }
func (e *Event) Stack() []byte                { return e.stack }

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
