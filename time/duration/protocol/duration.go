/* For license and copyright information please see the LEGAL file in the code repository */

package duration_p

import (
	string_p "memar/codec/string/protocol"
)

// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
// The representation limits the largest representable duration to approximately 290 earth years.
type Duration interface {
	string_p.Stringer_To
}
