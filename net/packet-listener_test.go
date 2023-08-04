/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
)

var _ protocol.Network_PacketListener = &PacketListener{}
