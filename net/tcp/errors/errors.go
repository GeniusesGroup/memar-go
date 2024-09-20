/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
)

// Errors
var (
	SegmentTooShort    er.Error
	SegmentWrongLength er.Error
)

func init() {
	SegmentTooShort.Init("domain/tcp.wg.ietf.org; type=error; name=packet-too-short")
	SegmentWrongLength.Init("domain/tcp.wg.ietf.org; type=error; name=packet-wrong-length")
}
