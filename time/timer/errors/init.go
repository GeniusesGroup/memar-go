/* For license and copyright information please see the LEGAL file in the code repository */

package errs

func init() {
	ErrTimerNotInit.Init()
	ErrTimerAlreadyInit.Init()
	ErrTimerAlreadyStarted.Init()
	ErrNegativeDuration.Init()
	ErrNegativePeriodNumber.Init()

	ErrTimerBadStatus.Init()
	ErrTimerRacyAccess.Init()
}
