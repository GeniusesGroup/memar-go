/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"github.com/GeniusesGroup/libgo/time/monotonic"
)

// maxWhen is the maximum value for timer's when field.
const maxWhen monotonic.Time = 1<<63 - 1 // math.MaxInt64

// verifyTimers can be set to true to add debugging checks that the
// timer heaps are valid.
const verifyTimers = false

// The default heap ary is 4-ary. See siftUpTimer and siftDownTimer.
const heapAry = 4
