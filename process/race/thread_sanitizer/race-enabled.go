//go:build race

/* For license and copyright information please see the LEGAL file in the code repository */

package thread_sanitizer

import (
	"unsafe"
)

// Copy from `internal\race\race.go`

// Public race detection API, present iff build with -race.
const DetectorEnabled = true

//go:linkname Acquire runtime.RaceAcquire
func Acquire(addr unsafe.Pointer)

//go:linkname AcquireCTX runtime.Acquirectx
func AcquireCTX(ctx uintptr, addr unsafe.Pointer)

//go:linkname Release runtime.RaceRelease
func Release(addr unsafe.Pointer)

//go:linkname ReleaseMerge runtime.RaceReleaseMerge
func ReleaseMerge(addr unsafe.Pointer)

//go:linkname Disable runtime.RaceDisable
func Disable()

//go:linkname Enable runtime.RaceEnable
func Enable()

//go:linkname Read runtime.RaceRead
func Read(addr unsafe.Pointer)

//go:linkname Write runtime.RaceWrite
func Write(addr unsafe.Pointer)

//go:linkname ReadRange runtime.RaceReadRange
func ReadRange(addr unsafe.Pointer, len int)

//go:linkname WriteRange runtime.RaceWriteRange
func WriteRange(addr unsafe.Pointer, len int)

//go:linkname Errors runtime.RaceErrors
func Errors() int
