/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

var _ protocol.NetworkInterface = &BridgePort{}
