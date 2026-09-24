/* For license and copyright information please see the LEGAL file in the code repository */

package utc_epoch

type Field_HourEpoch interface {
	HourEpoch() HourEpoch
}

type HourEpoch uint8

const (
	HourEpoch_Unset HourEpoch = iota
	HourEpoch_InDay
	HourEpoch_InWeek
	HourEpoch_InMonth
	HourEpoch_InSeason
	HourEpoch_InYear
	HourEpoch_InCentury
)
