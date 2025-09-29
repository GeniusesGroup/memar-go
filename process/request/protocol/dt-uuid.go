/* For license and copyright information please see the LEGAL file in the code repository */

package request_p

import (
	record_p "memar/storage/record/protocol"
)

//	Field audit request (logging) is the process of documenting activity within the software systems used across your organization.
//
// Audit logs record the occurrence of an event, the time at which it occurred, the responsible user or service,
// and the impacted entity. All of the devices in your network, your cloud services,
// and your applications emit logs that may be used for auditing purposes.
type Field_RequestID interface {
	RequestID() UUID
}

type Field_RequestIDs interface {
	RequestsID() []UUID
}

// Request RecordID to prove how domain objects created in the chain of other objects.
type UUID record_p.UUID
