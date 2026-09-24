/* For license and copyright information please see the LEGAL file in the code repository */

package utc

import (
	"memar/time/earth"
)

type Field_DayElapsed interface {
	// Epoch is not indicate here due to utc has unique one.
	UTC_DayElapsed() DayElapsed
}

// DayElapsed show number of days pass from specific epoch.
// e.g. `unix.Now().DayElapsed()` return number of days from unix epoch.
type DayElapsed earth.Day
