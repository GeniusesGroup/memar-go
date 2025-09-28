/* For license and copyright information please see the LEGAL file in the code repository */

package event_p

import (
	datatype_p "memar/datatype/protocol"
	logic_p "memar/math/logic/protocol"
	time_p "memar/time/protocol"
)

// Events CAN be any data-type in any domain.
// Event are runtime data-types that DON't need to store in storage layer.
// 
// Events dispatch after base logic that make this Event.
// e.g. When a service request want to serve, a handler receive the request and decide to save the request for auditing
// after validate request codec and save request, a related event dispatch.
// This service can provide some mechanism to change the flow of request processing
//
// Other frameworks:
// https://www.w3.org/TR/DOM-Level-3-Events/#event-flow
// https://developer.mozilla.org/en-US/docs/Web/API/Event
// https://developer.mozilla.org/en-US/docs/Web/Events
// https://learn.microsoft.com/en-us/dotnet/api/system.componentmodel.eventdescriptor
type Event interface {
	datatype_p.Field_DataType
	time_p.Field_Time
	logic_p.Equivalence[Event]
}
