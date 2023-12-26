/* For license and copyright info please see LEGAL file in repository */

package cpu

import "runtime"

// ActiveCoreID or WhichCoreAmIOn return active core id that thread(goroutine) run on it.
func ActiveCoreID() uint64

// LogicalCount returns the number of logical CPUs usable by the current process.
func LogicalCount() uint { return uint(runtime.NumCPU()) }

// PhysicalCount returns the number of physical CPUs usable by the current process.
func PhysicalCount() uint { return uint(runtime.NumCPU()) }
