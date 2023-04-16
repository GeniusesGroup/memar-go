/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	er "libgo/error"
)

// Errors
var (
	ErrTimerNotInit         er.Error
	ErrTimerAlreadyInit     er.Error
	ErrTimerAlreadyStarted  er.Error
	ErrNegativeDuration     er.Error
	ErrNegativePeriodNumber er.Error
	
	ErrTimerBadStatus  er.Error
	ErrTimerRacyAccess er.Error
)

func init() {
	ErrTimerNotInit.Init("domain/timer.protocol; type=error; name=timer-not-initialize")
	ErrTimerAlreadyInit.Init("domain/timer.protocol; type=error; name=timer-already-initialized")
	ErrTimerAlreadyStarted.Init("domain/timer.protocol; type=error; name=timer-already-started")
	ErrNegativeDuration.Init("domain/timer.protocol; type=error; name=negative-duration")
	ErrNegativePeriodNumber.Init("domain/timer.protocol; type=error; name=negative-period-number")

	ErrTimerBadStatus.Init("domain/timer.protocol; type=error; name=timer-bad-status")
	ErrTimerRacyAccess.Init("domain/timer.protocol; type=error; name=timer-racy-access")
}
