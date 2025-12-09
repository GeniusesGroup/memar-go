/* For license and copyright information please see the LEGAL file in the code repository */

package time_p

import (
	"memar/time/duration"
)

type Field_Time interface {
	Time() Time
}

// Time is the interface that must implement by any time capsule.
// It is base on Epoch and Second terms to work anywhere (in any planet in the universe).
type Time interface {
	Field_Epoch
	SecondElapsed() duration.Second             // From Epoch
	NanoInSecondElapsed() duration.NanoInSecond // From second
}

// Time_Stringer is base on other factor than Time like timezone, ...
type Time_Stringer[STR string_p.String] interface {
	// Stringer MUST include factors like timezone, ...
	string_p.Stringer[STR]
}
