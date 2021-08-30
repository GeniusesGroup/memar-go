/* For license and copyright information please see LEGAL file in repository */

package protocol

const BaseUnixTime = "00:00:00 UTC on 1 January 1970"

// Time is the interface that must implement by any time object!
// Time codec len must always be 64bit same as int64
type Time interface {
	Unix() int64
	Pass(baseTime Time) (pass bool)
	// AddDuration(d Duration) (new Time)

	// Codec
}

// Duration is the interface that must implement by any duration object!
type Duration interface {
}
