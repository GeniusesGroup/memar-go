/* For license and copyright information please see the LEGAL file in the code repository */

package time_p

import (
	string_p "memar/string/protocol"
	"memar/time/duration"
)

// Time is the interface that must implement by any time capsule.
// It is base on Epoch and Second terms to work anywhere (in any planet in the universe).
type Time interface {
	Epoch() Epoch
	SecondElapsed() duration.Second             // From Epoch
	NanoInSecondElapsed() duration.NanoInSecond // From second
}

// Time_Stringer is base on other factor than Time like timezone, ...
type Time_Stringer[STR string_p.String] interface {
	string_p.Stringer[STR]
}
