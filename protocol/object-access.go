/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Object_Access or in some ways Object_Visibility indicate
type Object_Access uint32

// default is when flags not set
const (
	Object_Access_Unset Object_Access = 0 // or set defaults

	Object_Access_Mutable Object_Access = (1 << iota) // default:Immutable

	/*
		Encapsulation
	*/

	// means accessible from outside the first encapsulation(scope) it is live in it e.g.
	// - object member accessible from outside of the object
	// - object accessible from other packages
	// opposite means cannot be accessed (or viewed) from outside the structure
	Object_Access_OutsideOfFirstEncapsulation

	// means accessible from outside the second encapsulation(scope) it is live in it e.g.
	// - object member accessible from outside of the package
	Object_Access_OutsideOfSecondEncapsulation

	// protected - members cannot be accessed from outside the class,
	// however, they can be accessed in inherited structures.
	// - OutsideOfFirstEncapsulation+InheritedEncapsulation means just access when member inherited in other structure inside the same package
	// - OutsideOfSecondEncapsulation+InheritedEncapsulation means just access when member inherited in other structure in any packages
	Object_Access_InheritedEncapsulation

	/* Concurrency */

	// true if the member is safe to send it to another thread.
	// It depend on how and where object allocated. e.g.
	// Some times allocated object exist on a thread stack and
	// not safe to send to other thread without copy it to the global heap
	Object_Access_SafeToSend

	// true if the member is safe to share between threads(use concurrently).
	// it use some sync mechanism like atomic operation to let other access
	// false means it must access just in one thread.
	Object_Access_SafeToShare
)
