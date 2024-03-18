/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	"memar/protocol"
)

var _ protocol.EventTarget[*Event, Options] = &EventTarget[*Event]{}
