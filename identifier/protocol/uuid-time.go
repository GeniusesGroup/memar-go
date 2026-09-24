/* For license and copyright information please see the LEGAL file in the code repository */

package identifier_p

import (
	string_p "memar/codec/string/protocol"
	time_p "memar/time/protocol"
)

type UUID_Time interface {
	ExistenceTime() time_p.Time

	string_p.Stringer
}
