/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type DataType_DefaultValue[T any] interface {
	Default() T
	SetDefault() // default value
}

type DataType_Accessor[T any] interface {
	Get() T

	// It will check(validate) given value and return proper error for
	Set(new T) (err Error)
}

type DataType_AtomicAccessor[T any] interface {
	Load() T
	Store(new T) (err Error)
	Swap(new T) (old T, err Error)
	CompareAndSwap(old, new T) (err Error) // return more than swapped(bool)
}

// DataType_Clone is explicit, may be expensive, and may be re-implement arbitrarily.
// Clone is designed for arbitrary duplications:
// a Clone implementation for a type T can do arbitrarily complicated operations required to create a new T.
// It is a normal trait (other than being in the prelude), and so requires being used like a normal trait, with method calls, etc.
type DataType_Clone[T any] interface {
	// Returns a clone of the itself
	Clone() (c T, err Error)
	// Performs clone-assignment from source
	CloneFrom(source T) (err Error)
	// Performs clone-assignment to destination
	CloneTo(destination T) (err Error)
}

// DataType_Copy is implicit, inexpensive, and cannot be re-implemented (memcpy).
// The Copy trait represents values that can be safely duplicated via memcpy:
// things like reassignments and passing an argument by-value to a function are always memcpys, and so for Copy types,
// the compiler understands that it doesn't need to consider those a move.
// every Copy type is also required to be Clone
type DataType_Copy[T any] interface {
	// Returns a copy of the itself
	Copy() (c T, err Error)
	// Performs copy-assignment from source
	CopyFrom(source T) (err Error)
	// Performs copy-assignment to destination
	CopyTo(destination T) (err Error)
}

// DataType_Equal provide comparable two data type.
type DataType_Equal[T any] interface {
	Equal(with T) bool
}

// **ATTENTION**::: strongly suggest use DataType_Accessor to prevent invalid state at first place.
type DataType_Validation interface {
	Validate() (err Error)
}

type DataType_Locker interface {
	Lock() (err Error)
	Unlock() (err Error)
}

// DataType_ExpectedLength indicate min and max expected length.
type DataType_ExpectedLength interface {
	MinLength() NumberOfElement
	MaxLength() NumberOfElement
}

type DataType_OptionalValue interface {
	// false means data required and must be exist.
	Optional() bool
}

type DataType_Existence interface {
	Heap() bool // Pointer
}
