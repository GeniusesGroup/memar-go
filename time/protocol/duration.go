/* For license and copyright information please see the LEGAL file in the code repository */

package time_p

import (
	string_p "memar/string/protocol"
)

// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
// The representation limits the largest representable duration to approximately 290 earth years.
type Duration[STR string_p.String] interface {
	string_p.Stringer_To[STR]
}
