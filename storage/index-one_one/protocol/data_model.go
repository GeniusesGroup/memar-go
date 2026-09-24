/* For license and copyright information please see the LEGAL file in the code repository */

package pehrest_p

import (
	request_p "memar/process/request/protocol"
	audit_time_p "memar/process/time/protocol"
	record_p "memar/storage/record/protocol"
)

// Store to know all request created by specific rule.
type DataModel interface {
	Key
	record_p.Field_RecordUUID

	request_p.Field_RequestID
	audit_time_p.Field_Request
}
