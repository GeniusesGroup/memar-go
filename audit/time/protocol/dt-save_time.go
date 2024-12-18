/* For license and copyright information please see the LEGAL file in the code repository */

package audit_time_p

import (
	time_p "memar/time/protocol"
)

// Request time or save Time of the request not the created record by this record.
type Field_SaveTime interface {
	SaveTime() time_p.Time
}
