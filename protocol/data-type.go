/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// DataType in computer science and computer programming, is a collection or grouping of data values,
// usually specified by a set of possible values, a set of allowed operations on these values,
// and/or a representation of these values as machine types even in compilers or runtime packages.
// It can use for any data like CLA flags or json fields or any other data structures
// https://en.wikipedia.org/wiki/Data_type
type DataType /*[T any]*/ interface {
	// DataType_DefaultValue[T]
	// DataType_Accessor[T]
	// DataType_AtomicAccessor[T]
	// DataType_Clone[T]
	// DataType_Copy[T]

	// DataType_Locker

	// TODO::: add more

	// existence length
	// DataType_ExpectedLen
	// Len
	DataType_Details
	Stringer
}

type DataType_Details interface {
	Status() SoftwareStatus
	ReferenceURI() string
	IssueDate() Time
	ExpiryDate() Time
	ExpireInFavorOf() DataType

	Details
}

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
	// Returns a copy of the itself
	Clone() T
	// Performs copy-assignment from source
	CloneFrom(source T)
}

// DataType_Copy is implicit, inexpensive, and cannot be re-implemented (memcpy).
// The Copy trait represents values that can be safely duplicated via memcpy:
// things like reassignments and passing an argument by-value to a function are always memcpys, and so for Copy types,
// the compiler understands that it doesn't need to consider those a move.
// every Copy type is also required to be Clone
type DataType_Copy[T any] interface {
	// Returns a copy of the itself
	Copy() T
	// Performs copy-assignment from source
	CopyFrom(source T)
}

type DataType_Locker interface {
	Lock()
	Unlock()
}

// DataType_ExpectedLen or Expected length
type DataType_ExpectedLen interface {
	MinLen() int // in byte or 8bit
	MaxLen() int // in byte or 8bit
}

type DataType_OptionalValue interface {
	// false means data required and must be exist.
	Optional() bool
}

type DataType_Existence interface {
	Heap() bool // Pointer
}
