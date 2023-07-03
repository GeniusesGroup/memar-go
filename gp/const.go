/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/protocol"
	"libgo/time/earth"
)

const (
	// AddrLen address lengths 32 bit equal 4 byte.
	AddrLen = 16

	// FrameLen is GP frame length.
	FrameLen = protocol.Network_FrameID_Length + AddrLen + AddrLen // 33 = 1+16+16
)

const (
	ConnectionIdleTimeout = 24 * earth.Hour // 24 hour
)
