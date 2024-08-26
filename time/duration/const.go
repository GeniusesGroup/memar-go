/* For license and copyright information please see the LEGAL file in the code repository */

package duration

// Common durations.
const (
	OneNanosecond  NanoSecond = 1
	OneMicrosecond            = 1000 * OneNanosecond  // 1e3
	OneMillisecond            = 1000 * OneMicrosecond // 1e6
	OneSecond                 = 1000 * OneMillisecond // 1e9
)
