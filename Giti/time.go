/* For license and copyright information please see LEGAL file in repository */

package giti

const BaseUnixTime = "00:00:00 UTC on 1 January 1970"

// Time is the interface that must implement by any time object!
type Time interface {
	Pass(baseTime Time) (pass bool)
	// AddDuration(d Duration) (new Time)
}

// Duration is the interface that must implement by any duration object!
type Duration interface {
}
