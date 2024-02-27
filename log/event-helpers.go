/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/protocol"
)

func NewEvent(domain protocol.MediaType, level protocol.LogLevel, message protocol.String) (event *Event) {
	//go:gc stack
	var e Event
	e.Init(domain, level, message, nil)
	return &e
}

func TraceEvent(domain protocol.MediaType, level protocol.LogLevel, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, level, message, nil)
	e.DefaultStack()
	return &e
}

func NoticeEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Notice, message, nil)
	return &e
}

func ErrorEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Error, message, nil)
	e.DefaultStack()
	return &e
}

func DebugEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Debug, message, nil)
	return &e
}

func DeepDebugEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_DeepDebug, message, nil)
	e.DefaultStack()
	return &e
}

func InfoEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Information, message, nil)
	return &e
}

func WarnEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Warning, message, nil)
	return &e
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Fatal, message, nil)
	e.DefaultStack()
	return &e
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(domain protocol.MediaType, message protocol.String) (event *Event) {
	var e Event
	e.Init(domain, protocol.LogLevel_Confidential, message, nil)
	return &e
}
