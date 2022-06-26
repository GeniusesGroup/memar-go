/* For license and copyright information please see LEGAL file in repository */

package timer

// maxWhen is the maximum value for timer's when field.
const maxWhen int64 = 1<<63 - 1 // math.MaxInt64

// verifyTimers can be set to true to add debugging checks that the
// timer heaps are valid.
const verifyTimers = false
