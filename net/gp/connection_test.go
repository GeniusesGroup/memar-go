/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	net_p "memar/net/protocol"
)

var _ net_p.OSI_Network = &Connection{}
