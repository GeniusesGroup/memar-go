/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

func NewEvent(level protocol.LogType, domain, message string) (event *Event) {
	var e Event
	e.Init(level, domain, message, false)
	return &e
}

func TraceEvent(level protocol.LogType, domain, message string) (event *Event) {
	var e Event
	e.Init(level, domain, message, true)
	return &e
}

func InfoEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_Information, domain, message, false)
	return &e
}

func NoticeEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_Notice, domain, message, false)
	return &e
}

func DebugEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_Debug, domain, message, false)
	return &e
}

func DeepDebugEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_DeepDebug, domain, message, false)
	return &e
}

func WarnEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_Warning, domain, message, false)
	return &e
}

// FatalEvent return new event with panic level and added stack trace.
func PanicEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_Panic, domain, message, true)
	return &e
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_Fatal, domain, message, true)
	return &e
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(domain, message string) (event *Event) {
	var e Event
	e.Init(protocol.LogEvent_Confidential, domain, message, false)
	return &e
}
