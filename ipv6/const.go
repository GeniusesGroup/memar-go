/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

import "libgo/time/monotonic"

const (
	// Indicate max number of packet bundle in one buffer and processed
	maxEvents = 0

	// Max wait time a buffer fill
	timeout = 10 * monotonic.Millisecond

	// handle constants indicate IP must support their packets or drop them
	handleMode_Listener = false
)

const (
	// Version of protocol
	Version = 6

	// AddrLen address lengths 128 bit || 16 byte.
	AddrLen = 16

	// HeaderLen is minimum header length of IPv6 header
	HeaderLen = 40
)
