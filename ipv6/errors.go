/* For license and copyright information please see LEGAL file in repository */

package ipv6

import "../error"

// Declare IPv6 errors number
const (
	packetTooShort = 00000 + (iota + 1)
)

// Declare Errors Details
var (
	PacketTooShort = error.NewError("IPv6 packet is empty or too short than standard header", packetTooShort, 0)
)
