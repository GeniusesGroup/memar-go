/* For license and copyright information please see the LEGAL file in the code repository */

package audit_time_p

import (
	time_p "memar/time/protocol"
)

// Time of the data is modifying.
type Field_Modification interface {
	ModifiedAt() time_p.Time
}
