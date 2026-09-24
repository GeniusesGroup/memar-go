/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	net_p "memar/net/protocol"
	"memar/time/earth"
)

const (
	// AddrLen address lengths 32 bit equal 4 byte.
	AddrLen = 16

	// FrameLen is GP frame length.
	FrameLen = net_p.FrameType_Length + AddrLen + AddrLen // 33 = 1+16+16
)

const (
	ConnectionIdleTimeout = earth.SecondInDay // 24 hour
)
