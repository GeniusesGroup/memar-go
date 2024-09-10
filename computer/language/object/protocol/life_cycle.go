/* For license and copyright information please see the LEGAL file in the code repository */

package object_p

import (
	error_p "memar/error/protocol"
)

type LifeCycle interface {
	// new call to allocate the object and initialize it.
	// It must be called on pointer to the struct not direct use e.g. `t *embed`.
	// New()

	// Allocate()

	// ?? IncrementReference()
	// ?? DecrementReference()

	// Constructors should instantiate the fields of an object and do any other initialization necessary to make the object ready to use.
	// This is generally means constructors are small, but there are scenarios where this would be a substantial amount of work.
	// TODO::: how let custom initialize and get some args??
	// Init() (err Error)

	// TODO::: how let custom reinitialize and get some args??
	// Reinit() (err Error)

	Deinit() (err error_p.Error)

	// Deallocate()

	// Opposite of New()
	// Release(), Destroy(), Delete(), Drop()
	// remove, dispose, clear, close
}
