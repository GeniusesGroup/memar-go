/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// DataType_Access or in some ways DataType_Visibility indicate
type DataType_Access uint32

// default is when flags not set
const (
	DataType_Access_Unset DataType_Access = 0 // or set defaults

	DataType_Access_Mutable DataType_Access = (1 << iota) // default:Immutable

	/*
		Encapsulation
	*/

	// means accessible from outside the first encapsulation(scope) it is live in it e.g.
	// - object member accessible from outside of the object
	// - object accessible from other packages
	// opposite means cannot be accessed (or viewed) from outside the structure
	DataType_Access_OutsideOfFirstEncapsulation

	// means accessible from outside the second encapsulation(scope) it is live in it e.g.
	// - object member accessible from outside of the package
	DataType_Access_OutsideOfSecondEncapsulation

	// protected - members cannot be accessed from outside the class,
	// however, they can be accessed in inherited structures.
	// - OutsideOfFirstEncapsulation+InheritedEncapsulation means just access when member inherited in other structure inside the same package
	// - OutsideOfSecondEncapsulation+InheritedEncapsulation means just access when member inherited in other structure in any packages
	DataType_Access_InheritedEncapsulation

	/* Concurrency */

	// true if the member is safe to send it to another thread.
	// It depend on how and where object allocated. e.g.
	// Some times allocated object exist on a thread stack and
	// not safe to send to other thread without copy it to the global heap
	DataType_Access_SafeToSend

	// true if the member is safe to share between threads(use concurrently).
	// it use some sync mechanism like atomic operation to let other access
	// false means it must access just in one thread.
	DataType_Access_SafeToShare
)
