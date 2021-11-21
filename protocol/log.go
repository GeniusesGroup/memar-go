/* For license and copyright information please see LEGAL file in repository */

package protocol

// Logger
type Logger interface {
	PanicHandler()

	LogConfidential(any ...interface{})
	LogInfo(any ...interface{})
	LogWarn(any ...interface{})
	LogDebug(any ...interface{})
	LogDeepDebug(any ...interface{})
	LogPanic(any ...interface{})
	LogFatal(any ...interface{})
}
