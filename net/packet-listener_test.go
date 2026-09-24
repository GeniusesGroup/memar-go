/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	net_p "memar/net/protocol"
)

var _ net_p.PacketListener = &PacketListener{}
