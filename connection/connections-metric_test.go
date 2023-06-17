/* For license and copyright information please see the LEGAL file in the code repository */

package connection

import (
	"libgo/protocol"
)

var _ protocol.ConnectionsMetrics = &ConnectionsMetric{}
