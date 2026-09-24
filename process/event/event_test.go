/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	event_p "memar/process/event/protocol"
)

var _ event_p.Event = &Event{}
