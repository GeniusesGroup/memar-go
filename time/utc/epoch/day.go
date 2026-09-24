/* For license and copyright information please see the LEGAL file in the code repository */

package utc_epoch

type Field_DayEpoch interface {
	DayEpoch() DayEpoch
}

type DayEpoch uint8

const (
	DayEpoch_Unset DayEpoch = iota
	DayEpoch_InWeek
	DayEpoch_InMonth
	DayEpoch_InSeason
	DayEpoch_InYear
	DayEpoch_InDecade
	DayEpoch_InCentury
)
