/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/time/duration"
)

type Month int64

// Common durations.
const (
	NanoSecondInMonth                 = 2629743 * duration.OneSecond
	SecondInMonth     duration.Second = 2629743 // 30.44 days
)
