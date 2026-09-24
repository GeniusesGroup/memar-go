//go:build !race

/* For license and copyright information please see the LEGAL file in the code repository */

package thread_sanitizer

import (
	"unsafe"
)

// Copy from `internal\race\norace.go`

const DetectorEnabled = false

const panicMsg = "Because DetectorEnabled is false, none of race functions should be called."

// Dummy race detection API, used when not built with -race.
// Because DetectorEnabled is false, none of these functions should be called.

func Acquire(addr unsafe.Pointer)                 { panic(panicMsg) }
func AcquireCTX(ctx uintptr, addr unsafe.Pointer) { panic(panicMsg) }
func Release(addr unsafe.Pointer)                 { panic(panicMsg) }
func ReleaseMerge(addr unsafe.Pointer)            { panic(panicMsg) }
func Disable()                                    { panic(panicMsg) }
func Enable()                                     { panic(panicMsg) }
func Read(addr unsafe.Pointer)                    { panic(panicMsg) }
func Write(addr unsafe.Pointer)                   { panic(panicMsg) }
func ReadRange(addr unsafe.Pointer, len int)      { panic(panicMsg) }
func WriteRange(addr unsafe.Pointer, len int)     { panic(panicMsg) }
func Errors() int                                 { panic(panicMsg) }
