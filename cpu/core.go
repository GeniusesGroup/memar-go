/* For license and copyright info please see LEGAL file in repository */

package cpu

import "runtime"

type CoreID uint64

func (id *CoreID) Active() {
	*id = CoreID(activeCoreID())
}

// activeCoreID or WhichCoreAmIOn return active core id that thread(goroutine) run on it.
func activeCoreID() uint64

// LogicalCount returns the number of logical CPUs usable by the current process.
func LogicalCount() uint { return uint(runtime.NumCPU()) }

// PhysicalCount returns the number of physical CPUs usable by the current process.
func PhysicalCount() uint { return uint(runtime.NumCPU()) }
