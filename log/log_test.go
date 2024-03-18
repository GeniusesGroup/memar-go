/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/event"
	"memar/protocol"
)

var _ protocol.Logger[*Event, event.Options] = &Logger
