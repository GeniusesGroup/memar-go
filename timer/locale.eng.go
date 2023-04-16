//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"libgo/protocol"
)

const domainEnglish = "Timer"

func init() {
	ErrTimerNotInit.SetDetail(protocol.LanguageEnglish, domainEnglish, "Timer Not Initialized",
		"Timer must initialized before Start(), Reset() or Stop()",
		"",
		"",
		nil)
	ErrTimerAlreadyInit.SetDetail(protocol.LanguageEnglish, domainEnglish, "Time Already Initialized",
		"Don't initialize a timer twice. Use Reset() method to change the timer.",
		"",
		"",
		nil)
	ErrTimerAlreadyStarted.SetDetail(protocol.LanguageEnglish, domainEnglish, "Timer Already Started",
		"Start called with started timer",
		"",
		"",
		nil)
	ErrNegativeDuration.SetDetail(protocol.LanguageEnglish, domainEnglish, "Negative Duration",
		"Timer or Ticker must have positive duration or interval.",
		"",
		"",
		nil)
	ErrNegativePeriodNumber.SetDetail(protocol.LanguageEnglish, domainEnglish, "Negative Period Number",
		"periodNumber must be more than one on LimitTicker.",
		"",
		"",
		nil)

	ErrTimerBadStatus.SetDetail(protocol.LanguageEnglish, domainEnglish, "Bad Status",
		"Timer or Ticker is in a bad status and we don't process your request to Start, Stop or Reset it",
		"",
		"",
		nil)
	ErrTimerRacyAccess.SetDetail(protocol.LanguageEnglish, domainEnglish, "Racy Access",
		"Timer fields must not change illegally or called it's method concurrently.",
		"",
		"",
		nil)
}
