/* For license and copyright information please see the LEGAL file in the code repository */

package user_p

import (
	ut16 "memar/identifier/uuid-time-16byte"
	time_p "memar/time/protocol"
)

type UUID = ut16.UUID

type ID interface {
	UUID() UUID

	// Below fields extract from(part of) above UUID
	ExistenceTime() time_p.Time
	Type() Type
	ID() [3]byte
}
