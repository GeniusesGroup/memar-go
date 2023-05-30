/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import "libgo/protocol"

func nsecToSec(d protocol.Duration) (sec int64, nsec int32) {
	sec = int64(d / Second)
	if sec == 0 {
		nsec = int32(d)
		return
	}
	var secPass = sec * int64(Second)
	nsec = int32(int64(d) % secPass)
	return
}
