/* For license and copyright information please see the LEGAL file in the code repository */

package runtime_p

// Async indicate any object need to run in a new thread,
// due to have blocking mechanism in its logic.
type Async interface {
	DoAsync()
}
