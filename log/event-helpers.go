/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/protocol"
)

func NewEvent(domain protocol.MediaType, level protocol.LogLevel, message string) (event *Event) {
	//go:gc stack
	var e Event
	e.Init(domain, level, message, false)
	return &e
}

func TraceEvent(domain protocol.MediaType, level protocol.LogLevel, message string) (event *Event) {
	var e Event
	e.Init(domain, level, message, true)
	return &e
}

func NoticeEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Notice, message, false)
	return &e
}

func ErrorEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Error, message, true)
	return &e
}

func DebugEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Debug, message, false)
	return &e
}

func DeepDebugEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_DeepDebug, message, true)
	return &e
}

func InfoEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Information, message, false)
	return &e
}

func WarnEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Warning, message, false)
	return &e
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Fatal, message, true)
	return &e
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(domain protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Confidential, message, false)
	return &e
}
