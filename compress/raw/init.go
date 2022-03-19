/* For license and copyright information please see LEGAL file in repository */

package raw

import (
	"../../protocol"
)

func init() {
	// Check due to os can be nil almost in tests and benchmarks build
	if protocol.OS != nil {
		protocol.OS.RegisterCompressType(&RAW)
	}
}
