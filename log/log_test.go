/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/event"
	"memar/log/protocol"
)

var _ log_p.Logger[*Event, event.Options] = &Logger
