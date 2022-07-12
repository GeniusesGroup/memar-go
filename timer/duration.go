/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"../protocol"
)

// Common durations.
const (
	Nanosecond  protocol.Duration = 1
	Microsecond                   = 1000 * Nanosecond
	Millisecond                   = 1000 * Microsecond
	Second                        = 1000 * Millisecond
)