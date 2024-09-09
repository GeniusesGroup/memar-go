/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/time/duration"
)

type Year int64

// Common durations.
const (
	NanoSecondInYear                 = 31556926 * duration.OneSecond
	SecondInYear     duration.Second = 31556926 // 365.24 days
)
