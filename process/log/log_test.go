/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	log_p "memar/audit/log/protocol"
)

var _ log_p.Logger[*Event] = &Logger
