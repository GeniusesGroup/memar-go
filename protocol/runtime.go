/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Runtime interface {
}

type Runtime_Stack interface {
	// if log need to trace, specially in panic situation.
	// In Go USUALLY fill by `debug.Stack()`
	RuntimeStack() String
}
