/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	event_p "memar/event/protocol"
)

var _ event_p.Target[*Event] = &Target[*Event]{}
