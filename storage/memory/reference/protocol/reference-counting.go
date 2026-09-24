/* For license and copyright information please see the LEGAL file in the code repository */

package reference_p

import (
	error_p "memar/process/error/protocol"
)

// In computer science, reference counting is a programming technique of storing the number of references,
// pointers, or handles to a resource, such as an object, a block of memory, disk space, and others.
//
// In garbage collection algorithms, reference counts may be used to deallocate objects that are no longer needed.
//
// https://en.wikipedia.org/wiki/Reference_counting
// https://doc.rust-lang.org/std/rc/struct.Rc.html
type ReferenceCounting interface {
	IncrementReference() (err error_p.Error)
	DecrementReference() (err error_p.Error)
}

//  weak_count
