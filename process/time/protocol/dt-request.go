/* For license and copyright information please see the LEGAL file in the code repository */

package audit_time_p

import (
	time_p "memar/time/protocol"
)

// Time of the request not the save time of relation records.
type Field_Request interface {
	RequestedAt() time_p.Time
}
