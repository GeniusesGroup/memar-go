/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type ObjectLifeCycle interface {
	// new call to allocate the object and initialize it.
	// It must be called on pointer to the struct not direct use e.g. `t *embed`.
	// New()

	// Allocate()

	// ?? IncrementReference()
	// ?? DecrementReference()

	// TODO::: how let custom initialize and get some args??
	// Init() (err Error)

	// TODO::: how let custom reinitialize and get some args??
	// Reinit() (err Error)

	Deinit() (err Error)

	// Deallocate()
}
