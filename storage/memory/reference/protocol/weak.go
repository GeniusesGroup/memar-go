/* For license and copyright information please see the LEGAL file in the code repository */

package reference_p

// Weak reference is a reference that does not protect the referenced object
// from collection by a garbage collector, unlike a strong reference.
// 
// In other words, weak models temporary ownership, when an object needs to be accessed only if it exists,
// and it may be deleted at any time by someone else,
// 
// https://en.wikipedia.org/wiki/Weak_reference
//
// Other protocols:::
// https://en.cppreference.com/w/cpp/memory/weak_ptr.html
// https://learn.microsoft.com/en-us/dotnet/api/system.weakreference
// https://docs.oracle.com/javase/8/docs/api/java/lang/ref/WeakReference.html
// https://docs.python.org/3/library/weakref.html
type Weak interface {
	// Gets an indication whether the object referenced by the current WeakReference object has been garbage collected.
	// It will return error instead of simple boolean to support more conditions.
	// All returns errors can also use by `Dereference()` method.
	ReferenceAlive() (err error_p.Error)
}
