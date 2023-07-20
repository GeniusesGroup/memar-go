//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"libgo/detail"
	"libgo/protocol"
)

const domainEnglish = "Timer"

func init() {
	ErrTimerNotInit.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("Timer Not Initialized").
			SetOverview("Timer must initialized before Start(), Reset() or Stop()").
			SetUserNote("").
			SetDevNote("").
			SetTAGS([]string{})
		)
	ErrTimerAlreadyInit.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("Time Already Initialized").
			SetOverview("Don't initialize a timer twice. Use Reset() method to change the timer.").
			SetUserNote("").
			SetDevNote("").
			SetTAGS([]string{})
		)
	ErrTimerAlreadyStarted.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("Timer Already Started").
			SetOverview("Start called with started timer").
			SetUserNote("").
			SetDevNote("").
			SetTAGS([]string{})
		)

	ErrNegativeDuration.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("Negative Duration").
			SetOverview("Timer or Ticker must have positive duration or interval.").
			SetUserNote("").
			SetDevNote("").
			SetTAGS([]string{})
		)
	ErrNegativePeriodNumber.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("Negative Period Number").
			SetOverview("periodNumber must be more than one on LimitTicker.").
			SetUserNote("").
			SetDevNote("").
			SetTAGS([]string{})
		)


	ErrTimerBadStatus.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("Bad Status").
			SetOverview("Timer or Ticker is in a bad status and we don't process your request to Start, Stop or Reset it").
			SetUserNote("").
			SetDevNote("").
			SetTAGS([]string{})
		)
	ErrTimerRacyAccess.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
			SetName("").
			SetAbbreviation("").
			SetAliases([]string{}).
			SetSummary("Racy Access").
			SetOverview("Timer fields must not change illegally or called it's method concurrently.").
			SetUserNote("data corruption, maybe racy use of timers").
			SetDevNote(`The timer data structures have been corrupted, presumably due to racy use by the program.
dispatch log event here rather than panicing due to invalid slice access while holding locks.
See issue https://github.com/golang/go/issues/25686`).
			SetTAGS([]string{})
		)
}
