/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

import (
	er "libgo/error"
)

// Errors
var (
	ErrPacketTooShort er.Error
)

func init() {
	ErrPacketTooShort.Init("domain/ipv6.wg.ietf.org; type=error; name=packet-too-short")
}
