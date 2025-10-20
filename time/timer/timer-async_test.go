/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"memar/time/monotonic"
	timer_p "memar/time/timer/protocol"
)

var _ timer_p.Timer[monotonic.Time, Status] = &Async{}
