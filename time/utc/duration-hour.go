/* For license and copyright information please see the LEGAL file in the code repository */

package utc

import (
	"memar/time/earth"
)

type Field_HourElapsed interface {
	// Epoch is not indicate here due to utc has unique one.
	UTC_HourElapsed() HourElapsed
}

type HourElapsed earth.Hour
