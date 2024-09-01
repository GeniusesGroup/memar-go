/* For license and copyright information please see the LEGAL file in the code repository */

package log

// Application immutable runtime settings
const (

	// CNF_DevelopingMode use to indicate that app can do some more logic e.g.
	// - Save more logs
	// - Add more services like net/http/pprof for better debugging
	// - Add more pages that just need only for developers
	CNF_DevelopingMode = true
)
