/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// Common durations.
const (
	Nanosecond  protocol.Duration = 1
	Microsecond                   = 1000 * Nanosecond
	Millisecond                   = 1000 * Microsecond
	Second                        = 1000 * Millisecond
)
