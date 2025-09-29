/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	log_p "memar/audit/log/protocol"
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
)

// TODO::: Can't force compiler to inline below functions, Delete file to force developers use this way:
// Logger.DispatchEvent(log.ConfEvent(&dt, "???"))

// Trace make new event with given level and add stack trace and log it to Logger
func Trace(dt datatype_p.DataType, level log_p.Level, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, level, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Info make new event with "Information" level and log it to Logger
func Info(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_Information, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// Notice make new event with "Notice" level and log it to Logger
func Notice(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_Notice, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// Error make new event with "Error" level and add stack trace and log it to Logger
func Error(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_Error, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Debug make new event with "Debug" level and log it to Logger
func Debug(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_Debug, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// DeepDebug make new event with "DeepDebug" level and log it to Logger
func DeepDebug(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_DeepDebug, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Warn make new event with "Warning" level and log it to Logger
func Warn(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_Warning, message, false)
	return Logger.DispatchEvent(&e.Event)
}

// Fatal make new event with "Fatal" level and log it to Logger
func Fatal(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_Fatal, message, true)
	return Logger.DispatchEvent(&e.Event)
}

// Conf make new event with "Confidential" level and log it to Logger
func Conf(dt datatype_p.DataType, message string) (err error_p.Error) {
	var e Event_UTF8
	e.Init(dt, log_p.Level_Confidential, message, false)
	return Logger.DispatchEvent(&e.Event)
}
