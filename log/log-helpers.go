/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/protocol"
)

// TODO::: Can't force compiler to inline below functions, Delete file to force developers use this way:
// Logger.DispatchEvent(log.ConfEvent(&dt, "???"))

// Trace make new event with given level and add stack trace and log it to Logger
func Trace(dt protocol.DataType, level protocol.LogLevel, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, level, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Info make new event with "Information" level and log it to Logger
func Info(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_Information, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// Notice make new event with "Notice" level and log it to Logger
func Notice(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_Notice, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// Error make new event with "Error" level and add stack trace and log it to Logger
func Error(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_Error, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Debug make new event with "Debug" level and log it to Logger
func Debug(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_Debug, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// DeepDebug make new event with "DeepDebug" level and log it to Logger
func DeepDebug(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_DeepDebug, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Warn make new event with "Warning" level and log it to Logger
func Warn(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_Warning, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// Fatal make new event with "Fatal" level and log it to Logger
func Fatal(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_Fatal, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Conf make new event with "Confidential" level and log it to Logger
func Conf(dt protocol.DataType, message string) (err protocol.Error) {
	var e Event_UTF8
	e.Init(dt, protocol.LogLevel_Confidential, message, false)
	return Logger.DispatchEvent(&e.Event)
}
