/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	_ "unsafe" // for go:linkname
)

// now returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timeout, ...
// TODO::: move assembly logic from go/src/runtime to this package to prevent use unsafe package
//
//go:linkname now runtime.nanotime
func now() int64
