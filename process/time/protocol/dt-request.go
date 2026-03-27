/* For license and copyright information please see the LEGAL file in the code repository */

package audit_time_p

import (
	time_p "memar/time/protocol"
)

// Time of the request called.
// request CAN be in many layer:
// - In storage layer, it is the save time of relation records.
// - In business layer, It is the time user call related service.
type Field_Request interface {
	RequestedAt() time_p.Time
}
