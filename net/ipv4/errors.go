/* For license and copyright information please see the LEGAL file in the code repository */

package ipv4

import (
	er "memar/error"
)

// Errors
var (
	ErrPacketTooShort    er.Error
	ErrPacketWrongLength er.Error
)

func init() {
	ErrPacketTooShort.Init("domain/ipv4.wg.ietf.org; type=error; name=packet-too-short")
	ErrPacketWrongLength.Init("domain/ipv4.wg.ietf.org; type=error; name=packet-wrong-length")
}
