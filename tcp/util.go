/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"time"

	"github.com/GeniusesGroup/libgo/protocol"
)

func getDuration(t time.Time) (d protocol.Duration) {
	if !t.IsZero() {
		d = protocol.Duration(time.Until(t))
		if d == 0 {
			d = -1 // don't confuse deadline right now with no deadline
		}
	}
	return
}
