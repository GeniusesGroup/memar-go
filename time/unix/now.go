/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import (
	_ "unsafe" // for go:linkname
)

// Provided by package runtime.
// TODO::: move assembly logic from go/src/runtime to this package to prevent use unsafe package
//
//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)
