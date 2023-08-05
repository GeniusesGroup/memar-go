/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	er "memar/error"
)

// Errors
var (
	ErrSegmentTooShort    er.Error
	ErrSegmentWrongLength er.Error
)

func init() {
	ErrSegmentTooShort.Init("domain/tcp.wg.ietf.org; type=error; name=packet-too-short")
	ErrSegmentWrongLength.Init("domain/tcp.wg.ietf.org; type=error; name=packet-wrong-length")
}
