/* For license and copyright information please see the LEGAL file in the code repository */

package ipv4

import (
	"libgo/time/monotonic"
)

// https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
const (
	protocolNumber_tcp byte = 0x06
)

const (
	// Version of protocol
	Version = 4

	// AddrLen address lengths 32 bit equal 4 byte.
	AddrLen = 4

	// MinHeaderLen is minimum header length of IPv4 header
	MinHeaderLen = 40
)

const (
	// Indicate max number of packet bundle in one buffer and processed
	maxEvents = 0

	// Max wait time a buffer fill
	timeout = 10 * monotonic.Millisecond

	// handle constants indicate IP must support their packets or drop them
	handleMode_Listener = false
)
