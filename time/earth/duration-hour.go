/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/math/integer"
	"memar/time/duration"
)

type Field_Hour interface {
	// Epoch is not indicate here and MUST indicate by the user.
	Hour() Hour
}

type Hour integer.S64

// Common durations.
const (
	NanoSecondInHour duration.NanoSecond = 60 * NanoSecondInMinute
	SecondInHour     duration.Second     = 60 * SecondInMinute
	MinuteInHour     Minute              = 60
)

func (h *Hour) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*h = Hour(d / NanoSecondInHour)
}

func (h *Hour) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*h = Hour(d / SecondInHour)
}
