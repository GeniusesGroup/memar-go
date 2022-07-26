/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// TODO::: Can't force compiler to inline below functions, Delete file to force developers use this way:
// protocol.App.Log(log.ConfEvent(domainEnglish, "???"))

// Trace make new event with given level and add stack trace and log it to protocol.App
func Trace(level protocol.LogType, domain, message string) (err protocol.Error) {
	// var e Event
	// e.Init(level, domain, message, true)
	return protocol.App.Log(TraceEvent(level, domain, message))
}

// Info make new event with "Information" level and log it to protocol.App
func Info(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_Information, domain, message, false)
	return protocol.App.Log(&e)
}

// Notice make new event with "Notice" level and log it to protocol.App
func Notice(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_Notice, domain, message, false)
	return protocol.App.Log(&e)
}

// Debug make new event with "Debug" level and log it to protocol.App
func Debug(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_Debug, domain, message, false)
	return protocol.App.Log(&e)
}

// DeepDebug make new event with "DeepDebug" level and log it to protocol.App
func DeepDebug(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_DeepDebug, domain, message, false)
	return protocol.App.Log(&e)
}

// Warn make new event with "Warning" level and log it to protocol.App
func Warn(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_Warning, domain, message, false)
	return protocol.App.Log(&e)
}

// Panic make new event with "Panic" level and log it to protocol.App
func Panic(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_Panic, domain, message, true)
	return protocol.App.Log(&e)
}

// Fatal make new event with "Fatal" level and log it to protocol.App
func Fatal(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_Fatal, domain, message, true)
	return protocol.App.Log(&e)
}

// Conf make new event with "Confidential" level and log it to protocol.App
func Conf(domain, message string) (err protocol.Error) {
	var e Event
	e.Init(protocol.LogEvent_Confidential, domain, message, false)
	return protocol.App.Log(&e)
}
