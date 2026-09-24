/* For license and copyright information please see the LEGAL file in the code repository */

package record_p

import (
	request_p "memar/process/request/protocol"
	audit_time_p "memar/process/time/protocol"

	// datatype_p "org/modules/computer/datatype/protocol"
)

// Store to know all request created by specific rule.
type DataModel interface {
	Field_RecordUUID
	Field_Data

	// datatype_p.Field_DataTypeID
	request_p.Field_RequestID
	audit_time_p.Field_Request
}
