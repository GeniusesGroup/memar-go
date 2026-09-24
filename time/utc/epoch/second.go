/* For license and copyright information please see the LEGAL file in the code repository */

package utc_epoch

type Field_SecondEpoch interface {
	SecondEpoch() SecondEpoch
}

type SecondEpoch uint8

const (
	SecondEpoch_Unset SecondEpoch = iota
	SecondEpoch_InMinute
	SecondEpoch_InHour
	SecondEpoch_InDay
	SecondEpoch_InWeek
	SecondEpoch_InMonth
	SecondEpoch_InSeason
	SecondEpoch_InYear
	SecondEpoch_InDecade
	SecondEpoch_InCentury
)
