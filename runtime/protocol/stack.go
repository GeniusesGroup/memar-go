/* For license and copyright information please see the LEGAL file in the code repository */

package runtime_p

import (
	string_p "memar/string/protocol"
)

type Stack interface {
	// if log need to trace, specially in panic situation.
	// In Go USUALLY fill by `debug.Stack()`
	RuntimeStack() string_p.String
}

type Existence interface {
	Heap() bool // Pointer
}
